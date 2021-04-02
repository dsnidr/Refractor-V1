package refractor

type ChatReceiveBody struct {
	ServerID     int64  `json:"serverId"`
	PlayerGameID string `json:"playerGameID"`
	Name         string `json:"name"`
	Message      string `json:"message"`
	SentByUser   bool   `json:"sentByUser"`
}

type ChatService interface {
	OnChatReceive(msgBody *ChatReceiveBody, serverID int64, gameConfig *GameConfig)
	OnUserSendChat(msgBody *ChatSendBody)
}
