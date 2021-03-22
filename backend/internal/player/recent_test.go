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
