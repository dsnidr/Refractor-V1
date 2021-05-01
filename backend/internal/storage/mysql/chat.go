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

type chatRepo struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) refractor.ChatRepository {
	return &chatRepo{
		db: db,
	}
}

func (r *chatRepo) Create(message *refractor.ChatMessage) (*refractor.ChatMessage, error) {
	if message.DateRecorded == 0 {
		message.DateRecorded = time.Now().Unix()
	}

	query := `INSERT INTO ChatMessages (PlayerID, ServerID, Message, DateRecorded, Flagged)
			VALUES (?, ?, ?, FROM_UNIXTIME(?), ?);`

	res, err := r.db.Exec(query, message.PlayerID, message.ServerID, message.Message, message.DateRecorded, message.Flagged)
	if err != nil {
		return nil, wrapError(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, wrapError(err)
	}

	message.MessageID = id

	return message, nil
}

func (r *chatRepo) FindByID(id int64) (*refractor.ChatMessage, error) {
	query := `SELECT MessageID, PlayerID, ServerID, Message, UNIX_TIMESTAMP(DateRecorded) AS DateRecorded, Flagged
			FROM ChatMessages WHERE MessageID = ?;`

	row := r.db.QueryRow(query, id)

	var message *refractor.ChatMessage

	if err := r.scanRow(row, message); err != nil {
		return nil, wrapError(err)
	}

	return message, nil
}

func (r *chatRepo) FindMany(args refractor.FindArgs) ([]*refractor.ChatMessage, error) {
	query, values := buildFindQuery("ChatMessage", args)

	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, wrapError(err)
	}

	var foundMessages []*refractor.ChatMessage

	for rows.Next() {
		message := &refractor.ChatMessage{}

		if err := r.scanRows(rows, message); err != nil {
			return nil, wrapError(err)
		}

		foundMessages = append(foundMessages, message)
	}

	return foundMessages, nil
}

// Scan helpers
func (r *chatRepo) scanRow(row *sql.Row, msg *refractor.ChatMessage) error {
	return row.Scan(&msg.MessageID, &msg.PlayerID, &msg.ServerID, &msg.MessageID, &msg.DateRecorded, &msg.Flagged)
}

func (r *chatRepo) scanRows(rows *sql.Rows, msg *refractor.ChatMessage) error {
	return rows.Scan(&msg.MessageID, &msg.PlayerID, &msg.ServerID, &msg.MessageID, &msg.DateRecorded, &msg.Flagged)
}
