package main

import (
	"gomoku/ai"
	"gomoku/game"
)

func main() {
	comHud := make(chan string, 200)
	ai := ai.New(2, comHud)
	go ai.Start()
	game.Start(ai.ComIn, ai.ComOut, comHud)
}
