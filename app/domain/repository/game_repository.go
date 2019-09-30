package repository

import "github.com/pablitopm/go-minesweeper/app/domain/model"

type GameRepository interface {
	FindAll() ([]*model.Game, error)
	FindById(id int) (*model.Game, error)
	Count() int
	Upsert(*model.Game) error
	GameExists(id int) bool
}
