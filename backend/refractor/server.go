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

// ServerData is used to keep track of server data (player counts, levels, etc)
type ServerData struct {
	NeedsUpdate   bool
	ServerID      int64
	Game          string
	Online        bool
	PlayerCount   int
	OnlinePlayers map[string]*Player
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
	CreateServerData(id int64, gameName string)
	GetAllServers() ([]*Server, *ServiceResponse)
	GetAllServerData() ([]*ServerData, *ServiceResponse)
	GetServerData(id int64) (*ServerData, *ServiceResponse)
	OnPlayerJoin(id int64, player *Player)
	OnPlayerQuit(id int64, player *Player)
	OnServerOnline(serverID int64)
	OnServerOffline(serverID int64)
}

type ServerHandler interface {
	CreateServer(c echo.Context) error
	GetAllServers(c echo.Context) error
}
