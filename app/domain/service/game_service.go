package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pablitopm/go-minesweeper/app/domain/model"
	"github.com/pablitopm/go-minesweeper/app/domain/repository"
)

type GameService struct {
	repo repository.GameRepository
}

func NewGameService(repo repository.GameRepository) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (g *GameService) StartGame(game model.Game) (model.Game, error) {
	game.ID = g.repo.Count()
	game.StartTime = time.Now()
	game.Status = model.New
	createGrid(&game)
	g.repo.Save(&game)
	return game, nil
}

func createGrid(game *model.Game) error {
	//make 2D array
	game.Grid = make([][]model.Cell, game.Rows)
	cells := make([]model.Cell, game.Rows*game.Cols)
	for i := range game.Grid {
		game.Grid[i] = cells[i*game.Cols : (i+1)*game.Cols]
	}

	//setting mines
	i := 0
	rand.Seed(time.Now().UTC().UnixNano())
	for i < game.Mines {
		idx := rand.Intn(game.Rows * game.Cols)
		if !cells[idx].Mine {
			cells[idx].Mine = true
			i++
		}
	}

	// Set cell values
	for i, row := range game.Grid {
		for j, cell := range row {
			if cell.Mine {
				setAdjacentValues(game, i, j)
			}
		}
	}

	fmt.Println(game.Grid)
	return nil
}

func setAdjacentValues(game *model.Game, i, j int) {
	for z := i - 1; z < i+2; z++ {
		if z < 0 || z > game.Rows-1 {
			continue
		}
		for w := j - 1; w < j+2; w++ {
			if w < 0 || w > game.Cols-1 {
				continue
			}
			if z == i && w == j {
				continue
			}
			game.Grid[z][w].Value++
		}
	}
}
