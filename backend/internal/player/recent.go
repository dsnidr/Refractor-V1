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
