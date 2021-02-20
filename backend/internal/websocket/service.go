package websocket

import (
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/websocket"
	"github.com/sniddunc/refractor/refractor"
	"net"
)

type websocketService struct {
	pool *websocket.Pool
	log  log.Logger
}

func NewWebsocketService(log log.Logger) refractor.WebsocketService {
	return &websocketService{
		pool: websocket.NewPool(log),
		log:  log,
	}
}

func (s *websocketService) Broadcast(message *refractor.WebsocketMessage) {
	s.pool.Broadcast <- message
}

func (s *websocketService) CreateClient(userID int64, conn net.Conn) {
	client := websocket.NewClient(userID, conn, s.pool, s.log)

	s.pool.Register <- client
	client.Read()
}

func (s *websocketService) StartPool() {
	s.pool.Start()
}
