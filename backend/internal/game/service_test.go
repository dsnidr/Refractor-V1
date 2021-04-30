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

package game

import (
	"github.com/sniddunc/refractor/internal/mock"
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_gameService_AddGame(t *testing.T) {
	type fields struct {
		games map[string]refractor.Game
	}
	type args struct {
		newGame refractor.Game
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "game.addgame.1",
			fields: fields{
				games: map[string]refractor.Game{},
			},
			args: args{
				newGame: mock.NewMockGame(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameService := &gameService{
				games: tt.fields.games,
			}

			gameService.AddGame(tt.args.newGame)

			assert.NotNil(t, gameService.games[tt.args.newGame.GetName()])
		})
	}
}

func Test_gameService_GameExists(t *testing.T) {
	type fields struct {
		games map[string]refractor.Game
	}
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      bool
		wantError *refractor.ServiceResponse
	}{
		{
			name: "game.gameexists.1",
			fields: fields{
				games: map[string]refractor.Game{
					mock.NewMockGame().GetName(): mock.NewMockGame(),
				},
			},
			args: args{
				name: mock.NewMockGame().GetName(),
			},
			want:      true,
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameService := &gameService{
				games: tt.fields.games,
			}

			exists, err := gameService.GameExists(tt.args.name)

			assert.Nil(t, tt.wantError, err, "Error was returned")
			assert.Equal(t, tt.want, exists)
		})
	}
}

func Test_gameService_GetAllGames(t *testing.T) {
	type fields struct {
		games map[string]refractor.Game
	}
	tests := []struct {
		name   string
		fields fields
		want   []refractor.Game
	}{
		{
			name: "game.getallgames.1",
			fields: fields{
				games: map[string]refractor.Game{
					mock.NewMockGame().GetName(): mock.NewMockGame(),
				},
			},
			want: []refractor.Game{mock.NewMockGame()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameService := &gameService{
				games: tt.fields.games,
			}

			games, _ := gameService.GetAllGames()

			assert.Equal(t, tt.want, games)
		})
	}
}

func Test_gameService_GetGame(t *testing.T) {
	type fields struct {
		games map[string]refractor.Game
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   refractor.Game
	}{
		{
			name: "game.getgame.1",
			fields: fields{
				games: map[string]refractor.Game{
					mock.NewMockGame().GetName(): mock.NewMockGame(),
				},
			},
			args: args{
				name: mock.NewMockGame().GetName(),
			},
			want: mock.NewMockGame(),
		},
		{
			name: "game.getgame.2",
			fields: fields{
				games: map[string]refractor.Game{},
			},
			args: args{
				name: "doesnotexist",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameService := &gameService{
				games: tt.fields.games,
			}

			game, _ := gameService.GetGame(tt.args.name)

			assert.Equal(t, tt.want, game)
		})
	}
}
