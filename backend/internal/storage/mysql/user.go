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

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) refractor.UserRepository {
	return &userRepo{
		db: db,
	}
}

// The provided user must have the following fields set: Username, Email, Password, Permissions
func (r *userRepo) Create(user *refractor.User) error {
	query := "INSERT INTO Users (Username, Email, Password, Permissions) VALUES (?, ?, ?, ?);"

	res, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Permissions)
	if err != nil {
		return wrapError(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return wrapError(err)
	}

	// Update UserID on passed in user
	user.UserID = id

	return nil
}

func (r *userRepo) FindByID(id int64) (*refractor.User, error) {
	query := "SELECT * FROM Users WHERE UserID = ?;"

	row := r.db.QueryRow(query, id)

	foundUser := &refractor.User{}
	if err := r.scanRow(row, foundUser); err != nil {
		return nil, wrapError(err)
	}

	return foundUser, nil
}

func (r *userRepo) Exists(args refractor.FindArgs) (bool, error) {
	query, values := buildExistsQuery("Users", args)

	var exists bool

	if err := r.db.QueryRow(query, values...).Scan(&exists); err != nil {
		return false, wrapError(err)
	}

	return exists, nil
}

func (r *userRepo) FindOne(args refractor.FindArgs) (*refractor.User, error) {
	query, values := buildFindQuery("Users", args)

	foundUser := &refractor.User{}

	row := r.db.QueryRow(query, values...)
	if err := r.scanRow(row, foundUser); err != nil {
		return nil, wrapError(err)
	}

	return foundUser, nil
}

func (r *userRepo) FindMany(args refractor.FindArgs) ([]*refractor.User, error) {
	query, values := buildFindQuery("Users", args)

	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, wrapError(err)
	}

	var foundUsers []*refractor.User

	for rows.Next() {
		user := &refractor.User{}

		if err := r.scanRows(rows, user); err != nil {
			return nil, wrapError(err)
		}

		foundUsers = append(foundUsers, user)
	}

	return foundUsers, nil
}

func (r *userRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.User, error) {
	query, values := buildUpdateQuery("Users", id, "UserID", args)

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return nil, wrapError(err)
	}

	query = "SELECT * FROM Users WHERE UserID = ?;"
	row := r.db.QueryRow(query, id)

	updatedUser := &refractor.User{}
	if err = r.scanRow(row, updatedUser); err != nil {
		return nil, wrapError(err)
	}

	return updatedUser, nil
}

func (r *userRepo) FindAll() ([]*refractor.User, error) {
	query := "SELECT * FROM Users;"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, wrapError(err)
	}

	var users []*refractor.User

	for rows.Next() {
		user := &refractor.User{}

		err := r.scanRows(rows, user)
		if err != nil {
			return nil, wrapError(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *userRepo) GetCount() int {
	query := "SELECT COUNT(1) AS Count FROM Users;"

	var userCount int

	row := r.db.QueryRow(query)

	if err := row.Scan(&userCount); err != nil {
		return 0
	}

	return userCount
}

// Scan helpers
func (r *userRepo) scanRow(row *sql.Row, user *refractor.User) error {
	return row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Permissions, &user.Activated, &user.NeedsPasswordChange)
}

func (r *userRepo) scanRows(rows *sql.Rows, user *refractor.User) error {
	return rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Permissions, &user.Activated, &user.NeedsPasswordChange)
}
