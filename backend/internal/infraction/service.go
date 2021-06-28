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

package infraction

import (
	"database/sql"
	"fmt"
	"github.com/sniddunc/bitperms"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/perms"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"net/url"
	"time"
)

type infractionService struct {
	repo                        refractor.InfractionRepository
	playerService               refractor.PlayerService
	serverService               refractor.ServerService
	userService                 refractor.UserService
	infractionCreateSubscribers []refractor.InfractionCreateSubscriber
	log                         log.Logger
}

func NewInfractionService(repo refractor.InfractionRepository, playerService refractor.PlayerService,
	serverService refractor.ServerService, userService refractor.UserService, log log.Logger) refractor.InfractionService {
	return &infractionService{
		repo:                        repo,
		playerService:               playerService,
		serverService:               serverService,
		userService:                 userService,
		infractionCreateSubscribers: []refractor.InfractionCreateSubscriber{},
		log:                         log,
	}
}

func (s *infractionService) CreateWarning(userID int64, body params.CreateWarningParams) (*refractor.Infraction, *refractor.ServiceResponse) {
	// Create nullable fields from values
	duration := sql.NullInt32{}
	reason := sql.NullString{String: body.Reason, Valid: true}

	warning, res := s.createInfraction(body.PlayerID, userID, body.ServerID, refractor.INFRACTION_TYPE_WARNING, reason,
		duration, time.Now().Unix(), false)

	return warning, res
}

func (s *infractionService) CreateMute(userID int64, body params.CreateMuteParams) (*refractor.Infraction, *refractor.ServiceResponse) {
	// Create nullable fields from values
	duration := sql.NullInt32{Int32: int32(body.Duration), Valid: true}
	reason := sql.NullString{String: body.Reason, Valid: true}

	mute, res := s.createInfraction(body.PlayerID, userID, body.ServerID, refractor.INFRACTION_TYPE_MUTE, reason,
		duration, time.Now().Unix(), false)

	return mute, res
}

func (s *infractionService) CreateKick(userID int64, body params.CreateKickParams) (*refractor.Infraction, *refractor.ServiceResponse) {
	// Create nullable fields from values
	duration := sql.NullInt32{}
	reason := sql.NullString{String: body.Reason, Valid: true}

	kick, res := s.createInfraction(body.PlayerID, userID, body.ServerID, refractor.INFRACTION_TYPE_KICK, reason,
		duration, time.Now().Unix(), false)

	return kick, res
}

func (s *infractionService) CreateBan(userID int64, body params.CreateBanParams) (*refractor.Infraction, *refractor.ServiceResponse) {
	// Create nullable fields from values
	duration := sql.NullInt32{Int32: int32(body.Duration), Valid: true}
	reason := sql.NullString{String: body.Reason, Valid: true}

	ban, res := s.createInfraction(body.PlayerID, userID, body.ServerID, refractor.INFRACTION_TYPE_BAN, reason,
		duration, time.Now().Unix(), false)

	return ban, res
}

// We don't just make this function a member of the infraction service interface because there is a good chance we'll need to wrap
// other code around this logic in the future. To avoid code repetition, the creation logic was moved into this function.
func (s *infractionService) createInfraction(playerID int64, userID int64, serverID int64, infractionType string,
	reason sql.NullString, duration sql.NullInt32, timestamp int64, systemAction bool) (*refractor.Infraction, *refractor.ServiceResponse) {

	// Make sure player exists
	player, _ := s.playerService.GetPlayerByID(playerID)
	if player == nil {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			ValidationErrors: url.Values{
				"playerId": []string{"Invalid player ID"},
			},
		}
	}

	// Make sure server exists
	server, _ := s.serverService.GetServerByID(serverID)
	if server == nil {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			ValidationErrors: url.Values{
				"serverId": []string{"Invalid server ID"},
			},
		}
	}

	newInfraction := &refractor.DBInfraction{
		PlayerID:     playerID,
		UserID:       userID,
		ServerID:     serverID,
		Type:         infractionType,
		Reason:       reason,
		Duration:     duration,
		Timestamp:    timestamp,
		SystemAction: systemAction,
	}

	infraction, err := s.repo.Create(newInfraction)
	if err != nil {
		s.log.Error("Could not create new infraction in repo. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	// Notify subscribers
	if len(s.infractionCreateSubscribers) > 0 {
		infraction.PlayerName = player.CurrentName

		user, _ := s.userService.GetUserByID(userID)
		if user != nil {
			infraction.StaffName = user.Username
		}

		for _, subscriber := range s.infractionCreateSubscribers {
			subscriber(infraction)
		}
	}

	return infraction, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Infraction created",
	}
}

func (s *infractionService) DeleteInfraction(id int64, user params.UserMeta) *refractor.ServiceResponse {
	userPerms := bitperms.PermissionValue(user.Permissions)

	infraction, err := s.repo.FindByID(id)
	if err != nil {
		if err == refractor.ErrNotFound {
			return &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not find infraction by id %d. Error: %v", id, err)
		return refractor.InternalErrorResponse
	}

	hasPermission := false

	// Check if the user has full access or can delete any infraction. If they do, skip the permissions check
	if perms.UserHasFullAccess(userPerms) || userPerms.HasFlag(perms.DELETE_ANY_INFRACTION) {
		hasPermission = true
	}

	// Check if the user created this infraction. If they did, check if they have permission to delete their own infractions.
	if user.UserID == infraction.UserID && userPerms.HasFlag(perms.DELETE_OWN_INFRACTIONS) {
		hasPermission = true
	}

	if !hasPermission {
		return &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageNoPermission,
		}
	}

	// If the above statement didn't return, the user has permission. Delete infraction.
	if err := s.repo.Delete(id); err != nil {
		s.log.Error("Could not delete infraction ID %d. Error: %v", id, err)
		return refractor.InternalErrorResponse
	}

	return &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Infraction deleted",
	}
}

