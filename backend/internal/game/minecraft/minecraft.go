package minecraft

import (
	"fmt"
	"github.com/sniddunc/refractor/refractor"
	"regexp"
	"time"
)

type minecraft struct {
	config *refractor.GameConfig
}

func NewMinecraftGame() refractor.Game {
	return &minecraft{
		config: &refractor.GameConfig{
			UseRCON:                   true,
			SendAlivePing:             true,
			AlivePingInterval:         time.Second * 30,
			EnableBroadcasts:          false,
			PlayerListPollingInterval: time.Second * 5,
			BroadcastPatterns:         map[string]*regexp.Regexp{},
			CmdOutputPatterns: map[string]*regexp.Regexp{
				"PlayerList": regexp.MustCompile("^(?P<MCUUID>[0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12}):(?P<Name>[\\S]+)$"),
			},
			PlayerGameIDField: "MCUUID",
		},
	}
}

func (g *minecraft) GetName() string {
	return "Minecraft"
}

func (g *minecraft) GetConfig() *refractor.GameConfig {
	return g.config
}

// GetWarnCommand returns an empty string since Mordhau does not have a warn command
func (g *minecraft) GetWarnCommand(args refractor.CommandArgs) string {
	return ""
}

// GetMuteCommand returns a constructed mute command for Mordhau.
// The following fields must be present on CommandArgs: PlayerID, Duration
func (g *minecraft) GetMuteCommand(args refractor.CommandArgs) string {
	return fmt.Sprintf("Mute %s %d", args.PlayerID, args.Duration)
}

// GetKickCommand returns a constructed kick command for Mordhau.
// The following fields must be present on CommandArgs: PlayerID, Reason
func (g *minecraft) GetKickCommand(args refractor.CommandArgs) string {
	return fmt.Sprintf("Kick %s %s", args.PlayerID, args.Reason)
}

// GetBanCommand returns a constructed ban command for Mordhau.
// The following fields must be present on CommandArgs: PlayerID, Duration, Reason
func (g *minecraft) GetBanCommand(args refractor.CommandArgs) string {
	return fmt.Sprintf("Ban %s %d %s", args.PlayerID, args.Duration, args.Reason)
}

func (g *minecraft) GetPlayerListCommand() string {
	return "refractormc:playerlist" // use refractor minecraft plugin's command
}
