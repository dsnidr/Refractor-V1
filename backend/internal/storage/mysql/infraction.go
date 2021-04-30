/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package mysql

import (
	"database/sql"
	"github.com/sniddunc/refractor/refractor"
	"time"
)

type infractionRepo struct {
	db *sql.DB
}

func NewInfractionRepository(db *sql.DB) refractor.InfractionRepository {
	return &infractionRepo{
		db: db,
	}
}

func (r *infractionRepo) Create(infraction *refractor.DBInfraction) (*refractor.Infraction, error) {
	if infraction.Timestamp == 0 {
		infraction.Timestamp = time.Now().Unix()
	}

	query := "INSERT INTO Infractions(PlayerID, UserID, ServerID, Type, Reason, Duration, Timestamp, SystemAction) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"

	res, err := r.db.Exec(query, infraction.PlayerID, infraction.UserID, infraction.ServerID, infraction.Type,
		infraction.Reason, infraction.Duration, infraction.Timestamp, infraction.SystemAction)
	if err != nil {
		return nil, wrapError(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, wrapError(err)
	}

	infraction.InfractionID = id

	return infraction.Infraction(), nil
}

func (r *infractionRepo) FindByID(id int64) (*refractor.Infraction, error) {
	query := "SELECT * FROM Infractions WHERE InfractionID = ?;"
	row := r.db.QueryRow(query, id)

	foundInfraction := &refractor.DBInfraction{}
	if err := r.scanRow(row, foundInfraction); err != nil {
		return nil, wrapError(err)
	}

	return foundInfraction.Infraction(), nil
}

func (r *infractionRepo) Exists(args refractor.FindArgs) (bool, error) {
	query, values := buildExistsQuery("Infractions", args)

	var exists bool

	row := r.db.QueryRow(query, values...)
	if err := row.Scan(&exists); err != nil {
		return false, wrapError(err)
	}

	return exists, nil
}

func (r *infractionRepo) FindOne(args refractor.FindArgs) (*refractor.Infraction, error) {
	query, values := buildFindQuery("Infractions", args)

	foundInfraction := &refractor.DBInfraction{}

	row := r.db.QueryRow(query, values...)
	if err := r.scanRow(row, foundInfraction); err != nil {
		return nil, wrapError(err)
	}

	return foundInfraction.Infraction(), nil
}

func (r *infractionRepo) FindMany(args refractor.FindArgs) ([]*refractor.Infraction, error) {
	query, values := buildFindQuery("Infractions", args)

	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, wrapError(err)
	}

	var foundInfractions []*refractor.Infraction

	for rows.Next() {
		infraction := &refractor.DBInfraction{}

		if err := r.scanRows(rows, infraction); err != nil {
			return nil, wrapError(err)
		}

		foundInfractions = append(foundInfractions, infraction.Infraction())
	}

	return foundInfractions, nil
}

func (r *infractionRepo) FindManyByPlayerID(playerID int64) ([]*refractor.Infraction, error) {
	query := "SELECT * FROM Infractions WHERE PlayerID = ?;"

	rows, err := r.db.Query(query, playerID)
	if err != nil {
		return nil, wrapError(err)
	}

	var foundInfractions []*refractor.Infraction

	for rows.Next() {
		infraction := &refractor.DBInfraction{}

		if err := r.scanRows(rows, infraction); err != nil {
			return nil, wrapError(err)
		}

		foundInfractions = append(foundInfractions, infraction.Infraction())
	}

	return foundInfractions, nil
}

func (r *infractionRepo) FindAll() ([]*refractor.Infraction, error) {
	query := "SELECT * FROM Infractions;"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, wrapError(err)
	}

	var foundInfractions []*refractor.Infraction

	for rows.Next() {
		infraction := &refractor.DBInfraction{}

		if err := r.scanRows(rows, infraction); err != nil {
			return nil, wrapError(err)
		}

		foundInfractions = append(foundInfractions, infraction.Infraction())
	}

	return foundInfractions, nil
}

func (r *infractionRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.Infraction, error) {
	query, values := buildUpdateQuery("Infractions", id, "InfractionID", args)

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return nil, wrapError(err)
	}

	// Retrieve updated infraction
	query = "SELECT * FROM Infractions WHERE InfractionID = ?;"
	row := r.db.QueryRow(query, id)

	updatedInfraction := &refractor.DBInfraction{}
	if err := r.scanRow(row, updatedInfraction); err != nil {
		return nil, wrapError(err)
	}

	return updatedInfraction.Infraction(), nil
}

func (r *infractionRepo) Delete(id int64) error {
	query := "DELETE FROM Infractions WHERE InfractionID = ?;"

	res, err := r.db.Exec(query, id)
	if err != nil {
		return wrapError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return wrapError(err)
	}

	if rowsAffected <= 0 {
		return wrapError(sql.ErrNoRows)
	}

	return nil
}

