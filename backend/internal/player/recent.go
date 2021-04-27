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

package player

import "github.com/sniddunc/refractor/refractor"

type recentPlayers struct {
	players []*refractor.Player
	maxSize int
}

func newRecentPlayers(maxSize int) *recentPlayers {
	return &recentPlayers{
		players: []*refractor.Player{},
		maxSize: maxSize,
	}
}

func (rp *recentPlayers) push(player *refractor.Player) {
	// Check if player already exists in array
	for i, p := range rp.players {
		if p.PlayerID == player.PlayerID {
			// Remove existing
			rp.players = append(rp.players[:i], rp.players[i+1:]...)
		}
	}

	if len(rp.players) == rp.maxSize {
		// If player is full, then remove the last entry
		rp.players = rp.players[:len(rp.players)-1]
	}

	// Prepend player to the front
	rp.players = append([]*refractor.Player{player}, rp.players...)
}

func (rp *recentPlayers) getAll() []*refractor.Player {
	return rp.players
}
