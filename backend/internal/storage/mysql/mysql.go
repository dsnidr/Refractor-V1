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
