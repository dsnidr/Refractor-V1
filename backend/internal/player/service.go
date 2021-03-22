package player

import (
	"database/sql"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"reflect"
	"time"
)

type playerService struct {
	repo          refractor.PlayerRepository
	log           log.Logger
	recentPlayers *recentPlayers
}

func NewPlayerService(repo refractor.PlayerRepository, log log.Logger) refractor.PlayerService {
	return &playerService{
		repo:          repo,
		log:           log,
		recentPlayers: newRecentPlayers(config.RecentPlayersMaxSize),
	}
}

func (s *playerService) CreatePlayer(newPlayer *refractor.DBPlayer) (*refractor.Player, *refractor.ServiceResponse) {
	if err := s.repo.Create(newPlayer); err != nil {
		s.log.Error("Could not create a new player with name %s. Error: %v", newPlayer.CurrentName, err)
		return nil, refractor.InternalErrorResponse
	}

	return newPlayer.Player(), &refractor.ServiceResponse{
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

func (s *playerService) GetRecentPlayers() ([]*refractor.Player, *refractor.ServiceResponse) {
	return s.recentPlayers.getAll(), &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Recent players fetched",
	}
}

func (s *playerService) SetPlayerWatch(id int64, watch bool) *refractor.ServiceResponse {
	if _, err := s.repo.Update(id, refractor.UpdateArgs{
		"Watched": watch,
	}); err != nil {
		if err == refractor.ErrNotFound {
			return &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not updated player. Error: %v", err)
		return refractor.InternalErrorResponse
	}

	res := &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Player added to the watchlist",
	}

	if !watch {
		res.Message = "Player removed from the watchlist"
	}

	return res
}

func (s *playerService) OnPlayerJoin(serverID int64, playerGameID string, currentName string, gameConfig *refractor.GameConfig) (*refractor.Player, *refractor.ServiceResponse) {
	// Check if the player is recorded in storage
	foundPlayer, err := s.repo.FindOne(refractor.FindArgs{
		gameConfig.PlayerGameIDField: playerGameID,
	})
	if err != nil && err != refractor.ErrNotFound {
		// If there is an error and it isn't an instance of ErrNotFound, an actual error occurred that we should
		// log for traceability.
		s.log.Error("Could not check if player with %s of %s exists. Error: %v", gameConfig.PlayerGameIDField, playerGameID, err)
		return nil, refractor.InternalErrorResponse
	}

	// If foundPlayer == nil we know they don't exist, so we record them in storage
	if foundPlayer == nil {
		nullablePlayerGameID := sql.NullString{String: playerGameID, Valid: true}

		newDBPlayer := &refractor.DBPlayer{
			CurrentName: currentName,
			LastSeen:    time.Now().Unix(),
		}

		// Set proper field using reflection
		r := reflect.ValueOf(newDBPlayer)
		field := reflect.Indirect(r).FieldByName(gameConfig.PlayerGameIDField)
		field.Set(reflect.ValueOf(nullablePlayerGameID))

		newPlayer, _ := s.CreatePlayer(newDBPlayer)

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
		s.log.Info("Updating name for player (%d) %s to %s", foundPlayer.PlayerID, foundPlayer.CurrentName, currentName)

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

func (s *playerService) OnPlayerQuit(serverID int64, playerGameID string, gameConfig *refractor.GameConfig) (*refractor.Player, *refractor.ServiceResponse) {
	foundPlayer, err := s.repo.FindOne(refractor.FindArgs{
		gameConfig.PlayerGameIDField: playerGameID,
	})
	if err != nil && err != refractor.ErrNotFound {
		// If there is an error and it isn't an instance of ErrNotFound, an actual error occurred that we should
		// log for traceability.
		s.log.Error("Could not get player by %s from repo. Error: %v", gameConfig.PlayerGameIDField, err)
		return nil, refractor.InternalErrorResponse
	}

	if foundPlayer == nil {
		s.log.Error("foundPlayer with %s: %s was nil", gameConfig.PlayerGameIDField, playerGameID)
		return nil, refractor.InternalErrorResponse
	}

	// Update player's last seen field
	if _, err := s.repo.Update(foundPlayer.PlayerID, refractor.UpdateArgs{
		"LastSeen": time.Now().Unix(),
	}); err != nil {
		s.log.Error("Could not update LastSeen field for player with PlayFabID: %s. Error: %v", playerGameID, err)
		return nil, refractor.InternalErrorResponse
	}

	// Add player to recent players
	s.recentPlayers.push(foundPlayer)

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
