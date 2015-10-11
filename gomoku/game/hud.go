package game

import (
	tl "termloop"
)

type hud struct {
	screen  *tl.Screen
	com     chan string
	toDisp  [20]*tl.Text
	offsetX int
	offsetY int
}

func (h *hud) Draw(screen *tl.Screen) {
	for _, v := range h.toDisp {
		v.Draw(screen)
	}
}

func (h *hud) Tick(event tl.Event) {
mainLoop:
	for {
		select {
		case str := <-h.com:
			i := 19
			for i > 0 {
				h.toDisp[i] = h.toDisp[i-1]
				if h.toDisp[i] != nil {
					h.toDisp[i].SetPosition(h.offsetX, h.offsetY+i)
				}
				i--
			}
			h.toDisp[0] = tl.NewText(h.offsetX, h.offsetY, str, Fg, Bg)
		default:
			break mainLoop
		}
	}
}

func newHud(com chan string, screen *tl.Screen) (h *hud) {
	h = &hud{
		screen: screen,
		com:    com,
	}
	for k, _ := range h.toDisp {
		h.toDisp[k] = tl.NewText(h.offsetX, h.offsetY+k, "", Fg, Bg)
	}
	return h
}
