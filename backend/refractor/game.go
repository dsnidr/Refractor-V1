package refractor

type Game interface {
	GetName() string
	GameCommands
}

// CommandArgs is a struct used to supply a game's command builders with the data they need
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
}
