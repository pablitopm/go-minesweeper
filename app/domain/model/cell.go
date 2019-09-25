package model

type Cell struct {
	Mine    bool `json:"mine"`
	Clicked bool `json:"clicked"`
	Value   int  `json:"value"`
}
