package rcon

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/refractor"
)

func (s *rconService) HandleJoinBroadcast(bcast *broadcast.Broadcast, serverID int64, gameConfig *refractor.GameConfig) {
	for _, sub := range s.joinSubscribers {
		sub(bcast.Fields, serverID, gameConfig)
	}
}

func (s *rconService) HandleQuitBroadcast(bcast *broadcast.Broadcast, serverID int64, gameConfig *refractor.GameConfig) {
	for _, sub := range s.quitSubscribers {
		sub(bcast.Fields, serverID, gameConfig)
	}
}

func (s *rconService) HandleChatBroadcast(bcast *broadcast.Broadcast, serverID int64, gameConfig *refractor.GameConfig) {
	fields := bcast.Fields

	for _, sub := range s.chatSubscribers {
		msgBody := &refractor.ChatReceiveBody{
			ServerID:     serverID,
			PlayerGameID: fields[gameConfig.PlayerGameIDField],
			Name:         fields["Name"],
			Message:      fields["Message"],
			SentByUser:   false,
		}

		sub(msgBody, serverID, gameConfig)
	}
}
