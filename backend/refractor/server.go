package refractor

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
)

type Server struct {
	ServerID     int64  `json:"id"`
	Name         string `json:"name"`
	Game         string `json:"game"`
	Address      string `json:"address"`
	RCONPort     string `json:"rconPort"`
	RCONPassword string `json:"rconPassword"`
}

type ServerInfo struct {
	ServerID int64  `json:"id"`
	Name     string `json:"name"`
	Game     string `json:"game"`
	Address  string `json:"address"`
}

type ServerRepository interface {
	Create(server *Server) error
	FindByID(id int64) (*Server, error)
	Exists(args FindArgs) (bool, error)
	FindOne(args FindArgs) (*Server, error)
	FindAll() ([]*Server, error)
	Update(id int64, args UpdateArgs) (*Server, error)
	Delete(id int64) error
}

type ServerService interface {
	CreateServer(body params.CreateServerParams) (*Server, *ServiceResponse)
	GetAllServers() ([]*Server, *ServiceResponse)
}

type ServerHandler interface {
	CreateServer(c echo.Context) error
}
