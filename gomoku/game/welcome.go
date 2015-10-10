package game

import (
	tl "termloop"
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
	w := &welcome{
		entity: tl.NewEntityFromCanvas(0, 0, welcomeMsg),
	}
	return w

}

func init() {
	welcomeMsg = tl.CanvasFromString("hello")
}
