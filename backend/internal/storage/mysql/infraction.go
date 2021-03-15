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

func (r *infractionRepo) Create(infraction *refractor.Infraction) (*refractor.Infraction, error) {
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

	infraction.ServerID = id

	return infraction, nil
}

func (r *infractionRepo) FindByID(id int64) (*refractor.Infraction, error) {
	query := "SELECT * FROM Servers WHERE ServerID = ?;"
	row := r.db.QueryRow(query, id)

	foundInfraction := &refractor.Infraction{}
	if err := r.scanRow(row, foundInfraction); err != nil {
		return nil, wrapError(err)
	}

	return foundInfraction, nil
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

	foundInfraction := &refractor.Infraction{}

	row := r.db.QueryRow(query, values...)
	if err := r.scanRow(row, foundInfraction); err != nil {
		return nil, wrapError(err)
	}

	return foundInfraction, nil
}

func (r *infractionRepo) FindManyByPlayerID(playerID int64) ([]*refractor.Infraction, error) {
	query := "SELECT * FROM Infractions WHERE PlayerID = ?;"

	rows, err := r.db.Query(query, playerID)
	if err != nil {
		return nil, wrapError(err)
	}

	var foundInfractions []*refractor.Infraction

	for rows.Next() {
		infraction := &refractor.Infraction{}

		if err := r.scanRows(rows, infraction); err != nil {
			return nil, wrapError(err)
		}

		foundInfractions = append(foundInfractions, infraction)
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
		infraction := &refractor.Infraction{}

		if err := r.scanRows(rows, infraction); err != nil {
			return nil, wrapError(err)
		}

		foundInfractions = append(foundInfractions, infraction)
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

	updatedInfraction := &refractor.Infraction{}
	if err := r.scanRow(row, updatedInfraction); err != nil {
		return nil, wrapError(err)
	}

	return updatedInfraction, nil
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

// Scan helpers
func (r *infractionRepo) scanRow(row *sql.Row, infr *refractor.Infraction) error {
	return row.Scan(infr.InfractionID, infr.PlayerID, infr.UserID, infr.ServerID, infr.Type, infr.Reason, infr.Duration, infr.Timestamp, infr.SystemAction)
}

func (r *infractionRepo) scanRows(row *sql.Rows, infr *refractor.Infraction) error {
	return row.Scan(infr.InfractionID, infr.PlayerID, infr.UserID, infr.ServerID, infr.Type, infr.Reason, infr.Duration, infr.Timestamp, infr.SystemAction)
}
