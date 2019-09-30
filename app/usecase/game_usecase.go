package usecase

import (
	"fmt"

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
	ClickCell(game *model.Game, row, col int) (*model.Game, error)
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
	g.repo.Upsert(&ng)
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

func (g *gameUsecase) ClickCell(game *model.Game, row, col int) (*model.Game, error) {
	if (row >= game.Rows || row < 0) || (col >= game.Cols || col < 0) {
		return nil, fmt.Errorf("Row or col exceeded grid limits")
	}

	if game.Status == model.FinishedGame {
		return nil, fmt.Errorf("Cannot played a finished game")
	}

	game.Status = model.OnGoingGame

	if game.Grid[row][col].Clicked {
		return nil, fmt.Errorf("Cell in row:%d, col:%d already clicked", row, col)
	}

	game.Grid[row][col].Clicked = true
	game.CellsRevealed++

	if game.Grid[row][col].Mine {
		game.Status = model.FinishedGame
		game.Result = model.Lose
		return game, nil
	}

	if game.Grid[row][col].Value == 0 {
		revealEmptyCells(game, row, col)
	}
	if win(game) {
		game.Status = model.FinishedGame
		game.Result = model.Win
	}
	g.repo.Upsert(game)

	return game, nil
}

func revealEmptyCells(game *model.Game, row, col int) {
	for z := row - 1; z <= row+1; z++ {
		if z < 0 || z > game.Rows-1 {
			continue
		}
		for w := col - 1; w <= col+1; w++ {
			if w < 0 || w > game.Cols-1 {
				continue
			}
			if z == row && w == col {
				continue
			}
			if game.Grid[z][w].Clicked {
				continue
			}
			game.Grid[z][w].Clicked = true
			game.CellsRevealed++
			if game.Grid[z][w].Value == 0 {
				revealEmptyCells(game, z, w)
			}
		}
	}
}

func win(game *model.Game) bool {
	return game.CellsRevealed == ((game.Rows + game.Cols) - game.Mines)
}
