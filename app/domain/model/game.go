package model

import "time"

type Game struct {
	Id        int       `json:"id"`
	Rows      int       `json:"rows"`
	Cols      int       `json:"cols"`
	Mines     int       `json:"mines"`
	Status    string    `json:"status"`
	StartTime time.Time `json:"startTime"`
	Grid      [][]Cell  `json:"grid,omitempty"`
	//User      User      `json:"user"`
}

func NewGame(cols int, rows int, mines int) *Game {
	return &Game{
		Rows:      rows,
		Cols:      cols,
		Mines:     mines,
		StartTime: time.Now(),
		Grid:      [][]Cell{},
	}
}
