package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sniddunc/refractor/internal/auth"
	"github.com/sniddunc/refractor/internal/game"
	"github.com/sniddunc/refractor/internal/game/mordhau"
	"github.com/sniddunc/refractor/internal/gameserver"
	"github.com/sniddunc/refractor/internal/http/api"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/internal/rcon"
	"github.com/sniddunc/refractor/internal/server"
	"github.com/sniddunc/refractor/internal/storage/mysql"
	"github.com/sniddunc/refractor/internal/user"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/env"
	logger "github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"log"
	"math"
	"os"
)

func main() {
	if err := godotenv.Load("./.env"); err == nil {
		fmt.Println("Environment variables loaded from .env file")
	}

	if err := env.RequireEnv("DB_URI").
		RequireEnv("JWT_SECRET").
		RequireEnv("DB_URI").
		RequireEnv("GAME").
		GetError(); err != nil {
		log.Fatal(err)
	}

	secureMode := os.Getenv("SECURE") == "true"

	// Get port if defined
	var port string = ":5000"
	if portVal := os.Getenv("PORT"); portVal != "" {
		port = fmt.Sprintf(":%s", portVal)
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
	gameService := game.NewGameService()
	gameService.AddGame(mordhau.NewMordhauGame())

	userRepo := mysql.NewUserRepository(db)
	userService := user.NewUserService(userRepo, loggerInst)
	userHandler := api.NewUserHandler(userService)

	authService := auth.NewAuthService(userRepo, loggerInst, os.Getenv("JWT_SECRET"))
	authHandler := api.NewAuthHandler(authService, secureMode)

	serverRepo := mysql.NewServerRepository(db)
	serverService := server.NewServerService(serverRepo, gameService, loggerInst)
	serverHandler := api.NewServerHandler(serverService, loggerInst)

	gameServerService := gameserver.NewGameServerService(gameService, serverService, loggerInst)
	gameServerHandler := api.NewGameServerHandler(gameServerService)

	rconService := rcon.NewRCONService(gameService, loggerInst)

	// Set up initial user if no users currently exist
	if count := userRepo.GetCount(); count == 0 {
		if err := setupInitialUser(userService); err != nil {
			log.Fatalf("Could not create initial user. Error: %v", err)
		}

		loggerInst.Info("Initial user created from environment variables")
	}

	// Set up RCON clients for all existing servers
	if err := setupServerClients(rconService, serverService, loggerInst); err != nil {
		log.Fatalf("Could not set up server RCON clients. Error: %v", err)
	}

	// API Setup
	apiHandlers := &api.Handlers{
		AuthHandler:       authHandler,
		UserHandler:       userHandler,
		ServerHandler:     serverHandler,
		GameServerHandler: gameServerHandler,
	}

	// Done. Begin serving.
	loggerInst.Info("Refractor startup complete!")

	API := api.NewAPI(apiHandlers, port, loggerInst)
	if err := API.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
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
	newUser, res := userService.CreateUser(body)
	if !res.Success {
		return fmt.Errorf("could not create initial user: %s", res.Message)
	}

	// Make user a super-admin
	_, res = userService.SetUserAccessLevel(params.SetUserAccessLevelParams{
		UserID:      newUser.UserID,
		AccessLevel: config.AL_SUPERADMIN,
		UserMeta: &params.UserMeta{
			UserID:      0,
			AccessLevel: math.MaxInt32,
		},
	})

	if !res.Success {
		return fmt.Errorf("could not set new user to be a superadmin: %s", res.Message)
	}

	return nil
}

func setupServerClients(rconService refractor.RCONService, serverService refractor.ServerService, log logger.Logger) error {
	allServers, res := serverService.GetAllServers()
	if !res.Success {
		return fmt.Errorf("could not get all servers. Message: %s", res.Message)
	}

	for _, server := range allServers {
		serverService.CreateServerData(server.ServerID)

		if err := rconService.CreateClient(server); err != nil {
			return err
		}

		log.Info("RCON Client connected to", server.Name)
	}

	return nil
}
