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

package gameserver

import (
	"fmt"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
)

type gameServerService struct {
	gameService   refractor.GameService
	serverService refractor.ServerService
	log           log.Logger
}

func NewGameServerService(gameService refractor.GameService, serverService refractor.ServerService, log log.Logger) refractor.GameServerService {
	return &gameServerService{
		gameService:   gameService,
		serverService: serverService,
		log:           log,
	}
}

func (s *gameServerService) GetGameServers() ([]*refractor.GameServer, *refractor.ServiceResponse) {
	games, res := s.gameService.GetAllGames()
	if !res.Success {
		s.log.Error("GetAllGameInfo() response success was false")
		return nil, &refractor.ServiceResponse{
			Success:    res.Success,
			StatusCode: res.StatusCode,
			Message:    res.Message,
		}
	}

	servers, res := s.serverService.GetAllServers()
	if !res.Success {
		s.log.Error("GetAllServers() response success was false")
		return nil, &refractor.ServiceResponse{
			Success:    res.Success,
			StatusCode: res.StatusCode,
			Message:    res.Message,
		}
	}

	gameServers := map[string]*refractor.GameServer{}
	for _, game := range games {
		config := game.GetConfig()

		gameServers[game.GetName()] = &refractor.GameServer{
			Name: game.GetName(),
			Config: &refractor.GameServerConfig{
				EnableChat: config.EnableChat,
			},
			Servers: []*refractor.ServerInfo{},
		}
	}

	for _, server := range servers {
		gameServers[server.Game].Servers = append(gameServers[server.Game].Servers, &refractor.ServerInfo{
			ServerID: server.ServerID,
			Name:     server.Name,
			Game:     server.Game,
			Address:  server.Address,
		})
	}

	var gameServersOutput []*refractor.GameServer
	for _, gameServer := range gameServers {
		gameServersOutput = append(gameServersOutput, gameServer)
	}

	return gameServersOutput, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Fetched %d game server payloads", len(gameServersOutput)),
	}
}
