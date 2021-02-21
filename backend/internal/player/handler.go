package player

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/refractor"
)

type playerHandler struct {
	service       refractor.PlayerService
	serverService refractor.ServerService
	gameService   refractor.GameService
}

func NewPlayerHandler(service refractor.PlayerService, serverService refractor.ServerService, gameService refractor.GameService) refractor.PlayerHandler {
	return &playerHandler{
		service:       service,
		serverService: serverService,
		gameService:   gameService,
	}
}

func (h *playerHandler) OnPlayerJoin(fields broadcast.Fields, serverID int64) {
	// Get game config from server data
	serverData, _ := h.serverService.GetServerData(serverID)
	game, _ := h.gameService.GetGame(serverData.Game)
	gameConfig := game.GetConfig()

	h.service.OnPlayerJoin(serverID, fields[gameConfig.PlayerGameIDField], fields["name"])
}

func (h *playerHandler) OnPlayerQuit(fields broadcast.Fields, serverID int64) {
	// Get game config from server data
	serverData, _ := h.serverService.GetServerData(serverID)
	game, _ := h.gameService.GetGame(serverData.Game)
	gameConfig := game.GetConfig()

	h.service.OnPlayerQuit(serverID, fields[gameConfig.PlayerGameIDField])
}
