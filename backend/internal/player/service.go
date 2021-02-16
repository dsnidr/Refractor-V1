package player

import (
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
)

type playerService struct {
	repo refractor.PlayerRepository
	log  log.Logger
}

func NewPlayerService(repo refractor.PlayerRepository, log log.Logger) refractor.PlayerService {
	return &playerService{
		repo: repo,
		log:  log,
	}
}

func (s *playerService) CreatePlayer(newPlayer *refractor.Player) (*refractor.Player, *refractor.ServiceResponse) {
	panic("implement me")
}
