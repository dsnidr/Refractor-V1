package search

import (
	"fmt"
	"github.com/sniddunc/refractor/internal/params"
	logger "github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
)

type searchService struct {
	playerRepo refractor.PlayerRepository
	log        logger.Logger
}

func NewSearchService(playerRepo refractor.PlayerRepository, log logger.Logger) refractor.SearchService {
	return &searchService{
		playerRepo: playerRepo,
		log:        log,
	}
}

func (s *searchService) SearchPlayers(body params.SearchPlayersParams) (int, []*refractor.Player, *refractor.ServiceResponse) {
	switch body.SearchType {
	case "playfabid":
		return s.searchByPlayerPlayFabID(body.SearchTerm)
	case "mcuuid":
		return s.searchByPlayerMCUUID(body.SearchTerm)
	case "name":
		return s.searchByPlayerName(body.SearchTerm, body.SearchParams.Limit, body.SearchParams.Offset)
	default:
		s.log.Warn("Attempted to search an invalid search type: %s", body.SearchType)
		return 0, []*refractor.Player{}, refractor.InternalErrorResponse
	}
}

func (s *searchService) searchByPlayerPlayFabID(playFabID string) (int, []*refractor.Player, *refractor.ServiceResponse) {
	player, err := s.playerRepo.FindByPlayFabID(playFabID)
	if err != nil {
		if err == refractor.ErrNotFound {
			return 0, []*refractor.Player{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Found 0 matching players",
			}
		}

		s.log.Error("Could not get player by PlayFabID. Error: %v", err)
		return 0, nil, refractor.InternalErrorResponse
	}

	return 1, []*refractor.Player{player}, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Found 1 matching players",
	}
}

func (s *searchService) searchByPlayerMCUUID(MCUUID string) (int, []*refractor.Player, *refractor.ServiceResponse) {
	player, err := s.playerRepo.FindByMCUUID(MCUUID)
	if err != nil {
		if err == refractor.ErrNotFound {
			return 0, []*refractor.Player{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Found 0 matching players",
			}
		}

		s.log.Error("Could not get player by MCUUID. Error: %v", err)
		return 0, nil, refractor.InternalErrorResponse
	}

	return 1, []*refractor.Player{player}, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Found 1 matching players",
	}
}

func (s *searchService) searchByPlayerName(name string, limit int, offset int) (int, []*refractor.Player, *refractor.ServiceResponse) {
	count, players, err := s.playerRepo.SearchByName(name, limit, offset)
	if err != nil {
		if err == refractor.ErrNotFound {
			return 0, []*refractor.Player{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Found 0 matching players",
			}
		}

		s.log.Error("Could not search players by name. Error: %v", err)
		return 0, nil, refractor.InternalErrorResponse
	}

	return count, players, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Found %d matching players", len(players)),
	}
}

func (s *searchService) SearchInfractions(body params.SearchInfractionsParams) (int, []*refractor.Infraction, *refractor.ServiceResponse) {
	return 0, nil, nil
}