package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sniddunc/refractor/refractor"
)

func Setup(db *sql.DB) error {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Could not create transaction. Error: %v", err)
	}

	// Create users table
	if _, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS Users(
			UserID INT NOT NULL AUTO_INCREMENT,
			Username VARCHAR(32) UNIQUE NOT NULL,
			Email VARCHAR(254) UNIQUE NOT NULL,
			Password BINARY(60) NOT NULL,
			AccessLevel INT DEFAULT 0,
			Activated BOOLEAN DEFAULT TRUE,
			NeedsPasswordChange BOOLEAN DEFAULT TRUE,
			
			PRIMARY KEY (UserID)
		);
	`); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}

		return fmt.Errorf("could not create Users table. Error: %v", err)
	}

	return tx.Commit()
}

// MySQL query builder and helper functions
func wrapError(err error) error {
	switch err {
	case nil:
		return err
	case sql.ErrNoRows:
		return refractor.ErrNotFound
	default:
		return err
	}
}

func buildExistsQuery(table string, args map[string]interface{}) (string, []interface{}) {
	var query string = fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE ", table)
	var values []interface{}

	// Build query
	for key, val := range args {
		query += key + " = ? AND "
		values = append(values, val)
	}

	// Cut off trailing AND
	query = query[:len(query)-5] + ");"

	return query, values
}

func buildFindQuery(table string, args map[string]interface{}) (string, []interface{}) {
	var query string = fmt.Sprintf("SELECT * FROM %s WHERE ", table)
	var values []interface{}

	// Build query
	for key, val := range args {
		query += key + " = ? AND "
		values = append(values, val)
	}

	// Cut off trailing AND
	query = query[:len(query)-5] + ";"

	return query, values
}

func buildUpdateQuery(table string, id int64, idName string, args map[string]interface{}) (string, []interface{}) {
	var query string = fmt.Sprintf("UPDATE %s SET ", table)
	var values []interface{}

	// Build query
	for key, val := range args {
		query += key + " = ?, "
		values = append(values, val)
	}

	values = append(values, id)

	// Cut off trailing comma and space
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE %s = ?;", idName)

	return query, values
}
