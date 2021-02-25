package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
)

type serverHandler struct {
	service       refractor.ServerService
	playerService refractor.PlayerService
	log           log.Logger
}

func NewServerHandler(service refractor.ServerService, playerService refractor.PlayerService, log log.Logger) refractor.ServerHandler {
	return &serverHandler{
		service:       service,
		playerService: playerService,
		log:           log,
	}
}

func (h *serverHandler) CreateServer(c echo.Context) error {
	body := params.CreateServerParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	server, res := h.service.CreateServer(body)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: server,
		Errors:  res.ValidationErrors,
	})
}

func (h *serverHandler) GetAllServers(c echo.Context) error {
	allServers, res := h.service.GetAllServers()
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: allServers,
	})
}

type serverDataRes struct {
	ServerID    int64               `json:"id"`
	Name        string              `json:"name"`
	Address     string              `json:"address"`
	RCONPort    string              `json:"rconPort"`
	Online      bool                `json:"online"`
	PlayerCount int                 `json:"playerCount"`
	Players     []*refractor.Player `json:"players"`
}

func (h *serverHandler) GetAllServerData(c echo.Context) error {
	allServerData, res := h.service.GetAllServerData()

	// Parse all server data into serverDataRes structs
	var resServerData []*serverDataRes

	for _, serverData := range allServerData {
		var players []*refractor.Player

		for _, player := range serverData.OnlinePlayers {
			players = append(players, player)
		}

		// Get server info
		server, res := h.service.GetServerByID(serverData.ServerID)
		if !res.Success {
			h.log.Warn("Could not GetServerByID %d", serverData.ServerID)
			continue
		}

		resServerData = append(resServerData, &serverDataRes{
			ServerID:    serverData.ServerID,
			Name:        server.Name,
			Address:     server.Address,
			RCONPort:    server.RCONPort,
			Online:      serverData.Online,
			PlayerCount: serverData.PlayerCount,
			Players:     players,
		})
	}

	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: resServerData,
	})
}

func (h *serverHandler) OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	playerGameID := gameConfig.PlayerGameIDField

	player, res := h.playerService.GetPlayer(refractor.FindArgs{
		playerGameID: fields[playerGameID],
	})

	if !res.Success {
		h.log.Error("Could not get player by their PlayerGameID field. %v = %v", playerGameID, fields[playerGameID])
		return
	}

	h.service.OnPlayerJoin(serverID, player)
}

func (h *serverHandler) OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	playerGameID := gameConfig.PlayerGameIDField

	player, res := h.playerService.GetPlayer(refractor.FindArgs{
		playerGameID: fields[playerGameID],
	})

	if !res.Success {
		h.log.Error("Could not get player by their PlayerGameID field. %v = %v", playerGameID, fields[playerGameID])
		return
	}

	h.service.OnPlayerQuit(serverID, player)
}
