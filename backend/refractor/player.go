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

package refractor

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/pkg/broadcast"
)

type Player struct {
	PlayerID        int64    `json:"id"`
	PlayFabID       string   `json:"playFabId"`
	MCUUID          string   `json:"mcuuid"`
	LastSeen        int64    `json:"lastSeen"`
	CurrentName     string   `json:"currentName"`
	PreviousNames   []string `json:"previousNames,omitempty"`
	Watched         bool     `json:"watched"`
	InfractionCount *int     `json:"infractionCount,omitempty"` // not a db field
}

type DBPlayer struct {
	PlayerID      int64
	PlayFabID     sql.NullString
	MCUUID        sql.NullString
	LastSeen      int64
	CurrentName   string
	PreviousNames []string
	Watched       bool `json:"watched"`
}

func (dbp DBPlayer) Player() *Player {
	return &Player{
		PlayerID:      dbp.PlayerID,
		PlayFabID:     dbp.PlayFabID.String,
		MCUUID:        dbp.MCUUID.String,
		LastSeen:      dbp.LastSeen,
		CurrentName:   dbp.CurrentName,
		PreviousNames: dbp.PreviousNames,
		Watched:       dbp.Watched,
	}
}

type PlayerUpdateSubscriber func(updated *Player)

type PlayerRepository interface {
	Create(player *DBPlayer) error
	FindByID(id int64) (*Player, error)
	FindByPlayFabID(playFabID string) (*Player, error)
	FindByMCUUID(MCUUID string) (*Player, error)
	FindOne(args FindArgs) (*Player, error)
	Exists(args FindArgs) (bool, error)
	UpdateName(player *Player, currentName string) error
	Update(id int64, args UpdateArgs) (*Player, error)
	SearchByName(name string, limit int, offset int) (int, []*Player, error)
}

type PlayerService interface {
	CreatePlayer(newPlayer *DBPlayer) (*Player, *ServiceResponse)
	GetPlayerByID(id int64) (*Player, *ServiceResponse)
	GetPlayer(args FindArgs) (*Player, *ServiceResponse)
	GetRecentPlayers() ([]*Player, *ServiceResponse)
	SetPlayerWatch(id int64, watch bool) *ServiceResponse
	OnPlayerJoin(serverID int64, playerGameID string, currentName string, gameConfig *GameConfig) (*Player, *ServiceResponse)
	OnPlayerQuit(serverID int64, playerGameID string, gameConfig *GameConfig) (*Player, *ServiceResponse)
	SubscribeUpdate(subscriber PlayerUpdateSubscriber)
}

type PlayerHandler interface {
	GetRecentPlayers(c echo.Context) error
	SwitchPlayerWatch(watch bool) echo.HandlerFunc
	OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
	OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
}
