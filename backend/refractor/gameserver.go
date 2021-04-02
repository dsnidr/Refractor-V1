package refractor

import "github.com/labstack/echo/v4"

type GameServer struct {
	Name    string            `json:"name"`
	Config  *GameServerConfig `json:"config"`
	Servers []*ServerInfo     `json:"servers"`
}

// GameServerConfig holds any GameConfig fields which should be sent back to the client.
type GameServerConfig struct {
	EnableChat bool `json:"enableChat"`
}

type GameServerService interface {
	GetGameServers() ([]*GameServer, *ServiceResponse)
}

type GameServerHandler interface {
	GetAllGameServers(c echo.Context) error
}
