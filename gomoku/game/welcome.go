package game

import (
	tl "gomoku/termloop"
)

var welcomeMsg tl.Canvas

type welcome struct {
	entity *tl.Entity
}

func Tick(event tl.Event) {

}

func Draw(screen *tl.Screen) {

}

func newWelcome() *welcome {
	w := &welcome{}
	//entity: tl.NewEntityFromCanvas()
	return w

}

func init() {
	welcomeMsg = CanvasFromString("hello")
}
