package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sniddunc/refractor/internal/storage/mysql"
	"github.com/sniddunc/refractor/pkg/env"
	logger "github.com/sniddunc/refractor/pkg/log"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load("./.env"); err == nil {
		log.Println("Environment variables loaded from .env file")
	}

	if err := env.RequireEnv("DB_URI").
		RequireEnv("JWT_SECRET").
		RequireEnv("DB_URI").
		GetError(); err != nil {
		log.Fatal(err)
	}

	// Setup loggerInst
	loggerInst, err := logger.NewLogger(true, true)
	if err != nil {
		log.Fatalf("Could not set up loggerInst. Error: %v", err)
	}

	// Setup database
	_, err = setupDatabase(os.Getenv("DB_URI"))
	if err != nil {
		log.Fatalf("Could not setup database. Error: %v", err)
	}

	loggerInst.Info("Refractor startup complete!")
}

func setupDatabase(connString string) (*sql.DB, error) {
	// for now, just assume we're using mysql since it's the only database type supported on initial release.
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err := mysql.Setup(db); err != nil {
		return nil, err
	}

	return db, nil
}
