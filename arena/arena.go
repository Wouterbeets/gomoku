package arena

import (
	"gomoku/ai"
	//"gen"
	//"conf"
)

type Arena struct {
	ais  [2]*ai.Ai
	com1 chan string
	com2 chan string
}

//TODO: get depth from conf
func newArena(brain1) *Arena {
	com1 := make(chan string)
	com2 := make(chan string)
	a := Arena{
		ais: [2]*ai.Ai{
			ai.New(2, com1),
			ai.New(2, com2),
		},
		com1: com1,
		com2: com2,
	}
	return &a
}

func (a *Arena) Fight() {

}
