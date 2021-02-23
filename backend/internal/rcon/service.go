package rcon

import (
	"fmt"
	rcon "github.com/sniddunc/mordhau-rcon"
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"strconv"
)

type rconService struct {
	clients            map[int64]*refractor.RCONClient
	gameService        refractor.GameService
	log                log.Logger
	joinSubscribers    []refractor.BroadcastSubscriber
	quitSubscribers    []refractor.BroadcastSubscriber
	onlineSubscribers  []refractor.StatusSubscriber
	offlineSubscribers []refractor.StatusSubscriber
}

func NewRCONService(gameService refractor.GameService, log log.Logger) refractor.RCONService {
	return &rconService{
		clients:            map[int64]*refractor.RCONClient{},
		gameService:        gameService,
		log:                log,
		joinSubscribers:    []refractor.BroadcastSubscriber{},
		quitSubscribers:    []refractor.BroadcastSubscriber{},
		onlineSubscribers:  []refractor.StatusSubscriber{},
		offlineSubscribers: []refractor.StatusSubscriber{},
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
		DisconnectHandler:        s.getDisconnectHandler(server.ServerID),
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

	// If this point was reached, we know the RCON connection was successful so we notify server online subscribers
	// of this server online event.
	for _, sub := range s.onlineSubscribers {
		sub(server.ServerID)
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

// SubscribeJoin adds a function to a slice of functions to be called when a player joins a server
func (s *rconService) SubscribeJoin(subscriber refractor.BroadcastSubscriber) {
	s.joinSubscribers = append(s.joinSubscribers, subscriber)
}

// SubscribeQuit adds a function to a slice of functions to be called when a player quits a server
func (s *rconService) SubscribeQuit(subscriber refractor.BroadcastSubscriber) {
	s.quitSubscribers = append(s.quitSubscribers, subscriber)
}

// SubscribeOnline adds a function to a slice of functions to be called when an RCON connection to a server comes online
func (s *rconService) SubscribeOnline(subscriber refractor.StatusSubscriber) {
	s.onlineSubscribers = append(s.onlineSubscribers, subscriber)
}

// SubscribeOffline adds a function to a slice of functions to be called when an RCON connection to a server goes offline
func (s *rconService) SubscribeOffline(subscriber refractor.StatusSubscriber) {
	s.offlineSubscribers = append(s.offlineSubscribers, subscriber)
}

func (s *rconService) getBroadcastListener(serverID int64, gameConfig *refractor.GameConfig) func(string) {
	// We wrap this in a parent function so that we can pass in the server IDs which each client belongs to.
	// This allows us to uniquely identify which server a broadcast came from.
	return func(message string) {
		s.log.Info("Received broadcast from server ID %d: %v", serverID, message)

		bcast := broadcast.GetBroadcastType(message, gameConfig.BroadcastPatterns)

		switch bcast.Type {
		case broadcast.TYPE_JOIN:
			s.HandleJoinBroadcast(bcast, serverID, gameConfig)
			break
		case broadcast.TYPE_QUIT:
			s.HandleQuitBroadcast(bcast, serverID, gameConfig)
		}
	}
}

func (s *rconService) getDisconnectHandler(serverID int64) func(error, bool) {
	return func(err error, expected bool) {
		// Notify all subscribers of a server offline event
		for _, sub := range s.offlineSubscribers {
			sub(serverID)
		}
	}
}
