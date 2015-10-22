package game

//welcome manages he startup message and lets you hoose a gametpye
import (
	tl "github.com/JoelOtter/termloop"
)

const (
	msg = "please choose your gametype"
)

var header = [...]string{
	"      _____            _____         ______  _______           _____     ____    ____   ____   ____",
	"  ___|\\    \\      ____|\\    \\       |      \\/       \\     ____|\\    \\   |    |  |    | |    | |    |",
	"/    /\\    \\    /     /\\    \\     /          /\\     \\   /     /\\    \\  |    |  |    | |    | |    |",
	"|    |  |____|  /     /  \\    \\   /     /\\   / /\\     | /     /  \\    \\ |    | /    // |    | |    |",
	"|    |    ____ |     |    |    | /     /\\ \\_/ / /    /||     |    |    ||    |/ _ _//  |    | |    |",
	"|    |   |    ||     |    |    ||     |  \\|_|/ /    / ||     |    |    ||    |\\    \\'  |    | |    |",
	"|    |   |_,  ||\\     \\  /    /||     |       |    |  ||\\     \\  /    /||    | \\    \\  |    | |    |",
	"|\\ ___\\___/  /|| \\_____\\/____/ ||\\____\\       |____|  /| \\_____\\/____/ ||____|  \\____\\ |\\___\\_|____|",
	"| |   /____ / | \\ |    ||    | /| |    |      |    | /  \\ |    ||    | /|    |   |    || |    |    |",
	" \\|___|    | /   \\|____||____|/  \\|____|      |____|/    \\|____||____|/ |____|   |____| \\|____|____|",
	"   \\( |____|/       \\(    )/        \\(          )/          \\(    )/      \\(       )/      \\(   )/  ",
	"    '   )/           '    '          '          '            '    '        '       '        '   '   ",
}

type welcome struct {
	header   [12]*tl.Text
	msg      *tl.Text
	choices  [4]*tl.Text
	selected int8
	done     bool
	game     *tl.Game
}

func (w *welcome) Tick(event tl.Event) {
	if !w.done && event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			if w.selected < 3 {
				w.choices[w.selected].SetColor(Fg, Bg)
				w.selected++
			}
		case tl.KeyArrowLeft:
			if w.selected > 0 {
				w.choices[w.selected].SetColor(Fg, Bg)
				w.selected--
			}
		case tl.KeyArrowUp:
			if w.selected > 0 {
				w.choices[w.selected].SetColor(Fg, Bg)
				w.selected--
			}
		case tl.KeyArrowDown:
			if w.selected < 3 {
				w.choices[w.selected].SetColor(Fg, Bg)
				w.selected++
			}
		case tl.KeySpace:
			w.done = true
		}
	}
}

func (w *welcome) Draw(screen *tl.Screen) {
	if !w.done {
		width, height := screen.Size()
		w.game.Screen().Level().(*tl.BaseLevel).SetOffset(width/2-(len(header[0])/2), height/2-(len(header)/2))
		for _, v := range w.header {
			v.Draw(screen)
		}
		w.msg.Draw(screen)
		w.choices[w.selected].SetColor(Bg, Fg)
		for _, v := range w.choices {
			v.Draw(screen)
		}
	} else {
	}
}

func newWelcome(game *tl.Game) *welcome {
	c := [4]string{"HOTSEAT", "P1 vs AI", "AI vs P2", "AI vs AI"}
	hWidth := len(header[0])
	hHeight := len(header)
	w := &welcome{
		msg:    tl.NewText(hWidth/2-(len(msg)/2), hHeight+1, msg, Fg, Bg),
		header: [12]*tl.Text{},
		choices: [4]*tl.Text{
			tl.NewText(5, hHeight+3, c[0], Fg, Bg),
			tl.NewText(hWidth/len(c)+5, hHeight+3, c[1], Fg, Bg),
			tl.NewText((hWidth/len(c))*2+5, hHeight+3, c[2], Fg, Bg),
			tl.NewText((hWidth/len(c))*3+5, hHeight+3, c[3], Fg, Bg),
		},
		done: false,
		game: game,
	}
	for k, _ := range w.header {
		w.header[k] = tl.NewText(0, k, header[k], Fg, Bg)
	}
	return w
}
