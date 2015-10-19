package main

import (
	"fmt"
	"gomoku/ai"
	"gomoku/game"
)

func main() {
	fmt.Println("main")
	ai := ai.New()
	go ai.Start()
	game.Start(ai.ComIn, ai.ComOut)
}
