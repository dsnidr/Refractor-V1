package rcon

import (
	"fmt"
	rcon "github.com/sniddunc/mordhau-rcon"
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"strconv"
)

type BroadcastSubscriber func(fields broadcast.Fields, serverID int64)

type rconService struct {
	clients         map[int64]*refractor.RCONClient
	gameService     refractor.GameService
	log             log.Logger
	joinSubscribers []BroadcastSubscriber
	quitSubscribers []BroadcastSubscriber
}

func NewRCONService(gameService refractor.GameService, log log.Logger) refractor.RCONService {
	return &rconService{
		clients:         map[int64]*refractor.RCONClient{},
		gameService:     gameService,
		log:             log,
		joinSubscribers: []BroadcastSubscriber{},
		quitSubscribers: []BroadcastSubscriber{},
	}
}

func (s *rconService) CreateClient(server *refractor.Server) error {
	port, err := strconv.ParseInt(server.RCONPort, 10, 16)
	if err != nil {
		return err
	}

	// Get the server's game
	game, _ := s.gameService.GetGame(server.Game)
	if game == nil {
		return fmt.Errorf("invalid game for server ID %d: %s", server.ServerID, server.Game)
	}

	gameConfig := game.GetConfig()

	// Create client
	client := rcon.NewClient(&rcon.ClientConfig{
		Host:                     server.Address,
		Port:                     int16(port),
		Password:                 server.RCONPassword,
		SendHeartbeatCommand:     gameConfig.SendAlivePing,
		HeartbeatCommandInterval: gameConfig.AlivePingInterval,
		AttemptReconnect:         false,
		EnableBroadcasts:         gameConfig.EnableBroadcasts,
		BroadcastHandler:         s.getBroadcastListener(server.ServerID, gameConfig),
	})

	// Connect the main socket
	if err := client.Connect(); err != nil {
		return err
	}

	// Connect broadcast socket
	if gameConfig.EnableBroadcasts {
		errorChan := make(chan error)
		go client.ListenForBroadcasts([]string{"login"}, errorChan)

		go func() {
			select {
			case err := <-errorChan:
				s.log.Error("Broadcast listener error: %v\n", err)
				break
			}
		}()
	}

	// Add to list of clients
	s.clients[server.ServerID] = &refractor.RCONClient{
		Server: server,
		Client: client,
	}

	// TODO: Check if the client needs to update it's server's state. If so, do the update.

	// TODO: Once websockets are implemented, inform them of the server status.

	s.log.Info("A new RCON client was created for server ID: %d", server.ServerID)

	return nil
}

func (s *rconService) GetClients() map[int64]*refractor.RCONClient {
	panic("implement me")
}

func (s *rconService) DeleteClient(serverID int64) {
	panic("implement me")
}

func (s *rconService) getBroadcastListener(serverID int64, gameConfig *refractor.GameConfig) func(string) {
	// We wrap this in a parent function so that we can pass in the server IDs which each client belongs to.
	// This allows us to uniquely identify which server a broadcast came from.
	return func(message string) {
		s.log.Info("Received broadcast from server ID %d: %v", serverID, message)

		bcast := broadcast.GetBroadcastType(message, gameConfig.BroadcastPatterns)

		switch bcast.Type {
		case broadcast.TYPE_JOIN:
			s.HandleJoinBroadcast(bcast, serverID)
			break
		case broadcast.TYPE_QUIT:
			s.HandleQuitBroadcast(bcast, serverID)
		}
	}
}
