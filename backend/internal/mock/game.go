package mock

import (
	"github.com/sniddunc/refractor/refractor"
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
