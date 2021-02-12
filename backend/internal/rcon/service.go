package rcon

import (
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
)

type rconService struct {
	log log.Logger
}

func NewRCONService(log log.Logger) refractor.RCONService {
	return &rconService{
		log: log,
	}
}

func (s *rconService) CreateClient(server *refractor.Server) error {
	panic("implement me")
}

func (s *rconService) GetClients() map[int64]*refractor.RCONClient {
	panic("implement me")
}

func (s *rconService) DeleteClient(serverID int64) {
	panic("implement me")
}
