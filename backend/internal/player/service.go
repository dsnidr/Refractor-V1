package player

import (
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"time"
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
	if err := s.repo.Create(newPlayer); err != nil {
		s.log.Error("Could not create a new player. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	return newPlayer, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Player created",
	}
}

func (s *playerService) GetPlayerByID(id int64) (*refractor.Player, *refractor.ServiceResponse) {
	player, err := s.repo.FindByID(id)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not get player by ID. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	return player, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Player fetched",
	}
}

func (s *playerService) OnPlayerJoin(serverID int64, playerGameID string, currentName string) (*refractor.Player, *refractor.ServiceResponse) {
	// TODO: Differentiate between playerGameID per game. An additional method on the Game interface will likely be required.
	// For now, we just assume that this is Mordhau and work via PlayFabIDs.
	// Check if the player is recorded in storage
	foundPlayer, err := s.repo.FindByPlayFabID(playerGameID)
	if err != nil && err != refractor.ErrNotFound {
		// If there is an error and it isn't an instance of ErrNotFound, an actual error occurred that we should
		// log for traceability.
		s.log.Error("Could not check if player exists. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	// If foundPlayer == nil we know they don't exist, so we record them in storage
	if foundPlayer == nil {
		newPlayer, _ := s.CreatePlayer(&refractor.Player{
			PlayFabID:   playerGameID,
			CurrentName: currentName,
			LastSeen:    time.Now().Unix(),
		})

		if newPlayer == nil {
			s.log.Error("Could not create new player")
			return nil, refractor.InternalErrorResponse
		}

		return newPlayer, &refractor.ServiceResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Message:    "New player created",
		}
	}

	// If they player was already in storage, check if their name changed.
	if foundPlayer.CurrentName != currentName {
		if err := s.repo.UpdateName(foundPlayer, currentName); err != nil {
			s.log.Error("Could not update player name to new name. Error: %v", err)
			return nil, refractor.InternalErrorResponse
		}
	}

	return foundPlayer, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Player is already in storage",
	}
}

func (s *playerService) OnPlayerQuit(serverID int64, playerGameID string) (*refractor.Player, *refractor.ServiceResponse) {
	// TODO: Differentiate between playerGameID per game. An additional method on the Game interface will likely be required.
	// For now, we just assume that this is Mordhau and work via PlayFabIDs.
	// Check if the player is recorded in storage
	foundPlayer, err := s.repo.FindByPlayFabID(playerGameID)
	if err != nil && err != refractor.ErrNotFound {
		// If there is an error and it isn't an instance of ErrNotFound, an actual error occurred that we should
		// log for traceability.
		s.log.Error("Could not get player by PlayFabID from repo. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	if foundPlayer == nil {
		s.log.Error("foundPlayer with PlayFabID: %s was nil", playerGameID)
		return nil, refractor.InternalErrorResponse
	}

	// Update player's last seen field
	if _, err := s.repo.Update(foundPlayer.PlayerID, refractor.UpdateArgs{
		"LastSeen": time.Now().Unix(),
	}); err != nil {
		s.log.Error("Could not update LastSeen field for player with PlayFabID: %s. Error: %v", playerGameID, err)
		return nil, refractor.InternalErrorResponse
	}

	return foundPlayer, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Quit handled",
	}
}

func (s *playerService) GetPlayer(args refractor.FindArgs) (*refractor.Player, *refractor.ServiceResponse) {
	foundPlayer, err := s.repo.FindOne(args)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "No player found",
			}
		}

		s.log.Error("Could not find one player from storage. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	return foundPlayer, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Player found",
	}
}
