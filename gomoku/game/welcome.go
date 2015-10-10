package game

//welcome manages he startup message and lets you hoose a gametpye
import (
	tl "termloop"
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
	level    *tl.BaseLevel
}

func (w *welcome) Tick(event tl.Event) {
	if !w.done && event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			if w.selected < 3 {
				w.choices[w.selected].SetColor(tl.ColorBlack, tl.ColorWhite)
				w.selected++
			}
		case tl.KeyArrowLeft:
			if w.selected > 0 {
				w.choices[w.selected].SetColor(tl.ColorBlack, tl.ColorWhite)
				w.selected--
			}
		case tl.KeyArrowUp:
			if w.selected > 0 {
				w.choices[w.selected].SetColor(tl.ColorBlack, tl.ColorWhite)
				w.selected--
			}
		case tl.KeyArrowDown:
			if w.selected < 3 {
				w.choices[w.selected].SetColor(tl.ColorBlack, tl.ColorWhite)
				w.selected++
			}
		case tl.KeySpace:
			w.done = true
		}
	}
	w.choices[w.selected].SetColor(tl.ColorWhite, tl.ColorBlack)
}

func (w *welcome) Draw(screen *tl.Screen) {
	width, height := screen.Size()
	w.level.SetOffset(width/2-((boardSize*tileSizeX)/2), height/2-((boardSize*tileSizeY)/2))
	if !w.done {
		for _, v := range w.header {
			v.Draw(screen)
		}
		w.msg.Draw(screen)
		for _, v := range w.choices {
			v.Draw(screen)
		}
	} else {
	}
}

func newWelcome() *welcome {
	c := [4]string{"HOTSEAT", "P1 vs AI", "AI vs P2", "AI vs AI"}
	w := &welcome{
		msg:    tl.NewText(15, 15, msg, tl.ColorBlack, tl.ColorWhite),
		header: [12]*tl.Text{},
		choices: [4]*tl.Text{
			tl.NewText(20, 17, c[0], tl.ColorWhite, tl.ColorBlack),
			tl.NewText(20, 18, c[1], tl.ColorBlack, tl.ColorWhite),
			tl.NewText(20, 19, c[2], tl.ColorBlack, tl.ColorWhite),
			tl.NewText(20, 20, c[3], tl.ColorBlack, tl.ColorWhite),
		},
		done: false,
	}
	for k, _ := range w.header {
		w.header[k] = tl.NewText(0, k, header[k], tl.RgbTo256Color(92, 64, 51), tl.RgbTo256Color(60, 150, 180))
	}
	return w
}
