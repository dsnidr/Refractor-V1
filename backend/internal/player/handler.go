package player

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/refractor"
)

type playerHandler struct {
	service refractor.PlayerService
}

func NewPlayerHandler(service refractor.PlayerService) refractor.PlayerHandler {
	return &playerHandler{
		service: service,
	}
}

func (h *playerHandler) OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	h.service.OnPlayerJoin(serverID, fields[gameConfig.PlayerGameIDField], fields["Name"], gameConfig)
}

func (h *playerHandler) OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	h.service.OnPlayerQuit(serverID, fields[gameConfig.PlayerGameIDField], gameConfig)
}
