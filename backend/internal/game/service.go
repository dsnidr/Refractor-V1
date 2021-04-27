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
	"fmt"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
)

type gameService struct {
	games map[string]refractor.Game
}

func NewGameService() refractor.GameService {
	return &gameService{
		games: map[string]refractor.Game{},
	}
}

func (s *gameService) AddGame(newGame refractor.Game) {
	s.games[newGame.GetName()] = newGame
}

func (s *gameService) GetAllGames() ([]refractor.Game, *refractor.ServiceResponse) {
	var games []refractor.Game

	for _, game := range s.games {
		games = append(games, game)
	}

	return games, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("%d games retrieved", len(games)),
	}
}

func (s *gameService) GameExists(name string) (bool, *refractor.ServiceResponse) {
	return s.games[name] != nil, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
	}
}

func (s *gameService) GetGame(name string) (refractor.Game, *refractor.ServiceResponse) {
	return s.games[name], &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
	}
}
