package mock

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/refractor"
	"regexp"
	"time"
)

type mockGame struct {
	config *refractor.GameConfig
}

func NewMockGame() refractor.Game {
	return &mockGame{
		config: &refractor.GameConfig{
			UseRCON:           true,
			SendAlivePing:     true,
			AlivePingInterval: time.Second * 30,
			EnableBroadcasts:  true,
			BroadcastPatterns: map[string]*regexp.Regexp{
				broadcast.TYPE_JOIN: regexp.MustCompile("^(?P<name>.+) joined the game$"),
				broadcast.TYPE_QUIT: regexp.MustCompile("^(?P<name>.+) quit the game$"),
			},
		},
	}
}

func (g *mockGame) GetName() string {
	return "TestGame"
}

func (g *mockGame) GetConfig() *refractor.GameConfig {
	return g.config
}

func (g *mockGame) GetWarnCommand(args refractor.CommandArgs) string {
	return "mockwarn"
}

func (g *mockGame) GetMuteCommand(args refractor.CommandArgs) string {
	return "mockmute"
}

func (g *mockGame) GetKickCommand(args refractor.CommandArgs) string {
	return "mockkick"
}

func (g *mockGame) GetBanCommand(args refractor.CommandArgs) string {
	return "mockban"
}

func (g *mockGame) GetPlayerListCommand() string {
	return "mocklist"
}
