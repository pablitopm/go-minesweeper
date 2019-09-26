package usecase

import (
	"github.com/pablitopm/go-minesweeper/app/domain/model"
	"github.com/pablitopm/go-minesweeper/app/domain/repository"
	"github.com/pablitopm/go-minesweeper/app/domain/service"
)

type GameUsecase interface {
	StartGame(game model.Game) (model.Game, error)
	EndGame(gameId int) error
	ListGames() ([]*model.Game, error)
	GetGame(gameId int) (*model.Game, error)
	GameExists(gameId int) bool
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

func (g *gameUsecase) StartGame(game model.Game) (model.Game, error) {
	ng, _ := g.service.StartGame(game)
	g.repo.Save(&ng)
	return ng, nil
}

func (g *gameUsecase) EndGame(gameId int) error {
	return nil
}

func (g *gameUsecase) ListGames() ([]*model.Game, error) {
	return g.repo.FindAll()
}

func (g *gameUsecase) GetGame(gameId int) (*model.Game, error) {
	return g.repo.FindById(gameId)
}

func (g *gameUsecase) GameExists(gameId int) bool {
	return g.repo.GameExists(gameId)
}
