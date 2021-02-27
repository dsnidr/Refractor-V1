package mysql

import (
	"database/sql"
	"fmt"
	"github.com/sniddunc/refractor/refractor"
	"time"
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
func (r *playerRepo) Create(player *refractor.Player) error {
	query := "INSERT INTO Players (PlayFabID, LastSeen) VALUES (?, ?);"

	res, err := r.db.Exec(query, player.PlayFabID, player.LastSeen)
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

	if _, err = r.db.Exec(query, id, player.CurrentName, time.Now().Unix()); err != nil {
		return wrapError(err)
	}

	return nil
}

func (r *playerRepo) FindByID(id int64) (*refractor.Player, error) {
	query := "SELECT * FROM Players WHERE PlayerID = ?;"

	row := r.db.QueryRow(query, id)

	foundPlayer := &refractor.Player{}
	if err := r.scanRow(row, foundPlayer); err != nil {
		return nil, wrapError(err)
	}

	return foundPlayer, nil
}

func (r *playerRepo) FindByPlayFabID(playFabID string) (*refractor.Player, error) {
	query := "SELECT * FROM Players WHERE PlayFabID = ?;"

	row := r.db.QueryRow(query, playFabID)

	foundPlayer := &refractor.Player{}
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

	return foundPlayer, nil
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

	var foundPlayer = &refractor.Player{}

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

	return foundPlayer, nil
}

func (r *playerRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.Player, error) {
	query, values := buildUpdateQuery("Players", id, "PlayerID", args)

	if _, err := r.db.Exec(query, values...); err != nil {
		return nil, wrapError(err)
	}

	query = "SELECT * FROM Players WHERE PlayerID = ?;"
	row := r.db.QueryRow(query, id)

	updatedPlayer := &refractor.Player{}
	if err := r.scanRow(row, updatedPlayer); err != nil {
		return nil, wrapError(err)
	}

	return updatedPlayer, nil
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
func (r *playerRepo) scanRow(row *sql.Row, player *refractor.Player) error {
	return row.Scan(&player.PlayerID, &player.PlayFabID, &player.LastSeen)
}

func (r *playerRepo) scanRows(rows *sql.Rows, player *refractor.Player) error {
	return rows.Scan(&player.PlayerID, &player.PlayFabID, &player.LastSeen)
}
