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

package mock

import (
	"database/sql"
	"github.com/sniddunc/refractor/refractor"
	"strings"
)

type mockPlayerRepo struct {
	players map[int64]*refractor.DBPlayer
}

func NewMockPlayerRepository(mockPlayers map[int64]*refractor.DBPlayer) refractor.PlayerRepository {
	return &mockPlayerRepo{
		players: mockPlayers,
	}
}

func (r *mockPlayerRepo) Create(player *refractor.DBPlayer) error {
	newID := int64(len(r.players) + 1)
	r.players[newID] = player

	player.PlayerID = newID

	return nil
}

func (r *mockPlayerRepo) FindByID(id int64) (*refractor.Player, error) {
	foundPlayer := r.players[id]

	if foundPlayer == nil {
		return nil, refractor.ErrNotFound
	}

	return foundPlayer.Player(), nil
}

func (r *mockPlayerRepo) FindByPlayFabID(playFabID string) (*refractor.Player, error) {
	for _, player := range r.players {
		if player.PlayFabID.String == playFabID {
			return player.Player(), nil
		}
	}

	return nil, refractor.ErrNotFound
}

func (r *mockPlayerRepo) FindByMCUUID(MCUUID string) (*refractor.Player, error) {
	for _, player := range r.players {
		if player.MCUUID.String == MCUUID {
			return player.Player(), nil
		}
	}

	return nil, refractor.ErrNotFound
}

func (r *mockPlayerRepo) FindOne(args refractor.FindArgs) (*refractor.Player, error) {
	for _, player := range r.players {
		if args["PlayerID"] != nil && args["PlayerID"].(int64) != player.PlayerID {
			continue
		}

		if args["PlayFabID"] != nil && args["PlayFabID"].(string) != player.PlayFabID.String {
			continue
		}

		if args["MCUUID"] != nil && args["MCUUID"].(string) != player.PlayFabID.String {
			continue
		}

		if args["LastSeen"] != nil && args["LastSeen"].(int64) != player.LastSeen {
			continue
		}

		return player.Player(), nil
	}

	return nil, refractor.ErrNotFound
}

func (r *mockPlayerRepo) Exists(args refractor.FindArgs) (bool, error) {
	for _, player := range r.players {
		if args["PlayerID"] != nil && args["PlayerID"].(int64) != player.PlayerID {
			continue
		}

		if args["PlayFabID"] != nil && args["PlayFabID"].(string) != player.PlayFabID.String {
			continue
		}

		if args["MCUUID"] != nil && args["MCUUID"].(string) != player.PlayFabID.String {
			continue
		}

		if args["LastSeen"] != nil && args["LastSeen"].(int64) != player.LastSeen {
			continue
		}

		return true, nil
	}

	return false, nil
}

func (r *mockPlayerRepo) UpdateName(player *refractor.Player, currentName string) error {
	panic("not implemented")
}

func (r *mockPlayerRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.Player, error) {
	if r.players[id] == nil {
		return nil, refractor.ErrNotFound
	}

	if args["PlayerID"] != nil {
		r.players[id].PlayerID = args["PlayerID"].(int64)
	}

	if args["PlayFabID"] != nil {
		r.players[id].PlayFabID = sql.NullString{String: args["PlayFabID"].(string), Valid: true}
	}

	if args["MCUUID"] != nil {
		r.players[id].PlayFabID = sql.NullString{String: args["MCUUID"].(string), Valid: true}
	}

	if args["LastSeen"] != nil {
		r.players[id].LastSeen = args["LastSeen"].(int64)
	}

	return r.players[id].Player(), nil
}

func (r *mockPlayerRepo) SearchByName(name string, offset int, limit int) (int, []*refractor.Player, error) {
	var foundPlayers []*refractor.Player

	for _, player := range r.players {
		if strings.Contains(player.CurrentName, name) {
			foundPlayers = append(foundPlayers, player.Player())
		}
	}

	return len(foundPlayers), foundPlayers, nil
}

func (r *mockPlayerRepo) GetPlayerNames(id int64) (string, []string, error) {
	panic("implement me")
}
