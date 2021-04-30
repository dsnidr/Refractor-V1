package playerinfraction

import (
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
)

type playerInfractionService struct {
	playerRepo refractor.PlayerRepository
	infractionRepo refractor.InfractionRepository
	log         log.Logger
}

func NewPlayerInfractionService(playerRepo refractor.PlayerRepository,
	infractionRepo refractor.InfractionRepository, log log.Logger) refractor.PlayerInfractionService {
	return &playerInfractionService{
		playerRepo: playerRepo,
		infractionRepo: infractionRepo,
		log: log,
	}
}

func (s *playerInfractionService) GetPlayerInfractionCount(playerID int64) (int, *refractor.ServiceResponse) {
	count, err := s.infractionRepo.GetCountByPlayerID(playerID)
	if err != nil {
		s.log.Error("Could not get player infraction count. Error: %v", err)
		return 0, refractor.InternalErrorResponse
	}

	return count, &refractor.ServiceResponse{
		Success:          true,
		StatusCode:       http.StatusOK,
		Message:          "Fetched infraction count",
	}
}