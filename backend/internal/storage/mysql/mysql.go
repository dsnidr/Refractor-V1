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
			Permissions BIGINT UNSIGNED DEFAULT 0,
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

	// Create servers table
	if _, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS Servers(
			ServerID INT NOT NULL AUTO_INCREMENT,
			Game VARCHAR(32) NOT NULL,
			Name VARCHAR(32) UNIQUE NOT NULL,
			Address VARCHAR(15) NOT NULL,
		    RCONPort VARCHAR(5) NOT NULL,
		    RCONPassword VARCHAR(128) NOT NULL,
			
			PRIMARY KEY (ServerID)
		);
	`); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}

		return fmt.Errorf("could not create Servers table. Error: %v", err)
	}

	// Create players table
	if _, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS Players(
			PlayerID INT NOT NULL AUTO_INCREMENT,
			PlayFabID VARCHAR(32) UNIQUE,
		    MCUUID VARCHAR(36) UNIQUE,
			LastSeen BIGINT DEFAULT 0,
			
			PRIMARY KEY (PlayerID)
		);
	`); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}

		return fmt.Errorf("could not create Players table. Error: %v", err)
	}

	// Create player names table
	if _, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS PlayerNames(
			PlayerID INT NOT NULL,
			Name VARCHAR(128) CHARACTER SET utf8mb4 NOT NULL,
			DateRecorded BIGINT DEFAULT 0,
			
			FOREIGN KEY (PlayerID) REFERENCES Players(PlayerID),
		    PRIMARY KEY (PlayerID, Name)
		);
	`); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}

		return fmt.Errorf("could not create PlayerNames table. Error: %v", err)
	}

	// Alter player names table
	if _, err := tx.Exec(`
		ALTER TABLE PlayerNames CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
	`); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}

		return fmt.Errorf("could not alter PlayerNames table. Error: %v", err)
	}

	// Create infractions table
	if _, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS Infractions (
			InfractionID INT NOT NULL AUTO_INCREMENT,
			PlayerID INT NOT NULL,
			UserID INT NOT NULL,
			ServerID INT NOT NULL,
			Type ENUM("WARNING", "MUTE", "KICK", "BAN") NOT NULL,
			Reason TEXT,
			Duration INT,
			Timestamp INT UNSIGNED NOT NULL,
			SystemAction BOOLEAN DEFAULT FALSE,
			
			PRIMARY KEY (InfractionID),
			FOREIGN KEY (PlayerID) REFERENCES Players(PlayerID),
			FOREIGN KEY (UserID) REFERENCES Users(UserID),
			FOREIGN KEY (ServerID) REFERENCES Servers(ServerID)
		);
	`); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}

		return fmt.Errorf("could not create Infractions table. Error: %v", err)
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
