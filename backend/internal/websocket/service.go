package websocket

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/websocket"
	"github.com/sniddunc/refractor/refractor"
	"net"
)

type websocketService struct {
	pool                *websocket.Pool
	userService         refractor.UserService
	playerService       refractor.PlayerService
	log                 log.Logger
	chatSendSubscribers []refractor.ChatSendSubscriber
}

func NewWebsocketService(playerService refractor.PlayerService, userService refractor.UserService, log log.Logger) refractor.WebsocketService {
	return &websocketService{
		pool:                websocket.NewPool(log),
		playerService:       playerService,
		userService:         userService,
		log:                 log,
		chatSendSubscribers: []refractor.ChatSendSubscriber{},
	}
}

func (s *websocketService) Broadcast(message *refractor.WebsocketMessage) {
	s.pool.Broadcast <- message
}

func (s *websocketService) CreateClient(userID int64, conn net.Conn) {
	client := websocket.NewClient(userID, conn, s.pool, s.log, s.sendChatHandler)

	s.pool.Register <- client
	client.Read()
}

func (s *websocketService) sendChatHandler(msgBody *websocket.SendChatBody) {
	// Get user's name
	user, res := s.userService.GetUserByID(msgBody.UserID)
	if user == nil {
		s.log.Error("Could get get user ID. Res message: %s", res.Message)
		return
	}

	transformed := &refractor.ChatSendBody{
		ServerID: msgBody.ServerID,
		Message:  msgBody.Message,
		Sender:   user.Username,
	}

	for _, sub := range s.chatSendSubscribers {
		sub(transformed)
	}
}

func (s *websocketService) StartPool() {
	s.pool.Start()
}

type playerJoinQuitData struct {
	ServerID     int64  `json:"serverId"`
	PlayerID     int64  `json:"id"`
	PlayerGameID string `json:"playerGameId"`
	Name         string `json:"name"`
}

func (s *websocketService) OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	idField := gameConfig.PlayerGameIDField

	player, res := s.playerService.GetPlayer(refractor.FindArgs{
		idField: fields[idField],
	})

	if !res.Success {
		s.log.Warn("Could not GetPlayer. PlayerGameIDField = %s, field value = %v", idField, fields[idField])
		return
	}

	s.Broadcast(&refractor.WebsocketMessage{
		Type: "player-join",
		Body: playerJoinQuitData{
			ServerID:     serverID,
			PlayerID:     player.PlayerID,
			PlayerGameID: fields[idField],
			Name:         player.CurrentName,
		},
	})
}

func (s *websocketService) OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	idField := gameConfig.PlayerGameIDField

	player, res := s.playerService.GetPlayer(refractor.FindArgs{
		idField: fields[idField],
	})

	if !res.Success {
		s.log.Warn("Could not GetPlayer. PlayerGameIDField = %s, field value = %v", idField, fields[idField])
		return
	}

	if player == nil {
		s.log.Warn("GetPlayer player returned was nil!")
		return
	}

	s.Broadcast(&refractor.WebsocketMessage{
		Type: "player-quit",
		Body: playerJoinQuitData{
			ServerID:     serverID,
			PlayerID:     player.PlayerID,
			PlayerGameID: fields[idField],
			Name:         player.CurrentName,
		},
	})
}

func (s *websocketService) OnServerOnline(serverID int64) {
	s.Broadcast(&refractor.WebsocketMessage{
		Type: "server-online",
		Body: serverID,
	})
}

func (s *websocketService) OnServerOffline(serverID int64) {
	s.Broadcast(&refractor.WebsocketMessage{
		Type: "server-offline",
		Body: serverID,
	})
}

func (s *websocketService) SubscribeChatSend(subscriber refractor.ChatSendSubscriber) {
	s.chatSendSubscribers = append(s.chatSendSubscribers, subscriber)
}
