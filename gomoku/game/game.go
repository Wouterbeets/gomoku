package game

import (
	//"fmt"
	tl "termloop"
)

const (
	cornerUL rune = '┌'
	cornerUR rune = '┐'
	cornerLL rune = '└'
	cornerLR rune = '┘'

	hDash rune = '─'
	vDash rune = '│'

	blank  rune = ' '
	circle      = '●'

	tileSizeX  = 5
	tileSizeY  = 3
	boardSize  = 19
	boardTiles = 19 * 19
	tileStr    = "┌───┐\n│   │\n└───┘"
	P1         = iota
	P2
	UND
	AI1
	AI2
)

type board struct {
	tiles     [boardSize][boardSize]int8
	tilesDisp [boardSize][boardSize]*tile
	sY        int8
	sX        int8
	offsetX   int
	offsetY   int
	level     *tl.BaseLevel
	turn      int8 //current turn b.player1 or b.player2
	player1   int8 // p1 or ai1
	player2   int8 // p2 or ai2
	ComOut    chan [boardSize][boardSize]int8
	ComIn     chan [2]int8
	screen    *tl.Screen
}

func (b *board) setPiece() {
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
}

func (b *board) sendBoard() (out [boardSize][boardSize]int8) {
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			out[y][x] = b.tiles[y][x]
		}
	}
	return
}

func newBoard(offestX, offsetY int, level *tl.BaseLevel, screen *tl.Screen, player1 int8, player2 int8, comIn chan [2]int8, comOut chan [boardSize][boardSize]int8) (b *board) {
	b = new(board)
	b.offsetX = offestX
	b.offsetY = offsetY
	b.addTiles()
	b.level = level
	b.player1 = player1
	b.player2 = player2
	b.ComIn = comIn
	b.ComOut = comOut
	b.screen = screen
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
	w, h := screen.Size()
	b.level.SetOffset(w/2-((boardSize*tileSizeX)/2), h/2-((boardSize*tileSizeY)/2))
	b.tilesDisp[b.sY][b.sX].selected()
	for _, tY := range b.tilesDisp {
		for _, tX := range tY {
			tX.Draw(screen)
		}
	}
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
			b.setPiece()
		}
	}
}

func (b *board) handleAIInput() {
	in := <-b.ComIn
	b.tilesDisp[b.sY][b.sX].deselect()
	b.sY = in[0]
	b.sX = in[1]
	b.setPiece()
}

func (b *board) Tick(event tl.Event) {
	if b.turn == P1 || b.turn == P2 {
		b.handleHumanInput(event)
	} else {
		b.handleAIInput()
	}
	//check rules with rules package
}

func Start(comOut chan [boardSize][boardSize]int8, comIn chan [2]int8) {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.RgbTo256Color(60, 150, 180),
		Fg: tl.RgbTo256Color(92, 64, 51),
	})
	game.Screen().SetLevel(level)
	b := newBoard(0, 0, level, game.Screen(), P1, AI2, comIn, comOut)
	level.AddEntity(b)
	fps := tl.NewFpsText(0, 0, tl.ColorBlack, tl.ColorRed, 1)
	level.AddEntity(fps)
	b.turn = b.player1
	b.ComIn <- [2]int8{10, 9}
	game.Start()
}
