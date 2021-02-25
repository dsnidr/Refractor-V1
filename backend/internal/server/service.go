package server

import (
	"fmt"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"net/url"
	"reflect"
)

type serverService struct {
	repo        refractor.ServerRepository
	gameService refractor.GameService
	log         log.Logger
	serverData  map[int64]*refractor.ServerData
}

func NewServerService(repo refractor.ServerRepository, gameService refractor.GameService, log log.Logger) refractor.ServerService {
	return &serverService{
		repo:        repo,
		gameService: gameService,
		log:         log,
		serverData:  map[int64]*refractor.ServerData{},
	}
}

func (s *serverService) CreateServer(body params.CreateServerParams) (*refractor.Server, *refractor.ServiceResponse) {
	// Check if the game is valid
	gameExists, res := s.gameService.GameExists(body.Game)
	if !res.Success {
		return nil, refractor.InternalErrorResponse
	}

	if !gameExists {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			ValidationErrors: url.Values{
				"game": []string{"Invalid game"},
			},
		}
	}

	// Check if a server with this name exists
	args := refractor.FindArgs{
		"Name": body.Name,
	}

	exists, err := s.repo.Exists(args)
	if err != nil {
		s.log.Error("Could not check existence of server. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	if exists {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			ValidationErrors: url.Values{
				"name": []string{"A server with this name already exists"},
			},
		}
	}

	// Create the new server
	newServer := &refractor.Server{
		Game:         body.Game,
		Name:         body.Name,
		Address:      body.Address,
		RCONPort:     body.RCONPort,
		RCONPassword: body.RCONPassword,
	}

	if err := s.repo.Create(newServer); err != nil {
		s.log.Error("Could not insert new server into repository. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	return newServer, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Server created",
	}
}

func (s *serverService) CreateServerData(id int64, gameName string) {
	s.serverData[id] = &refractor.ServerData{
		NeedsUpdate:   true,
		ServerID:      id,
		Game:          gameName,
		PlayerCount:   0,
		OnlinePlayers: map[string]*refractor.Player{},
	}
}

func (s *serverService) GetAllServers() ([]*refractor.Server, *refractor.ServiceResponse) {
	servers, err := s.repo.FindAll()
	if err != nil {
		if err == refractor.ErrNotFound {
			return []*refractor.Server{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Fetched 0 servers",
			}
		}

		s.log.Error("Could not FindAll servers from repository. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	return servers, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Fetched %d servers", len(servers)),
	}
}

func (s *serverService) createServerData(id int64, game string) {
	s.serverData[id] = &refractor.ServerData{
		NeedsUpdate: true,
		ServerID:    id,
		Game:        game,
		Online:      false,
	}
}

func (s *serverService) GetAllServerData() ([]*refractor.ServerData, *refractor.ServiceResponse) {
	var allServerData []*refractor.ServerData

	for _, serverData := range s.serverData {
		allServerData = append(allServerData, serverData)
	}

	return allServerData, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Fetched server data for %d servers", len(allServerData)),
	}
}

func (s *serverService) GetServerData(id int64) (*refractor.ServerData, *refractor.ServiceResponse) {
	return s.serverData[id], &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
	}
}

func (s *serverService) GetServerByID(id int64) (*refractor.Server, *refractor.ServiceResponse) {
	server, err := s.repo.FindByID(id)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not FindByID server from repository. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	return server, &refractor.ServiceResponse{
		Success: true,
		StatusCode: http.StatusOK,
		Message: "Server fetched",
	}
}

func (s *serverService) OnPlayerJoin(serverID int64, player *refractor.Player) {
	// Get the game for this server
	game, _ := s.gameService.GetGame(s.serverData[serverID].Game)

	// Use reflection to get the proper PlayerGameIDField from the player
	r := reflect.ValueOf(player)
	field := reflect.Indirect(r).FieldByName(game.GetConfig().PlayerGameIDField).String()

	// Add the player to the server data
	s.serverData[serverID].OnlinePlayers[field] = player
}

func (s *serverService) OnPlayerQuit(serverID int64, player *refractor.Player) {
	// Get the game for this server
	game, _ := s.gameService.GetGame(s.serverData[serverID].Game)

	// Use reflection to get the proper PlayerGameIDField from the player
	r := reflect.ValueOf(player)
	field := reflect.Indirect(r).FieldByName(game.GetConfig().PlayerGameIDField).String()

	// Remove the player from the server data
	delete(s.serverData[serverID].OnlinePlayers, field)
}

func (s *serverService) OnServerOnline(serverID int64) {
	if s.serverData[serverID] == nil {
		s.log.Warn("OnServerOnline was called with an invalid serverID of %d", serverID)
		return
	}

	s.serverData[serverID].Online = true
}

func (s *serverService) OnServerOffline(serverID int64) {
	if s.serverData[serverID] == nil {
		s.log.Warn("OnServerOffline was called with an invalid serverID of %d", serverID)
		return
	}

	s.serverData[serverID].Online = false

	s.log.Warn("Server with ID %d has gone offline", serverID)
}
