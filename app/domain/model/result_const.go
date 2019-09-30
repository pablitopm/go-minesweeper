package model

type Result int

const (
	Win Result = iota
	Lose
	Undefined
)

func (s Result) String() string {
	return [...]string{"WIN", "LOSE", "UNDEFINED"}[s]
}
