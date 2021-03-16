package rcon

import (
	"fmt"
	rcon "github.com/sniddunc/mordhau-rcon"
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/regexutils"
	"github.com/sniddunc/refractor/refractor"
	"strconv"
	"time"
)

type rconService struct {
	clients            map[int64]*refractor.RCONClient
	gameService        refractor.GameService
	playerService      refractor.PlayerService
	log                log.Logger
	joinSubscribers    []refractor.BroadcastSubscriber
	quitSubscribers    []refractor.BroadcastSubscriber
	onlineSubscribers  []refractor.StatusSubscriber
	offlineSubscribers []refractor.StatusSubscriber

	// used to store players for future comparison if broadcasts are not enabled
	// prevPlayers[serverId][playerGameID] = onlinePlayer
	prevPlayers map[int64]map[string]*onlinePlayer
}

func NewRCONService(gameService refractor.GameService, playerService refractor.PlayerService, log log.Logger) refractor.RCONService {
	return &rconService{
		clients:            map[int64]*refractor.RCONClient{},
		gameService:        gameService,
		playerService:      playerService,
		log:                log,
		joinSubscribers:    []refractor.BroadcastSubscriber{},
		quitSubscribers:    []refractor.BroadcastSubscriber{},
		onlineSubscribers:  []refractor.StatusSubscriber{},
		offlineSubscribers: []refractor.StatusSubscriber{},
		prevPlayers:        map[int64]map[string]*onlinePlayer{},
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
	} else {
		go s.startPlayerListPolling(server.ServerID, game)
	}

	// Add to list of clients
	s.clients[server.ServerID] = &refractor.RCONClient{
		Server: server,
		Client: client,
	}

	// Get players currently on the server
	onlinePlayers := s.getOnlinePlayers(server.ServerID, game)

	for _, onlinePlayer := range onlinePlayers {
		for _, sub := range s.joinSubscribers {
			sub(broadcast.Fields{
				game.GetConfig().PlayerGameIDField: onlinePlayer.PlayerGameID,
				"Name":                             onlinePlayer.Name,
			}, server.ServerID, game.GetConfig())
		}
	}

	// If this point was reached, we know the RCON connection was successful so we notify server online subscribers
	// of this server online event.
	for _, sub := range s.onlineSubscribers {
		sub(server.ServerID)
	}

	s.log.Info("A new RCON client was created for server ID: %d", server.ServerID)

	return nil
}

func (s *rconService) startPlayerListPolling(serverID int64, game refractor.Game) {
	// Set up prevPlayers map for this server
	s.prevPlayers[serverID] = map[string]*onlinePlayer{}

	for {
		time.Sleep(game.GetConfig().PlayerListPollingInterval)

		client := s.clients[serverID]
		if client == nil {
			s.log.Warn("Player list polling routine could not get the client for server ID %d", serverID)
			s.log.Warn("Exiting player list polling routine for server ID %d", serverID)
			return
		}

		players := s.getOnlinePlayers(serverID, game)

		onlinePlayers := map[string]*onlinePlayer{}
		for _, player := range players {
			onlinePlayers[player.PlayerGameID] = player
		}

		prevPlayers := s.prevPlayers[serverID]

		// Check for new player joins
		for playerGameID, player := range onlinePlayers {
			if prevPlayers[playerGameID] == nil {
				s.log.Info("Player joined server ID %d: %s", serverID, player.Name)
				prevPlayers[playerGameID] = player

				// Player was not online previously so broadcast join
				for _, sub := range s.joinSubscribers {
					sub(broadcast.Fields{
						game.GetConfig().PlayerGameIDField: player.PlayerGameID,
						"Name":                             player.Name,
					}, serverID, game.GetConfig())
				}
			}
		}

		// Check for existing player quits
		for prevPlayerGameID, prevPlayer := range prevPlayers {
			if onlinePlayers[prevPlayerGameID] == nil {
				s.log.Info("Player left server ID %d: %s", serverID, prevPlayer.Name)
				delete(prevPlayers, prevPlayerGameID)

				// Player quit so broadcast quit
				for _, sub := range s.quitSubscribers {
					sub(broadcast.Fields{
						game.GetConfig().PlayerGameIDField: prevPlayer.PlayerGameID,
						"Name":                             prevPlayer.Name,
					}, serverID, game.GetConfig())
				}
			}
		}

		// Update prevPlayers for this server
		s.prevPlayers[serverID] = prevPlayers
	}
}

func (s *rconService) GetClients() map[int64]*refractor.RCONClient {
	return s.clients
}

func (s *rconService) DeleteClient(serverID int64) {
	delete(s.clients, serverID)
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

type onlinePlayer struct {
	PlayerGameID string
	Name         string
}

func (s *rconService) getOnlinePlayers(serverID int64, game refractor.Game) []*onlinePlayer {
	playerListCommand := game.GetPlayerListCommand()

	res, err := s.clients[serverID].ExecCommand(playerListCommand)
	if err != nil {
		s.log.Error("RCON ExecCommand %s failed with error: %v", playerListCommand, err)
		return nil
	}

	// Extract player info from RCON command response
	playerListPattern := game.GetConfig().CmdOutputPatterns["PlayerList"]
	players := playerListPattern.FindAllString(res, -1)

	var onlinePlayers []*onlinePlayer

	for _, player := range players {
		fields := regexutils.MapNamedMatches(playerListPattern, player)

		playerGameID := fields[game.GetConfig().PlayerGameIDField]
		name := fields["Name"]

		onlinePlayers = append(onlinePlayers, &onlinePlayer{
			PlayerGameID: playerGameID,
			Name:         name,
		})
	}

	return onlinePlayers
}
