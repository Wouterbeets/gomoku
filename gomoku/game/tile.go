package game

import (
	//"fmt"
	tl "termloop"
)

type tile struct {
	entity *tl.Entity
	player int8
}

func newTile(y, x, offY, offX int) (t *tile) {
	t = new(tile)
	t.player = UND
	t.entity = tl.NewEntity(x*tileSizeX+offX, y*tileSizeY+offY, tileSizeX, tileSizeY)
	tileCanvas := tl.CanvasFromString(tileStr)
	t.entity.ApplyCanvas(&tileCanvas)
	return t
}

func (t *tile) Draw(screen *tl.Screen) {
	t.entity.Draw(screen)
}

func (t *tile) Tick(event tl.Event) {
	return
}

func (t *tile) deselect() {
	c := &tl.Cell{
		Bg: tl.RgbTo256Color(60, 150, 180),
	}
	switch t.player {
	case UND:
		c.Ch = blank
	case P1:
		c.Ch = circle
		c.Fg = tl.ColorBlack
	case P2:
		c.Ch = circle
		c.Fg = tl.ColorWhite
	}
	t.entity.SetCell(2, 1, c)
}

func (t *tile) selected() {
	c := &tl.Cell{
		Bg: tl.ColorYellow,
	}
	switch t.player {
	case UND:
		c.Ch = blank
	case P1:
		c.Ch = circle
		c.Fg = tl.ColorBlack
	case P2:
		c.Ch = circle
		c.Fg = tl.ColorWhite
	}
	t.entity.SetCell(2, 1, c)
}

func (t *tile) white() {
	c := &tl.Cell{
		Bg: tl.RgbTo256Color(60, 150, 180),
		Fg: tl.ColorWhite,
		Ch: circle,
	}
	t.entity.SetCell(2, 1, c)
	t.player = P2
}

func (t *tile) black() {
	c := &tl.Cell{
		Bg: tl.RgbTo256Color(60, 150, 180),
		Fg: tl.ColorBlack,
		Ch: circle,
	}
	t.entity.SetCell(2, 1, c)
	t.player = P1
}
