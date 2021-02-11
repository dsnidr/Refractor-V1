package refractor

type GameServer struct {
	Servers []*Server
	Game
}

type GameServerService interface {
	GetGameServers() ([]*GameServer, *ServiceResponse)
}
