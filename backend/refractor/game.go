package refractor

import (
	"github.com/labstack/echo/v4"
	"time"
)

type Game interface {
	GetName() string
	GetConfig() *GameConfig
	GameCommands
}

type GameConfig struct {
	UseRCON           bool
	SendAlivePing     bool
	AlivePingInterval time.Duration
}

// CommandArgs is a struct used to supply a game's command builders with the data they need.
type CommandArgs struct {
	PlayerID string
	Reason   string
	Duration int
}

type GameCommands interface {
	GetWarnCommand(args CommandArgs) string
	GetMuteCommand(args CommandArgs) string
	GetKickCommand(args CommandArgs) string
	GetBanCommand(args CommandArgs) string
	GetPlayerListCommand() string
}

type GameInfo struct {
	Name string `json:"name"`
}

type GameService interface {
	AddGame(game Game)
	GetAllGameInfo() ([]*GameInfo, *ServiceResponse)
}

type GameHandler interface {
	GetAllGames(c echo.Context) error
}
