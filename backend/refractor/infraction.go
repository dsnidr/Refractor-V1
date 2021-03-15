package refractor

import "github.com/sniddunc/refractor/internal/params"

const (
	INFRACTION_TYPE_WARNING = "WARNING"
	INFRACTION_TYPE_MUTE    = "MUTE"
	INFRACTION_TYPE_KICK    = "KICK"
	INFRACTION_TYPE_BAN     = "BAN"
)

var InfractionTypes = []string{INFRACTION_TYPE_WARNING, INFRACTION_TYPE_MUTE, INFRACTION_TYPE_KICK, INFRACTION_TYPE_BAN}

type Infraction struct {
	InfractionID int64  `json:"id"`
	PlayerID     int64  `json:"playerId"`
	UserID       int64  `json:"userId"`
	ServerID     int64  `json:"serverId"`
	Type         string `json:"type"`
	Reason       string `json:"reason"`
	Duration     int    `json:"duration"`
	Timestamp    int64  `json:"timestamp"`
	SystemAction bool   `json:"systemAction"`
	StaffName    string `json:"staffName"` // not a database field
}

type InfractionRepository interface {
	Create(infraction *Infraction) (*Infraction, error)
	FindByID(id int64) (*Infraction, error)
	Exists(args FindArgs) (bool, error)
	FindOne(args FindArgs) (*Infraction, error)
	FindManyByPlayerID(playerID int64) ([]*Infraction, error)
	FindAll() ([]*Infraction, error)
	Update(id int64, args UpdateArgs) (*Infraction, error)
	Delete(id int64) error
}

type InfractionService interface {
	CreateWarning(userID int64, body params.CreateWarningParams) (*Infraction, *ServiceResponse)
	CreateMute(userID int64, body params.CreateMuteParams) (*Infraction, *ServiceResponse)
	CreateKick(userID int64, body params.CreateKickParams) (*Infraction, *ServiceResponse)
	CreateBan(userID int64, body params.CreateBanParams) (*Infraction, *ServiceResponse)
}
