/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package search

import (
	"fmt"
	"github.com/sniddunc/refractor/internal/params"
	logger "github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"strconv"
	"time"
)

type searchService struct {
	playerRepo     refractor.PlayerRepository
	infractionRepo refractor.InfractionRepository
	chatRepo       refractor.ChatRepository
	log            logger.Logger
}

func NewSearchService(playerRepo refractor.PlayerRepository, infractionRepo refractor.InfractionRepository,
	chatRepo refractor.ChatRepository, log logger.Logger) refractor.SearchService {
	return &searchService{
		playerRepo:     playerRepo,
		infractionRepo: infractionRepo,
		chatRepo:       chatRepo,
		log:            log,
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
	case "id":
		return s.searchByID(body.SearchTerm)
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

func (s *searchService) searchByID(idString string) (int, []*refractor.Player, *refractor.ServiceResponse) {
	// Parse ID
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return 0, []*refractor.Player{}, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			ValidationErrors: map[string][]string{
				"term": {"Invalid player ID"},
			},
		}
	}

	player, err := s.playerRepo.FindByID(id)
	if err != nil {
		if err == refractor.ErrNotFound {
			return 0, []*refractor.Player{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Found 0 matching players",
			}
		}

		s.log.Error("Could not get player by id. Error: %v", err)
		return 0, nil, refractor.InternalErrorResponse
	}

	return 1, []*refractor.Player{player}, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Found 1 matching players",
	}
}

func (s *searchService) SearchInfractions(body params.SearchInfractionsParams) (int, []*refractor.Infraction, *refractor.ServiceResponse) {
	searchArgs := refractor.FindArgs{}

	// add defined arguments from body into searchArgs
	if body.Type != "" {
		searchArgs["Type"] = body.Type
	}

	if body.Game != "" {
		searchArgs["Game"] = body.Game
	}

	if body.ParsedInfractionIDs.PlayerID != 0 {
		searchArgs["PlayerID"] = body.ParsedInfractionIDs.PlayerID
	}

	if body.ParsedInfractionIDs.ServerID != 0 {
		searchArgs["ServerID"] = body.ParsedInfractionIDs.ServerID
	}

	if body.ParsedInfractionIDs.UserID != 0 {
		searchArgs["UserID"] = body.ParsedInfractionIDs.UserID
	}

	if len(searchArgs) == 0 {
		return 0, []*refractor.Infraction{}, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "You must provide at least one search filter",
		}
	}

	// Execute search
	count, infractions, err := s.infractionRepo.Search(searchArgs, body.SearchParams.Limit, body.SearchParams.Offset)
	if err != nil {
		if err == refractor.ErrNotFound {
			return 0, []*refractor.Infraction{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Found 0 total results",
			}
		}

		s.log.Error("Could not search infractions. Error: %v", err)
		return 0, []*refractor.Infraction{}, refractor.InternalErrorResponse
	}

	return count, infractions, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Found %d total results", count),
	}
}

func (s *searchService) SearchChatMessages(body params.SearchChatMessagesParams) (int, []*refractor.ChatMessage, *refractor.ServiceResponse) {
	searchArgs := refractor.FindArgs{}

	if body.ParsedChatMessageIDs.PlayerID != 0 {
		searchArgs["PlayerID"] = body.ParsedChatMessageIDs.PlayerID
	}

	if body.ParsedChatMessageIDs.ServerID != 0 {
		searchArgs["ServerID"] = body.ParsedChatMessageIDs.ServerID
	}

	searchArgs["Message"] = body.Message
	searchArgs["StartDate"] = body.StartDate

	if body.EndDate == 0 {
		body.EndDate = time.Now().Unix()
	}
	searchArgs["EndDate"] = body.EndDate

	// Run search
	count, results, err := s.chatRepo.Search(searchArgs, body.SearchParams.Limit, body.SearchParams.Offset)
	if err != nil {
		if err == refractor.ErrNotFound {
			return 0, []*refractor.ChatMessage{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Found 0 total results",
			}
		}

		s.log.Error("Could not search chat messages. Error: %v", err)
		return 0, []*refractor.ChatMessage{}, refractor.InternalErrorResponse
	}

	return count, results, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Found %d total results", count),
	}
}
