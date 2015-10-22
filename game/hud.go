package game

import (
	tl "github.com/JoelOtter/termloop"
)

const hudMem = 60

type hud struct {
	screen  *tl.Screen
	header  *tl.Text
	header2 *tl.Text
	com     chan string
	toDisp  [hudMem]*tl.Text
	offsetX int
	offsetY int
}

func (h *hud) Draw(screen *tl.Screen) {
	w, _ := screen.Size()
	h.offsetX = w - 80
mainLoop:
	for {
		select {
		case str := <-h.com:
			i := hudMem - 1
			for i > 0 {
				h.toDisp[i] = h.toDisp[i-1]
				if h.toDisp[i] != nil {
					h.toDisp[i].SetPosition(h.offsetX, h.offsetY+i+2)
				}
				i--
			}
			h.toDisp[0] = tl.NewText(h.offsetX, h.offsetY+2, str, Fg, Bg)
		default:
			break mainLoop
		}
	}
	for _, v := range h.toDisp {
		v.Draw(screen)
	}
	h.header.SetPosition(h.offsetX, h.offsetY)
	h.header2.SetPosition(h.offsetX, h.offsetY+1)
	h.header.Draw(screen)
	h.header2.Draw(screen)
}

func (h *hud) Tick(event tl.Event) {

}

func newHud(com chan string, screen *tl.Screen) (h *hud) {
	h = &hud{
		screen: screen,
		com:    com,
	}
	h.header = tl.NewText(h.offsetX, h.offsetY, "Events:", Fg, Bg)
	h.header2 = tl.NewText(h.offsetX, h.offsetY, "--------------------------------------------------------------------------------", Fg, Bg)
	for k, _ := range h.toDisp {
		h.toDisp[k] = tl.NewText(h.offsetX, h.offsetY+k, "", Fg, Bg)
	}
	return h
}
