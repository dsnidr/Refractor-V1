package refractor

const (
	WARNING = "WARNING"
	MUTE    = "MUTE"
	KICK    = "KICK"
	BAN     = "BAN"
)

type Infraction struct {
	InfractionID int64  `json:"id"`
	PlayerID     int64  `json:"playerId"`
	UserID       int64  `json:"userId"`
	ServerID     int64  `json:"serverId"`
	Timestamp    int64  `json:"timestamp"`
	SystemAction bool   `json:"systemAction"`
	StaffName    string `json:"staffName"` // not a database field
}

type InfractionRepository interface {
	Create(infraction *Infraction) error
	FindByID(id int64) (*Infraction, error)
	Exists(args FindArgs) (bool, error)
	FindOne(args FindArgs) (*Infraction, error)
	FindManyByPlayerID(playerID int64) ([]*Infraction, error)
	FindAll() ([]*Infraction, error)
	Update(id int64, args UpdateArgs) (*Infraction, error)
	Delete(id int64) error
}

type InfractionService interface {
}
