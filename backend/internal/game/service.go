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

func (s *gameService) GetAllGameInfo() ([]*refractor.GameInfo, *refractor.ServiceResponse) {
	var gameInfoList []*refractor.GameInfo

	for _, listGame := range s.games {
		gameInfo := &refractor.GameInfo{
			Name: listGame.GetName(),
		}

		gameInfoList = append(gameInfoList, gameInfo)
	}

	return gameInfoList, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("%d games retrieved", len(gameInfoList)),
	}
}

func (s *gameService) GameExists(name string) (bool, *refractor.ServiceResponse) {
	return s.games[name] != nil, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
	}
}
