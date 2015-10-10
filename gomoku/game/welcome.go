package game

//welcome manages he startup message and lets you hoose a gametpye
import (
	tl "termloop"
)

const (
	header = `
      _____            _____         ______  _______           _____     ____    ____   ____   ____ 
  ___|\    \      ____|\    \       |      \/       \     ____|\    \   |    |  |    | |    | |    |
 /    /\    \    /     /\    \     /          /\     \   /     /\    \  |    |  |    | |    | |    |
|    |  |____|  /     /  \    \   /     /\   / /\     | /     /  \    \ |    | /    // |    | |    |
|    |    ____ |     |    |    | /     /\ \_/ / /    /||     |    |    ||    |/ _ _//  |    | |    |
|    |   |    ||     |    |    ||     |  \|_|/ /    / ||     |    |    ||    |\    \'  |    | |    |
|    |   |_,  ||\     \  /    /||     |       |    |  ||\     \  /    /||    | \    \  |    | |    |
|\ ___\___/  /|| \_____\/____/ ||\____\       |____|  /| \_____\/____/ ||____|  \____\ |\___\_|____|
| |   /____ / | \ |    ||    | /| |    |      |    | /  \ |    ||    | /|    |   |    || |    |    |
 \|___|    | /   \|____||____|/  \|____|      |____|/    \|____||____|/ |____|   |____| \|____|____|
   \( |____|/       \(    )/        \(          )/          \(    )/      \(       )/      \(   )/  
    '   )/           '    '          '          '            '    '        '       '        '   '   
	`
	msg = `
									please choose your gametype
			`
)

var welcomeCanvas tl.Canvas

type welcome struct {
	header   *tl.Entity
	msg      *tl.Text
	choices  [4]*tl.Text
	selected int8
	done     bool
}

func (w *welcome) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			if w.selected < 3 {
				w.selected++
			}
		case tl.KeyArrowLeft:
			if w.selected > 0 {
				w.selected--
			}
		case tl.KeyArrowUp:
			if w.selected > 0 {
				w.selected--
			}
		case tl.KeyArrowDown:
			if w.selected < 3 {
				w.selected++
			}
		case tl.KeySpace:
		}
	}
	w.choices[w.selected].SetColor(tl.ColorWhite, tl.ColorBlack)
}

func (w *welcome) Draw(screen *tl.Screen) {
	w.header.Draw(screen)
	w.header.Draw(screen)
	for _, v := range w.choices {
		v.Draw(screen)
	}
}

func newWelcome() *welcome {
	c := [4]string{"HOTSEAT", "P1 vs AI", "AI vs P2", "AI vs AI"}
	w := &welcome{
		header: tl.NewEntityFromCanvas(0, 0, welcomeCanvas),
		msg:    tl.NewText(15, 15, msg, tl.ColorBlack, tl.ColorWhite),
		choices: [4]*tl.Text{
			tl.NewText(20, 17, c[0], tl.ColorWhite, tl.ColorBlack),
			tl.NewText(20, 18, c[1], tl.ColorBlack, tl.ColorWhite),
			tl.NewText(20, 19, c[2], tl.ColorBlack, tl.ColorWhite),
			tl.NewText(20, 20, c[3], tl.ColorBlack, tl.ColorWhite),
		},
		done: false,
	}
	return w

}

func init() {
	welcomeCanvas = tl.CanvasFromString(msg)
}
