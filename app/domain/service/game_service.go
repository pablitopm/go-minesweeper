package service

import (
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
