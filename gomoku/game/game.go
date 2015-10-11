package game

import (
	"gomoku/rules"
	tl "termloop"
)

const (
	UND = iota
	P1
	P2
	AI1
	AI2
	cornerUL rune = '┌'
	cornerUR rune = '┐'
	cornerLL rune = '└'
	cornerLR rune = '┘'

	hDash rune = '─'
	vDash rune = '│'

	blank  rune = ' '
	circle rune = '●'

	tileSizeX  = 5
	tileSizeY  = 3
	boardSize  = 19
	boardTiles = 19 * 19
	tileStr    = "┌───┐\n│   │\n└───┘"
)

const (
	START = iota
	PLAY
	WINP1
	WINP2
)

var (
	Bg tl.Attr = tl.RgbTo256Color(60, 150, 180)
	Fg tl.Attr = tl.RgbTo256Color(92, 64, 51)
)

type board struct {
	tiles     [boardSize][boardSize]int8
	tilesDisp [boardSize][boardSize]*tile
	sY        int8 //selected Y
	sX        int8 //selected X
	offsetX   int
	offsetY   int
	level     *tl.BaseLevel
	turn      int8 //current turn b.player1 or b.player2
	player1   int8 // p1 or ai1
	player2   int8 // p2 or ai2
	ComOut    chan [boardSize][boardSize]int8
	ComIn     chan [2]int8
	screen    *tl.Screen
	wel       *welcome
	state     int8
	comHud    chan string
}

func (b *board) setPiece() {
	if err := rules.Check(b.sY, b.sX, &b.tiles); err != nil {
		b.comHud <- err.Error()
		if b.turn == AI1 || b.turn == AI2 {
			b.ComOut <- b.tiles
		}
		return
	}
	if b.turn == P1 || b.turn == AI1 {
		b.tilesDisp[b.sY][b.sX].black()
		b.tiles[b.sY][b.sX] = b.player1
		b.turn = b.player2
		if b.player2 == AI2 {
			b.ComOut <- b.tiles
		}
	} else {
		b.tilesDisp[b.sY][b.sX].white()
		b.tiles[b.sY][b.sX] = b.player2
		b.turn = b.player1
		if b.player1 == AI1 {
			b.ComOut <- b.tiles
		}
	}
	if err := rules.CheckWin(b.sY, b.sX, &b.tiles); err != nil {
		if b.turn == b.player1 {
			b.state = WINP2
			b.comHud <- err.Error() + "player 1"
		} else {
			b.state = WINP1
			b.comHud <- err.Error() + "player 2"
		}
	}
}

func newBoard(level *tl.BaseLevel, screen *tl.Screen, player1 int8, player2 int8, comIn chan [2]int8, comOut chan [boardSize][boardSize]int8, wel *welcome) (b *board) {
	b = new(board)
	b.addTiles()
	b.level = level
	b.player1 = player1
	b.player2 = player2
	b.ComIn = comIn
	b.ComOut = comOut
	b.screen = screen
	b.wel = wel
	return
}

func (b *board) addTiles() {
	for i := 0; i < boardTiles; i++ {
		y, x := i/boardSize, i%boardSize
		t := newTile(y, x, b.offsetY, b.offsetX)
		b.tilesDisp[y][x] = t
	}
}

func (b *board) Draw(screen *tl.Screen) {
	if b.wel.done {
		b.initPlayers()
		w, h := screen.Size()
		b.level.SetOffset(w/2-((boardSize*tileSizeX)/2), h/2-((boardSize*tileSizeY)/2))
		b.tilesDisp[b.sY][b.sX].selected()
		if b.turn == AI1 || b.turn == AI2 {
			b.handleAIInput()
		}
		for _, tY := range b.tilesDisp {
			for _, tX := range tY {
				tX.Draw(screen)
			}
		}
	}
	//TODO if b.state == WINP1 || b.state == WINP2 {
	// highlight win
}

func (b *board) handleHumanInput(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			if b.sX < boardSize-1 {
				b.tilesDisp[b.sY][b.sX].deselect()
				b.sX++
			}
		case tl.KeyArrowLeft:
			if b.sX > 0 {
				b.tilesDisp[b.sY][b.sX].deselect()
				b.sX--
			}
		case tl.KeyArrowUp:
			if b.sY > 0 {
				b.tilesDisp[b.sY][b.sX].deselect()
				b.sY--
			}
		case tl.KeyArrowDown:
			if b.sY < boardSize-1 {
				b.tilesDisp[b.sY][b.sX].deselect()
				b.sY++
			}
		case tl.KeySpace:
			if b.state == PLAY {
				b.setPiece()
			} else if b.state == START {
				b.state = PLAY
			}
		}
	}
}

func (b *board) handleAIInput() {
	b.comHud <- "your turn human.. good luck, you'll need it"
	in := <-b.ComIn
	b.tilesDisp[b.sY][b.sX].deselect()
	b.sY = in[0]
	b.sX = in[1]
	b.setPiece()
}

func (b *board) initPlayers() {
	if b.wel.done && b.wel.selected < 4 {
		switch b.wel.selected {
		case 0:
			b.player1 = P1
			b.player2 = P2
		case 1:
			b.player1 = P1
			b.player2 = AI2
		case 2:
			b.player1 = AI1
			b.player2 = P2
			b.ComOut <- b.tiles
		case 3:
			b.player1 = AI1
			b.player2 = AI2
			b.ComOut <- b.tiles
		}
		b.wel.selected = 4
		b.turn = b.player1
	}
}

func (b *board) Tick(event tl.Event) {
	if b.wel.done {
		b.handleHumanInput(event)
	}
}

//check rules with rules package

func Start(comOut chan [boardSize][boardSize]int8, comIn chan [2]int8) {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: Bg,
		Fg: Fg,
	})
	w := newWelcome(game)
	level.AddEntity(w)
	b := newBoard(level, game.Screen(), P1, AI2, comIn, comOut, w)
	level.AddEntity(b)
	comHud := make(chan string, 200)
	h := newHud(comHud, game.Screen())
	b.comHud = comHud
	game.Screen().AddEntity(h)
	game.Screen().SetLevel(level)
	game.Start()
}
