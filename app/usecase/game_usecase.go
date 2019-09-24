package usecase

import (
	"github.com/pablitopm/go-minesweeper/app/domain/model"
	"github.com/pablitopm/go-minesweeper/app/domain/repository"
	"github.com/pablitopm/go-minesweeper/app/domain/service"
)

type GameUsecase interface {
	StartGame(game model.Game) error
	EndGame(gameId string) error
}

type gameUsecase struct {
	repo    repository.GameRepository
	service *service.GameService
}

func NewGameUsecase(repo repository.GameRepository, service *service.GameService) *gameUsecase {
	return &gameUsecase{
		repo:    repo,
		service: service,
	}
}

func (g *gameUsecase) StartGame(game model.Game) error {
	return nil
}
