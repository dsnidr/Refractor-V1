package refractor

import "github.com/sniddunc/refractor/pkg/broadcast"

type Player struct {
	PlayerID      int64    `json:"id"`
	PlayFabID     string   `json:"playFabId"`
	LastSeen      int64    `json:"lastSeen"`
	CurrentName   string   `json:"currentName"`
	PreviousNames []string `json:"previousNames,omitempty"`
}

type PlayerRepository interface {
	Create(player *Player) error
	FindByID(id int64) (*Player, error)
	FindByPlayFabID(playFabID string) (*Player, error)
	FindOne(args FindArgs) (*Player, error)
	Exists(args FindArgs) (bool, error)
	UpdateName(player *Player, currentName string) error
	Update(id int64, args UpdateArgs) (*Player, error)
}

type PlayerService interface {
	CreatePlayer(newPlayer *Player) (*Player, *ServiceResponse)
	GetPlayerByID(id int64) (*Player, *ServiceResponse)
	GetPlayer(args FindArgs) (*Player, *ServiceResponse)
	OnPlayerJoin(serverID int64, playerGameID string, currentName string) (*Player, *ServiceResponse)
	OnPlayerQuit(serverID int64, playerGameID string) (*Player, *ServiceResponse)
}

type PlayerHandler interface {
	OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
	OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
}
