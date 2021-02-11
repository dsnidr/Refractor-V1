package refractor

import "github.com/labstack/echo/v4"

type GameServer struct {
	Game    string        `json:"game"`
	Servers []*ServerInfo `json:"servers"`
}

type GameServerService interface {
	GetGameServers() ([]*GameServer, *ServiceResponse)
}

type GameServerHandler interface {
	GetAllGameServers(c echo.Context) error
}
