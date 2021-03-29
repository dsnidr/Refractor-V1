package mysql

import (
	"database/sql"
	"fmt"
	"github.com/sniddunc/refractor/refractor"
	"time"
	"unicode/utf8"
)

type playerRepo struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) refractor.PlayerRepository {
	return &playerRepo{
		db: db,
	}
}

// Create inserts a player into the Players table as well as inserting their current name into the PlayerNames table.
// The following values must be present on the passed in Player reference for Create to function properly:
// PlayFabID, LastSeen and CurrentName.
func (r *playerRepo) Create(player *refractor.DBPlayer) error {
	query := "INSERT INTO Players (PlayFabID, MCUUID, LastSeen) VALUES (?, ?, ?);"

	res, err := r.db.Exec(query, player.PlayFabID, player.MCUUID, player.LastSeen)
	if err != nil {
		return wrapError(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return wrapError(err)
	}

	player.PlayerID = id

	// Insert into PlayerNames table
	query = "INSERT INTO PlayerNames (PlayerID, Name, DateRecorded) VALUES (?, ?, ?);"

	name := player.CurrentName

	if !utf8.ValidString(name) {
		outputName := make([]rune, 0, len(name))
		for i, r := range name {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(name[i:])
				if size == 1 {
					continue
				}
			}

			outputName = append(outputName, r)
		}

		name = string(outputName)
	}

	// If the name is empty, it means it was all invalid unicode characters so we replace with a known string.
	if name == "" {
		name = "Invalid name"
	}

	if _, err = r.db.Exec(query, id, name, time.Now().Unix()); err != nil {
		return wrapError(err)
	}

	return nil
}

func (r *playerRepo) FindByID(id int64) (*refractor.Player, error) {
	query := "SELECT * FROM Players WHERE PlayerID = ?;"

	row := r.db.QueryRow(query, id)

	foundPlayer := &refractor.DBPlayer{}
	if err := r.scanRow(row, foundPlayer); err != nil {
		return nil, wrapError(err)
	}

	// Get player names
	currentName, previousNames, err := r.getPlayerNames(foundPlayer.PlayerID)
	if err != nil {
		return nil, wrapError(err)
	}

	foundPlayer.CurrentName = currentName
	foundPlayer.PreviousNames = previousNames

	return foundPlayer.Player(), nil
}

func (r *playerRepo) FindByPlayFabID(playFabID string) (*refractor.Player, error) {
	query := "SELECT * FROM Players WHERE PlayFabID = ?;"

	row := r.db.QueryRow(query, playFabID)

	foundPlayer := &refractor.DBPlayer{}
	if err := r.scanRow(row, foundPlayer); err != nil {
		return nil, wrapError(err)
	}

	// Get player names
	currentName, previousNames, err := r.getPlayerNames(foundPlayer.PlayerID)
	if err != nil {
		return nil, wrapError(err)
	}

	foundPlayer.CurrentName = currentName
	foundPlayer.PreviousNames = previousNames

	return foundPlayer.Player(), nil
}

func (r *playerRepo) FindByMCUUID(MCUUID string) (*refractor.Player, error) {
	query := "SELECT * FROM Players WHERE MCUUID = ?;"

	row := r.db.QueryRow(query, MCUUID)

	foundPlayer := &refractor.DBPlayer{}
	if err := r.scanRow(row, foundPlayer); err != nil {
		return nil, wrapError(err)
	}

	// Get player names
	currentName, previousNames, err := r.getPlayerNames(foundPlayer.PlayerID)
	if err != nil {
		return nil, wrapError(err)
	}

	foundPlayer.CurrentName = currentName
	foundPlayer.PreviousNames = previousNames

	return foundPlayer.Player(), nil
}

func (r *playerRepo) Exists(args refractor.FindArgs) (bool, error) {
	query, values := buildExistsQuery("Players", args)

	var exists bool

	if err := r.db.QueryRow(query, values...).Scan(&exists); err != nil {
		return false, wrapError(err)
	}

	return exists, nil
}

func (r *playerRepo) UpdateName(player *refractor.Player, currentName string) error {
	query := `
		INSERT INTO PlayerNames (PlayerID, Name, DateRecorded) VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE DateRecorded = UNIX_TIMESTAMP();
	`

	runeName := []rune(currentName)

	if _, err := r.db.Exec(query, player.PlayerID, string(runeName), time.Now().Unix()); err != nil {
		return wrapError(err)
	}

	// Get updated names list
	updatedCurrentName, previousNames, err := r.getPlayerNames(player.PlayerID)
	if err != nil {
		return wrapError(err)
	}

	// Set names
	player.CurrentName = updatedCurrentName
	player.PreviousNames = previousNames

	return nil
}

