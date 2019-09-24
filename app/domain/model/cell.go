package model

type Cell struct {
	Mine    bool `json:"mine"`
	Clicked bool `json:"clicked"`
	Value   int  `json:"value"`
}

func NewCell() *Cell {
	return &Cell{
		Mine:    false,
		Clicked: false,
		Value:   0,
	}
}
