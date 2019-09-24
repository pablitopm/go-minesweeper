package repository

import "github.com/pablitopm/go-minesweeper/app/domain/model"

type GameRepository interface {
	FindAll() ([]*model.Game, error)
	FindById(id int) (*model.Game, error)
	FindByUser(user model.Game) (*model.Game, error)
	Save(*model.Game) error
}
