package model

type Status int

const (
	NewGame Status = iota
	OnGoingGame
	PausedGame
	FinishedGame
)

func (s Status) String() string {
	return [...]string{"NEW", "ON_GOING", "PAUSED", "FINISHED"}[s]
}
