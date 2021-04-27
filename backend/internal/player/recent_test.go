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

import (
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_recentPlayers(t *testing.T) {
	rp := newRecentPlayers(4)

	player1 := &refractor.Player{PlayerID: 1}
	player2 := &refractor.Player{PlayerID: 2}
	player3 := &refractor.Player{PlayerID: 3}
	player4 := &refractor.Player{PlayerID: 4}
	player5 := &refractor.Player{PlayerID: 5}

	rp.push(player1)
	rp.push(player2)
	rp.push(player3)
	rp.push(player4)
	rp.push(player5)

	expectedContents := []*refractor.Player{player5, player4, player3, player2}
	contents := rp.getAll()
	assert.Equal(t, expectedContents, contents)

	rp.push(player4)

	expectedContents = []*refractor.Player{player4, player5, player3, player2}
	contents = rp.getAll()
	assert.Equal(t, expectedContents, contents)
}
