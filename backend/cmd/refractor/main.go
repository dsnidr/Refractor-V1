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

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sniddunc/refractor/internal/auth"
	"github.com/sniddunc/refractor/internal/chat"
	"github.com/sniddunc/refractor/internal/game"
	"github.com/sniddunc/refractor/internal/game/minecraft"
	"github.com/sniddunc/refractor/internal/game/mordhau"
	"github.com/sniddunc/refractor/internal/gameserver"
	"github.com/sniddunc/refractor/internal/http/api"
	"github.com/sniddunc/refractor/internal/infraction"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/internal/player"
	"github.com/sniddunc/refractor/internal/playerinfraction"
	"github.com/sniddunc/refractor/internal/rcon"
	"github.com/sniddunc/refractor/internal/search"
	"github.com/sniddunc/refractor/internal/server"
	"github.com/sniddunc/refractor/internal/storage/mysql"
	"github.com/sniddunc/refractor/internal/summary"
	"github.com/sniddunc/refractor/internal/user"
	"github.com/sniddunc/refractor/internal/watchdog"
	"github.com/sniddunc/refractor/internal/websocket"
	"github.com/sniddunc/refractor/pkg/env"
	logger "github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/perms"
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
	gameService.AddGame(minecraft.NewMinecraftGame())

	userRepo := mysql.NewUserRepository(db)
	userService := user.NewUserService(userRepo, loggerInst)
	userHandler := api.NewUserHandler(userService)

	authService := auth.NewAuthService(userRepo, loggerInst, os.Getenv("JWT_SECRET"))
	authHandler := api.NewAuthHandler(authService, secureMode)

	playerRepo := mysql.NewPlayerRepository(db)
	playerService := player.NewPlayerService(playerRepo, loggerInst)
	playerHandler := api.NewPlayerHandler(playerService)

	infractionRepo := mysql.NewInfractionRepository(db)
	playerInfractionService := playerinfraction.NewPlayerInfractionService(playerRepo, infractionRepo, loggerInst)

	serverRepo := mysql.NewServerRepository(db)
	serverService := server.NewServerService(serverRepo, gameService, playerInfractionService, loggerInst)
	serverHandler := api.NewServerHandler(serverService, playerService, loggerInst)
	playerService.SubscribeUpdate(serverService.OnPlayerUpdate)

	gameServerService := gameserver.NewGameServerService(gameService, serverService, loggerInst)
	gameServerHandler := api.NewGameServerHandler(gameServerService)

	websocketService := websocket.NewWebsocketService(playerService, userService, playerInfractionService, loggerInst)
	go websocketService.StartPool()

	rconService := rcon.NewRCONService(gameService, playerService, loggerInst)
	rconService.SubscribeJoin(playerHandler.OnPlayerJoin)
	rconService.SubscribeQuit(playerHandler.OnPlayerQuit)
	rconService.SubscribeJoin(websocketService.OnPlayerJoin)
	rconService.SubscribeQuit(websocketService.OnPlayerQuit)
	rconService.SubscribeJoin(serverHandler.OnPlayerJoin)
	rconService.SubscribeQuit(serverHandler.OnPlayerQuit)
	rconService.SubscribeOnline(serverService.OnServerOnline)
	rconService.SubscribeOffline(serverService.OnServerOffline)
	rconService.SubscribeOnline(websocketService.OnServerOnline)
	rconService.SubscribeOffline(websocketService.OnServerOffline)
	rconService.SubscribePlayerListPoll(serverService.OnPlayerListUpdate)

	chatService := chat.NewChatService(websocketService, rconService, loggerInst)
	rconService.SubscribeChat(chatService.OnChatReceive)
	websocketService.SubscribeChatSend(rconService.SendChatMessage)
	websocketService.SubscribeChatSend(chatService.OnUserSendChat)

	infractionService := infraction.NewInfractionService(infractionRepo, playerService, serverService, userService, loggerInst)
	infractionHandler := api.NewInfractionHandler(infractionService)

	summaryService := summary.NewSummaryService(playerService, infractionService, loggerInst)
	summaryHandler := api.NewSummaryHandler(summaryService)

	searchService := search.NewSearchService(playerRepo, infractionRepo, loggerInst)
	searchHandler := api.NewSearchHandler(searchService)

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

	// Start RCON client watchdog
	go watchdog.StartRCONServerWatchdog(rconService, serverService, loggerInst)

	// API Setup
	apiHandlers := &api.Handlers{
		AuthHandler:       authHandler,
		UserHandler:       userHandler,
		ServerHandler:     serverHandler,
		PlayerHandler:     playerHandler,
		GameServerHandler: gameServerHandler,
		InfractionHandler: infractionHandler,
		SummaryHandler:    summaryHandler,
		SearchHandler:     searchHandler,
	}

	// Done. Begin serving.
	loggerInst.Info("Refractor startup complete!")

	API := api.NewAPI(apiHandlers, port, loggerInst, websocketService)
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
	_, res = userService.SetUserPermissions(params.SetUserPermissionsParams{
		UserID:      newUser.UserID,
		Permissions: perms.SUPER_ADMIN,
		UserMeta: &params.UserMeta{
			UserID:      0,
			Permissions: perms.SUPER_ADMIN,
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
		serverService.CreateServerData(server.ServerID, server.Game)

		if err := rconService.CreateClient(server); err != nil {
			log.Info("Could not connect RCON client to server: %s", server.Name)
			continue
		}

		log.Info("RCON Client connected to %s", server.Name)
	}

	return nil
}
