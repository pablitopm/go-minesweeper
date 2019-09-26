package memory

import (
	"fmt"
	"sync"

	"github.com/pablitopm/go-minesweeper/app/domain/model"
)

type gameRepository struct {
	mu    *sync.Mutex
	games map[int]*model.Game
}

func NewGameRepository() *gameRepository {
	return &gameRepository{
		mu:    &sync.Mutex{},
		games: map[int]*model.Game{},
	}
}

func (r *gameRepository) FindAll() ([]*model.Game, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	games := make([]*model.Game, len(r.games))
	i := 0
	for _, game := range r.games {
		games[i] = game
		i++
	}
	return games, nil
}

//TODO refactor this func
func (r *gameRepository) FindById(id int) (*model.Game, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, game := range r.games {
		if game.ID == id {
			return game, nil
		}
	}
	return nil, nil
}

func (r *gameRepository) Save(game *model.Game) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	game.ID = len(r.games) + 1
	fmt.Println(game.ID)
	r.games[game.ID] = game
	return nil
}

func (r *gameRepository) Count() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	return len(r.games)
}

func (r *gameRepository) GameExists(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, game := range r.games {
		if game.ID == id {
			return true
		}
	}
	return false
}
