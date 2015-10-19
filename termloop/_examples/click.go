package main

import tl "github.com/JoelOtter/termloop"

type Clickable struct {
	r *tl.Rectangle
}

func NewClickable(x, y, w, h int, col tl.Attr) *Clickable {
	return &Clickable{
		r: tl.NewRectangle(x, y, w, h, col),
	}
}

func (c *Clickable) Draw(s *tl.Screen) {
	c.r.Draw(s)
}

func (c *Clickable) Tick(ev tl.Event) {
	x, y := c.r.Position()
	if ev.Type == tl.EventMouse && ev.MouseX == x && ev.MouseY == y {
		if c.r.Color() == tl.ColorWhite {
			c.r.SetColor(tl.ColorBlack)
		} else {
			c.r.SetColor(tl.ColorWhite)
		}
	}
}

func main() {
	g := tl.NewGame()

	for i := 0; i < 40; i++ {
		for j := 0; j < 20; j++ {
			g.Screen().AddEntity(NewClickable(i, j, 1, 1, tl.ColorWhite))
		}
	}

	g.Start()
}
