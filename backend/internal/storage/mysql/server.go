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
)

type serverRepo struct {
	db *sql.DB
}

func NewServerRepository(db *sql.DB) refractor.ServerRepository {
	return &serverRepo{
		db: db,
	}
}

func (r *serverRepo) Create(server *refractor.Server) error {
	query := "INSERT INTO Servers (Game, Name, Address, RCONPort, RCONPassword) VALUES (?, ?, ?, ?, ?);"

	res, err := r.db.Exec(query, server.Game, server.Name, server.Address, server.RCONPort, server.RCONPassword)
	if err != nil {
		return wrapError(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return wrapError(err)
	}

	server.ServerID = id

	return nil
}

func (r *serverRepo) FindByID(id int64) (*refractor.Server, error) {
	query := "SELECT * FROM Servers WHERE ServerID = ?;"
	row := r.db.QueryRow(query, id)

	foundServer := &refractor.Server{}
	if err := r.scanRow(row, foundServer); err != nil {
		return nil, wrapError(err)
	}

	return foundServer, nil
}

func (r *serverRepo) Exists(args refractor.FindArgs) (bool, error) {
	query, values := buildExistsQuery("Servers", args)

	var exists bool

	row := r.db.QueryRow(query, values...)
	if err := row.Scan(&exists); err != nil {
		return false, wrapError(err)
	}

	return exists, nil
}

func (r *serverRepo) FindOne(args refractor.FindArgs) (*refractor.Server, error) {
	query, values := buildFindQuery("Servers", args)

	foundServer := &refractor.Server{}

	row := r.db.QueryRow(query, values...)
	if err := r.scanRow(row, foundServer); err != nil {
		return nil, wrapError(err)
	}

	return foundServer, nil
}

func (r *serverRepo) FindAll() ([]*refractor.Server, error) {
	query := "SELECT * FROM Servers;"

	var foundServers []*refractor.Server

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		server := &refractor.Server{}

		if err := r.scanRows(rows, server); err != nil {
			return nil, wrapError(err)
		}

		foundServers = append(foundServers, server)
	}

	return foundServers, nil
}

func (r *serverRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.Server, error) {
	query, values := buildUpdateQuery("Servers", id, "ServerID", args)

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return nil, wrapError(err)
	}

	// Retrieve updated server
	query = "SELECT * FROM Servers WHERE ServerID = ?;"
	row := r.db.QueryRow(query, id)

	updatedServer := &refractor.Server{}
	if err := r.scanRow(row, updatedServer); err != nil {
		return nil, wrapError(err)
	}

	return updatedServer, nil
}

func (r *serverRepo) Delete(id int64) error {
	query := "DELETE FROM Servers WHERE ServerID = ?;"

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

// Scan helpers
func (r *serverRepo) scanRow(row *sql.Row, server *refractor.Server) error {
	return row.Scan(&server.ServerID, &server.Game, &server.Name, &server.Address, &server.RCONPort, &server.RCONPassword)
}

func (r *serverRepo) scanRows(rows *sql.Rows, server *refractor.Server) error {
	return rows.Scan(&server.ServerID, &server.Game, &server.Name, &server.Address, &server.RCONPort, &server.RCONPassword)
}
