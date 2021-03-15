package infraction

import (
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"net/url"
	"time"
)

type infractionService struct {
	repo          refractor.InfractionRepository
	playerService refractor.PlayerService
	serverService refractor.ServerService
	log           log.Logger
}

func NewInfractionService(repo refractor.InfractionRepository, playerService refractor.PlayerService,
	serverService refractor.ServerService, log log.Logger) refractor.InfractionService {
	return &infractionService{
		repo:          repo,
		playerService: playerService,
		serverService: serverService,
		log:           log,
	}
}

func (s *infractionService) CreateWarning(userID int64, body params.CreateWarningParams) (*refractor.Infraction, *refractor.ServiceResponse) {
	// No need to check for user perms here since that's handled at the handler level

	warning, res := s.createInfraction(body.PlayerID, userID, body.ServerID, refractor.INFRACTION_TYPE_WARNING, body.Reason,
		-1, time.Now().Unix(), false)

	return warning, res
}

func (s *infractionService) CreateMute(userID int64, body params.CreateMuteParams) (*refractor.Infraction, *refractor.ServiceResponse) {
	// No need to check for user perms here since that's handled at the handler level

	mute, res := s.createInfraction(body.PlayerID, userID, body.ServerID, refractor.INFRACTION_TYPE_MUTE, body.Reason,
		body.Duration, time.Now().Unix(), false)

	return mute, res
}

func (s *infractionService) CreateKick(userID int64, body params.CreateKickParams) (*refractor.Infraction, *refractor.ServiceResponse) {
	// No need to check for user perms here since that's handled at the handler level

	kick, res := s.createInfraction(body.PlayerID, userID, body.ServerID, refractor.INFRACTION_TYPE_KICK, body.Reason,
		-1, time.Now().Unix(), false)

	return kick, res
}

func (s *infractionService) CreateBan(userID int64, body params.CreateBanParams) (*refractor.Infraction, *refractor.ServiceResponse) {
	// No need to check for user perms here since that's handled at the handler level

	ban, res := s.createInfraction(body.PlayerID, userID, body.ServerID, refractor.INFRACTION_TYPE_BAN, body.Reason,
		body.Duration, time.Now().Unix(), false)

	return ban, res
}

// We don't just make this function a member of the infraction service interface because there is a good chance we'll need to wrap
// other code around this logic in the future. To avoid code repetition, the creation logic was moved into this function.
func (s *infractionService) createInfraction(playerID int64, userID int64, serverID int64, infractionType string,
	reason string, duration int, timestamp int64, systemAction bool) (*refractor.Infraction, *refractor.ServiceResponse) {

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

	newInfraction := &refractor.Infraction{
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

	return infraction, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Infraction created",
	}
}