func (r *infractionRepo) Search(args refractor.FindArgs, limit int, offset int) (int, []*refractor.Infraction, error) {
	query := `
		SELECT
			res.*,
			u.Username AS StaffName
		FROM (
			SELECT
				i.*
			FROM Infractions i
			INNER JOIN Servers s ON i.ServerID = s.ServerID
			WHERE
				(? IS NULL OR i.Type = ?) AND
				(? IS NULL OR i.PlayerID = ?) AND
				(? IS NULL OR i.UserID = ?) AND
				(? IS NULL OR i.ServerID = ?) AND
				(? IS NULL OR s.Game = ?)
			) res
		JOIN Users u ON res.UserID = u.UserID
		GROUP BY InfractionID
		LIMIT ? OFFSET ?;
	`

	var (
		iType    = args["Type"]
		playerID = args["PlayerID"]
		userID   = args["UserID"]
		serverID = args["ServerID"]
		game     = args["Game"]
	)

	rows, err := r.db.Query(query, iType, iType, playerID, playerID, userID, userID, serverID, serverID, game, game, limit, offset)
	if err != nil {
		return 0, nil, wrapError(err)
	}

	var foundInfractions []*refractor.Infraction

	for rows.Next() {
		dbinfr := &refractor.DBInfraction{}

		var staffName string
		if err := rows.Scan(&dbinfr.InfractionID, &dbinfr.PlayerID, &dbinfr.UserID, &dbinfr.ServerID,
			&dbinfr.Type, &dbinfr.Reason, &dbinfr.Duration, &dbinfr.Timestamp, &dbinfr.SystemAction, &staffName); err != nil {
			return 0, nil, wrapError(err)
		}

		infraction := dbinfr.Infraction()

		// Get player's name here since I can't figure out how to do it in the query in a reasonable amount of time.
		nameQuery := `SELECT Name FROM PlayerNames WHERE PlayerID = ? ORDER BY DateRecorded DESC LIMIT 1`

		row := r.db.QueryRow(nameQuery, infraction.PlayerID)

		var playerName string
		if err := row.Scan(&playerName); err != nil {
			return 0, nil, wrapError(err)
		}

		// Set staff and player name
		infraction.StaffName = staffName
		infraction.PlayerName = playerName

		// Append to list of results
		foundInfractions = append(foundInfractions, infraction)
	}

	// Get total number of matches
	query = `
		SELECT
			COUNT(1) AS Count
		FROM Infractions i
		INNER JOIN Servers s ON i.ServerID = s.ServerID
		WHERE
			(? IS NULL OR i.Type = ?) AND
			(? IS NULL OR i.PlayerID = ?) AND
			(? IS NULL OR i.UserID = ?) AND
			(? IS NULL OR i.ServerID = ?) AND
			(? IS NULL OR s.Game = ?)
	`

	row := r.db.QueryRow(query, iType, iType, playerID, playerID, userID, userID, serverID, serverID, game, game)

	var count int
	if err := row.Scan(&count); err != nil {
		return 0, nil, wrapError(err)
	}

	return count, foundInfractions, nil
}

func (r *infractionRepo) GetRecent(count int) ([]*refractor.Infraction, error) {
	query := `
		SELECT
			i.*,
			u.Username AS StaffName
		FROM Infractions i
		INNER JOIN Users u ON u.UserID = i.UserID
		ORDER BY Timestamp DESC LIMIT ?;
	`

	rows, err := r.db.Query(query, count)
	if err != nil {
		return nil, wrapError(err)
	}

	var foundInfractions []*refractor.Infraction

	for rows.Next() {
		dbinfr := &refractor.DBInfraction{}

		var staffName string
		if err := rows.Scan(&dbinfr.InfractionID, &dbinfr.PlayerID, &dbinfr.UserID, &dbinfr.ServerID,
			&dbinfr.Type, &dbinfr.Reason, &dbinfr.Duration, &dbinfr.Timestamp, &dbinfr.SystemAction, &staffName); err != nil {
			return nil, wrapError(err)
		}

		infraction := dbinfr.Infraction()

		// Get player's name here since I can't figure out how to do it in the query in a reasonable amount of time.
		nameQuery := `SELECT Name FROM PlayerNames WHERE PlayerID = ? ORDER BY DateRecorded DESC LIMIT 1`

		row := r.db.QueryRow(nameQuery, infraction.PlayerID)

		var playerName string
		if err := row.Scan(&playerName); err != nil {
			return nil, wrapError(err)
		}

		// Set staff and player name
		infraction.StaffName = staffName
		infraction.PlayerName = playerName

		// Append to list of results
		foundInfractions = append(foundInfractions, infraction)
	}

	return foundInfractions, nil
}

func (r *infractionRepo) GetCountByPlayerID(playerID int64) (int, error) {
	query := "SELECT COUNT(1) FROM Infractions WHERE PlayerID = ?;"

	row := r.db.QueryRow(query, playerID)

	var count int

	if err := row.Scan(&count); err != nil {
		return 0, wrapError(err)
	}

	return count, nil
}

// Scan helpers
func (r *infractionRepo) scanRow(row *sql.Row, infr *refractor.DBInfraction) error {
	return row.Scan(&infr.InfractionID, &infr.PlayerID, &infr.UserID, &infr.ServerID, &infr.Type, &infr.Reason,
		&infr.Duration, &infr.Timestamp, &infr.SystemAction)
}

func (r *infractionRepo) scanRows(row *sql.Rows, infr *refractor.DBInfraction) error {
	return row.Scan(&infr.InfractionID, &infr.PlayerID, &infr.UserID, &infr.ServerID, &infr.Type, &infr.Reason,
		&infr.Duration, &infr.Timestamp, &infr.SystemAction)
}
