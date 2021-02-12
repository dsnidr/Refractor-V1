package refractor

import rcon "github.com/sniddunc/mordhau-rcon"

// RCONClient wraps around a mordhau-rcon Client and has an extra field containing the server
type RCONClient struct {
	Server *Server
	*rcon.Client
}

type RCONService interface {
	CreateClient(*Server) error
	GetClients() map[int64]*RCONClient
	DeleteClient(serverID int64)
}
