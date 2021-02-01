package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/internal/storage/mysql"
	"github.com/sniddunc/refractor/internal/user"
	"github.com/sniddunc/refractor/pkg/env"
	logger "github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load("./.env"); err == nil {
		fmt.Println("Environment variables loaded from .env file")
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
	db, err := setupDatabase(os.Getenv("DB_URI"))
	if err != nil {
		log.Fatalf("Could not setup database. Error: %v", err)
	}

	// Set up application components
	userRepo := mysql.NewUserRepository(db)
	userService := user.NewUserService(userRepo, loggerInst)

	// Set up initial user if no users currently exist
	if count := userRepo.GetCount(); count == 0 {
		if err := setupInitialUser(userService); err != nil {
			log.Fatalf("Could not create initial user. Error: %v", err)
		}

		loggerInst.Info("Initial user created from environment variables")
	}

	// Done. Begin serving.
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

func setupInitialUser(userService refractor.UserService) error {
	if err := env.RequireEnv("INITIAL_USER_USERNAME").
		RequireEnv("INITIAL_USER_EMAIL").
		RequireEnv("INITIAL_USER_PASSWORD").
		GetError(); err != nil {
		return err
	}

	body := params.CreateUserParams{
		Email:           os.Getenv("INITIAL_USER_EMAIL"),
		Username:        os.Getenv("INITIAL_USER_USERNAME"),
		Password:        os.Getenv("INITIAL_USER_PASSWORD"),
		PasswordConfirm: os.Getenv("INITIAL_USER_PASSWORD"),
	}

	if ok, errors := body.Validate(); !ok {
		return fmt.Errorf("validation errors occurred: %v", errors)
	}

	// Create user
	_, res := userService.CreateUser(body)
	if !res.Success {
		return fmt.Errorf("could not create initial user: %s", res.Message)
	}

	// TODO: Update user to make them a super-admin

	return nil
}