func (s *infractionService) UpdateInfraction(id int64, body params.UpdateInfractionParams) (*refractor.Infraction, *refractor.ServiceResponse) {
	// Make sure infraction exists
	foundInfraction, err := s.repo.FindByID(id)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not get infraction by id %d. Error: %v", id, err)
		return nil, refractor.InternalErrorResponse
	}

	userPerms := bitperms.PermissionValue(body.UserMeta.Permissions)

	// We need to make sure that the user has permission to update this infraction.
	// We do this by:
	//   a) checking if they're a super admin, have full access or can edit any infraction.
	//   b) checking if they created this infraction and if they have permission to edit their own infractions.

	hasPermission := false

	// Check if the user is a super admin, has full access or can edit any infraction
	if perms.UserIsSuperAdmin(userPerms) || perms.UserHasFullAccess(userPerms) || userPerms.HasFlag(perms.EDIT_ANY_INFRACTION) {
		hasPermission = true
	}

	// Check if the user created this infraction and has permission to edit their own infractions
	if !hasPermission && foundInfraction.UserID == body.UserMeta.UserID && userPerms.HasFlag(perms.EDIT_OWN_INFRACTIONS) {
		hasPermission = true
	}

	// Final check of hasPermission
	if !hasPermission {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageNoPermission,
		}
	}

	// If the above statement didn't return, we know the user has permission so we proceed with updating the infraction.
	updateArgs := refractor.UpdateArgs{}
	if body.Reason != nil {
		updateArgs["Reason"] = *body.Reason
	}

	// This is a bit messy and definitely not the best way to handle this.
	// Ideally we'd have an infraction type with a list of allowed fields on it, but for now we just check if
	// it's a mute or ban and only allow the duration field to be set if it is.
	if foundInfraction.Type == refractor.INFRACTION_TYPE_MUTE ||
		foundInfraction.Type == refractor.INFRACTION_TYPE_BAN {
		if body.Duration != nil {
			updateArgs["Duration"] = *body.Duration
		}
	}

	if len(updateArgs) == 0 {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "No update fields provided",
		}
	}

	updatedInfraction, err := s.repo.Update(foundInfraction.InfractionID, updateArgs)
	if err != nil {
		s.log.Error("Could not update infraction with id %d. Error: %v", id, err)
		return nil, refractor.InternalErrorResponse
	}

	return updatedInfraction, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Infraction updated",
	}
}

func (s *infractionService) GetPlayerInfractionsType(infractionType string, playerID int64) ([]*refractor.Infraction, *refractor.ServiceResponse) {
	infractions, err := s.repo.FindMany(refractor.FindArgs{
		"PlayerID": playerID,
		"Type":     infractionType,
	})
	if err != nil {
		if err == refractor.ErrNotFound {
			return []*refractor.Infraction{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Fetched 0 infractions",
			}
		}

		s.log.Error("Could not find many infractions. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	// Get staff names
	for _, infraction := range infractions {
		user, err := s.userService.GetUserByID(infraction.UserID)
		if err != nil {
			s.log.Error("Could not get infraction user by ID. Error: %v", err)
			continue
		}

		// Set StaffName to contain the staff member's username
		infraction.StaffName = user.Username
	}

	return infractions, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Fetched %d infractions", len(infractions)),
	}
}

func (s *infractionService) GetPlayerInfractions(playerID int64) ([]*refractor.Infraction, *refractor.ServiceResponse) {
	infractions, err := s.repo.FindManyByPlayerID(playerID)
	if err != nil {
		if err == refractor.ErrNotFound {
			return []*refractor.Infraction{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Fetched 0 infractions",
			}
		}

		s.log.Error("Could not find many infractions by player ID. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	// Get staff names
	for _, infraction := range infractions {
		user, err := s.userService.GetUserByID(infraction.UserID)
		if err != nil {
			s.log.Error("Could not get infraction user by ID. Error: %v", err)
			continue
		}

		// Set StaffName to contain the staff member's username
		infraction.StaffName = user.Username
	}

	return infractions, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Fetched %d infractions", len(infractions)),
	}
}

func (s *infractionService) GetRecentInfractions(count int) ([]*refractor.Infraction, *refractor.ServiceResponse) {
	infractions, err := s.repo.GetRecent(count)
	if err != nil {
		if err == refractor.ErrNotFound {
			return []*refractor.Infraction{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Fetched 0 recent infractions",
			}
		}

		s.log.Error("Could not get recent infractions. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	return infractions, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Fetched %d recent infractions", len(infractions)),
	}
}

func (s *infractionService) SubscribeInfractionCreate(subscriber refractor.InfractionCreateSubscriber) {
	s.infractionCreateSubscribers = append(s.infractionCreateSubscribers, subscriber)
}
