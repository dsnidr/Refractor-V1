package summary

import (
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
)

type summaryService struct {
	playerService     refractor.PlayerService
	infractionService refractor.InfractionService
	log               log.Logger
}

func NewSummaryService(playerService refractor.PlayerService, infractionService refractor.InfractionService,
	log log.Logger) refractor.SummaryService {
	return &summaryService{
		playerService:     playerService,
		infractionService: infractionService,
		log:               log,
	}
}

func (s *summaryService) GetPlayerSummary(playerID int64) (*refractor.PlayerSummary, *refractor.ServiceResponse) {
	player, res := s.playerService.GetPlayerByID(playerID)
	if !res.Success || player == nil {
		return nil, res
	}

	// Get all player infractions
	infractions, res := s.infractionService.GetPlayerInfractions(playerID)
	if !res.Success {
		return nil, res
	}

	// Explicitly define slice over using var to declare an empty array since when returned these as JSON
	// we don't want them to return as null. Instead, we want to return an empty array if there aren't any
	// infractions in any given category.
	warnings := []*refractor.Infraction{}
	mutes := []*refractor.Infraction{}
	kicks := []*refractor.Infraction{}
	bans := []*refractor.Infraction{}

	// Sort them into the different infraction types
	for _, infraction := range infractions {
		switch infraction.Type {
		case refractor.INFRACTION_TYPE_WARNING:
			warnings = append(warnings, infraction)
			break
		case refractor.INFRACTION_TYPE_MUTE:
			mutes = append(mutes, infraction)
			break
		case refractor.INFRACTION_TYPE_KICK:
			kicks = append(kicks, infraction)
			break
		case refractor.INFRACTION_TYPE_BAN:
			bans = append(bans, infraction)
			break
		}
	}

	// Build player summary
	playerSummary := &refractor.PlayerSummary{
		Warnings: warnings,
		Mutes:    mutes,
		Kicks:    kicks,
		Bans:     bans,
		Player:   player,
	}

	return playerSummary, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Player summary fetched",
	}
}
