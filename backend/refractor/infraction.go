package refractor

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
)

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
	StaffName    string `json:"-"` // not a database field
}

type DBInfraction struct {
	InfractionID int64
	PlayerID     int64
	UserID       int64
	ServerID     int64
	Type         string
	Reason       sql.NullString
	Duration     sql.NullInt32
	Timestamp    int64
	SystemAction bool
}

// Infraction builds a Infraction instance from the DBInstance it was called upon.
func (dbi *DBInfraction) Infraction() *Infraction {
	return &Infraction{
		InfractionID: dbi.InfractionID,
		PlayerID:     dbi.PlayerID,
		UserID:       dbi.UserID,
		ServerID:     dbi.ServerID,
		Reason:       dbi.Reason.String,
		Duration:     int(dbi.Duration.Int32),
		Type:         dbi.Type,
		Timestamp:    dbi.Timestamp,
		SystemAction: dbi.SystemAction,
	}
}

type InfractionRepository interface {
	Create(infraction *DBInfraction) (*Infraction, error)
	FindByID(id int64) (*Infraction, error)
	Exists(args FindArgs) (bool, error)
	FindOne(args FindArgs) (*Infraction, error)
	FindMany(args FindArgs) ([]*Infraction, error)
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
	DeleteInfraction(id int64, user params.UserMeta) *ServiceResponse
	UpdateInfraction(id int64, body params.UpdateInfractionParams) (*Infraction, *ServiceResponse)
	GetInfractions(infractionType string, playerID int64) ([]*Infraction, *ServiceResponse)
}

type InfractionHandler interface {
	CreateWarning(c echo.Context) error
	CreateMute(c echo.Context) error
	CreateKick(c echo.Context) error
	CreateBan(c echo.Context) error
	DeleteInfraction(c echo.Context) error
	UpdateInfraction(c echo.Context) error
	GetInfractions(infractionType string) echo.HandlerFunc
}