func (r *playerRepo) FindOne(args refractor.FindArgs) (*refractor.Player, error) {
	query, values := buildFindQuery("Players", args)

	row := r.db.QueryRow(query, values...)

	var foundPlayer = &refractor.DBPlayer{}

	if err := r.scanRow(row, foundPlayer); err != nil {
		return nil, wrapError(err)
	}

	// Get names
	currentName, previousNames, err := r.getPlayerNames(foundPlayer.PlayerID)
	if err != nil {
		return nil, wrapError(err)
	}

	// Set names
	foundPlayer.CurrentName = currentName
	foundPlayer.PreviousNames = previousNames

	return foundPlayer.Player(), nil
}

func (r *playerRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.Player, error) {
	query, values := buildUpdateQuery("Players", id, "PlayerID", args)

	if _, err := r.db.Exec(query, values...); err != nil {
		return nil, wrapError(err)
	}

	query = "SELECT * FROM Players WHERE PlayerID = ?;"
	row := r.db.QueryRow(query, id)

	updatedPlayer := &refractor.DBPlayer{}
	if err := r.scanRow(row, updatedPlayer); err != nil {
		return nil, wrapError(err)
	}

	// Get names
	currentName, previousNames, err := r.getPlayerNames(updatedPlayer.PlayerID)
	if err != nil {
		return nil, wrapError(err)
	}

	// Set names
	updatedPlayer.CurrentName = currentName
	updatedPlayer.PreviousNames = previousNames

	return updatedPlayer.Player(), nil
}

func (r *playerRepo) SearchByName(name string, limit int, offset int) (int, []*refractor.Player, error) {
	query := `
		SELECT
			res.*
		FROM (
			SELECT p.* FROM PlayerNames pn
			INNER JOIN Players p ON p.PlayerID = pn.PlayerID
			WHERE pn.Name LIKE CONCAT('%', ?, '%')
			LIMIT ? OFFSET ?
		) res
		GROUP BY PlayerID
		ORDER BY LastSeen DESC
	`

	rows, err := r.db.Query(query, name, limit, offset)
	if err != nil {
		return 0, nil, wrapError(err)
	}

	var foundPlayers []*refractor.Player

	for rows.Next() {
		foundPlayer := &refractor.DBPlayer{}

		if err := r.scanRows(rows, foundPlayer); err != nil {
			return 0, nil, wrapError(err)
		}

		// Get names list
		currentName, previousNames, err := r.getPlayerNames(foundPlayer.PlayerID)
		if err != nil {
			return 0, nil, wrapError(err)
		}

		// Set names
		foundPlayer.CurrentName = currentName
		foundPlayer.PreviousNames = previousNames

		foundPlayers = append(foundPlayers, foundPlayer.Player())
	}

	// Get number of possible matches
	query = `
		SELECT
			COUNT(1) AS MatchCount
		FROM (
			SELECT 1 FROM PlayerNames pn
			WHERE pn.Name LIKE CONCAT('%', ?, '%')
		    GROUP BY PlayerID
		) AS matches;
	`

	row := r.db.QueryRow(query, name)

	var count int

	if err := row.Scan(&count); err != nil {
		return 0, nil, wrapError(err)
	}

	return count, foundPlayers, nil
}

func (r *playerRepo) getPlayerNames(playerID int64) (string, []string, error) {
	query := "SELECT Name FROM PlayerNames WHERE PlayerID = ? ORDER BY DateRecorded DESC;"

	rows, err := r.db.Query(query, playerID)
	if err != nil {
		return "", nil, err
	}

	var names []string

	for rows.Next() {
		name := ""

		err = rows.Scan(&name)
		if err != nil {
			return "", nil, err
		}

		names = append(names, name)
	}

	if names == nil {
		return "", nil, fmt.Errorf("names slice was nil")
	}

	// PlayerNames are ordered in descending order by DateRecorded so index 0 is the most recent name
	return names[0], names[1:], nil
}

// Scan helpers
func (r *playerRepo) scanRow(row *sql.Row, player *refractor.DBPlayer) error {
	return row.Scan(&player.PlayerID, &player.PlayFabID, &player.MCUUID, &player.LastSeen, &player.Watched)
}

func (r *playerRepo) scanRows(rows *sql.Rows, player *refractor.DBPlayer) error {
	return rows.Scan(&player.PlayerID, &player.PlayFabID, &player.MCUUID, &player.LastSeen, &player.Watched)
}
