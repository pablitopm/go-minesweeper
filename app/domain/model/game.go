package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Status int

const (
	New Status = iota
	OnGoing
	Paused
	Finished
)

func (s Status) String() string {
	return [...]string{"NEW", "ON_GOING", "PAUSED", "FINISHED"}[s]
}

type Game struct {
	ID        int       `json:"id"`
	Rows      int       `json:"rows"`
	Cols      int       `json:"cols"`
	Mines     int       `json:"mines"`
	Status    Status    `json:"status"`
	StartTime time.Time `json:"startTime"`
	Grid      [][]Cell  `json:"grid,omitempty"`
	//User      User      `json:"user"`
}

func (g Game) Validate() error {
	return validation.ValidateStruct(&g,
		// Rows cannot be empty, and the length must be between 1 and 100
		validation.Field(&g.Mines, validation.Required, validation.Max(100)),
		validation.Field(&g.Mines, validation.Required, validation.Min(1)),
		// Cols cannot be empty, and the length must be between 1 and 100
		validation.Field(&g.Cols, validation.Required, validation.Max(100)),
		validation.Field(&g.Cols, validation.Required, validation.Min(1)),
		// Mines cannot be empty, and the lenght must be between 0 and rows*cols
		validation.Field(&g.Mines, validation.Required, validation.Max(g.Rows*g.Cols)),
		validation.Field(&g.Mines, validation.Required, validation.Min(0)),
	)
}
