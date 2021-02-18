package rcon

import "github.com/sniddunc/refractor/pkg/broadcast"

type playerJoinData struct {
	ServerID     int64  `json:"serverId"`
	PlayerID     int64  `json:"id"`
	PlayerGameID string `json:"playerGameId"`
	Name         string `json:"name"`
}

func (s *rconService) HandleJoinBroadcast(bcast *broadcast.Broadcast, serverID int64) {
	for _, sub := range s.joinSubscribers {
		sub(bcast.Fields, serverID)
	}
}

func (s *rconService) HandleQuitBroadcast(bcast *broadcast.Broadcast, serverID int64) {
	for _, sub := range s.quitSubscribers {
		sub(bcast.Fields, serverID)
	}
}
