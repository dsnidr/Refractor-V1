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

func (r *userRepo) Create(user *refractor.User) error {
	query := "INSERT INTO Users (Username, Email, Password) VALUES (?, ?, ?);"

	res, err := r.db.Exec(query, user.Username, user.Email, user.Password)
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
	panic("implement me")
}

func (r *userRepo) FindAll() ([]*refractor.User, error) {
	panic("implement me")
}

func (r *userRepo) GetCount() int {
	panic("implement me")
}

// Scan helpers
func (r *userRepo) scanRow(row *sql.Row, user *refractor.User) error {
	return row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.AccessLevel, &user.Activated, &user.NeedsPasswordChange)
}

func (r *userRepo) scanRows(rows *sql.Rows, user *refractor.User) error {
	return rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.AccessLevel, &user.Activated, &user.NeedsPasswordChange)
}
