/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package refractor

import (
	"github.com/labstack/echo/v4"
	"regexp"
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
	EnableBroadcasts  bool
	BroadcastPatterns map[string]*regexp.Regexp
	CmdOutputPatterns map[string]*regexp.Regexp

	// Not all games will have support for live chat. If a game does, this should be set to true.
	EnableChat bool

	// If EnableBroadcasts is set to false, we will use polling for the playerlist instead of broadcasts.
	// Alternatively, if EnableBroadcasts is set to true this duration is used for the player refresh polling routine
	// to keep the player list in sync for games which support broadcasts.
	PlayerListPollingInterval time.Duration

	// PlayerGameIDField holds the name of the regex named properly containing the player's unique identifier for a game.
	// Using Mordhau as an example, it would be "PlayFabID".
	PlayerGameIDField string
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

type GameService interface {
	AddGame(game Game)
	GetAllGames() ([]Game, *ServiceResponse)
	GameExists(name string) (bool, *ServiceResponse)
	GetGame(name string) (Game, *ServiceResponse)
}

type GameHandler interface {
	GetAllGames(c echo.Context) error
}
