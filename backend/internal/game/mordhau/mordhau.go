package mordhau

import (
	"fmt"
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/refractor"
	"regexp"
	"time"
)

type mordhau struct {
	config *refractor.GameConfig
}

func NewMordhauGame() refractor.Game {
	return &mordhau{
		config: &refractor.GameConfig{
			UseRCON:           true,
			SendAlivePing:     true,
			AlivePingInterval: time.Second * 30,
			EnableBroadcasts:  true,
			BroadcastPatterns: map[string]*regexp.Regexp{
				broadcast.TYPE_JOIN: regexp.MustCompile("^Login: (?P<date>[0-9\\.-]+): (?P<name>.+) \\((?P<playfabid>[0-9a-fA-F]+)\\) logged in$"),
				broadcast.TYPE_QUIT: regexp.MustCompile("^Login: (?P<date>[0-9\\.-]+): (?P<name>.+) \\((?P<playfabid>[0-9a-fA-F]+)\\) logged out$"),
			},
		},
	}
}

func (g *mordhau) GetName() string {
	return "Mordhau"
}

func (g *mordhau) GetConfig() *refractor.GameConfig {
	return g.config
}

// GetWarnCommand returns an empty string since Mordhau does not have a warn command
func (g *mordhau) GetWarnCommand(args refractor.CommandArgs) string {
	return ""
}

// GetMuteCommand returns a constructed mute command for Mordhau.
// The following fields must be present on CommandArgs: PlayerID, Duration
func (g *mordhau) GetMuteCommand(args refractor.CommandArgs) string {
	return fmt.Sprintf("Mute %s %d", args.PlayerID, args.Duration)
}

// GetKickCommand returns a constructed kick command for Mordhau.
// The following fields must be present on CommandArgs: PlayerID, Reason
func (g *mordhau) GetKickCommand(args refractor.CommandArgs) string {
	return fmt.Sprintf("Kick %s %s", args.PlayerID, args.Reason)
}

// GetBanCommand returns a constructed ban command for Mordhau.
// The following fields must be present on CommandArgs: PlayerID, Duration, Reason
func (g *mordhau) GetBanCommand(args refractor.CommandArgs) string {
	return fmt.Sprintf("Ban %s %d %s", args.PlayerID, args.Duration, args.Reason)
}

func (g *mordhau) GetPlayerListCommand() string {
	return "PlayerList"
}
