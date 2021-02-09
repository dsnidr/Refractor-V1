package refractor

type Server struct {
	ServerID     int64  `json:"id"`
	Game         string `json:"game"`
	Address      string `json:"address"`
	RCONPort     string `json:"rconPort"`
	RCONPassword string `json:"rconPassword"`
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
}

type ServerHandler interface {
}
