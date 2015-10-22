package ai

import (
	"fmt"
	"gomoku/conf"
	"gomoku/rules"
	"math"
	"math/rand"
	"time"
)

type mes struct {
	move [2]int8
}

const (
	boardSize  = 19
	boardTiles = 19 * 19
)

type Ai struct {
	ComOut chan [2]int8
	ComIn  chan [boardSize][boardSize]int8
	comHud chan string
	depth  int
	p1     int8
	p2     int8
	heurF  func(*[boardSize][boardSize]int8, int8) float64
}

type board [boardSize][boardSize]int8

func (b board) String() (str string) {
	for _, row := range b {
		for _, col := range row {
			str += fmt.Sprint(col, " ")
		}
		str += "\n"
	}
	return
}

func (ai *Ai) initChans() {
	ai.ComOut = make(chan [2]int8, 1)
	ai.ComIn = make(chan [boardSize][boardSize]int8, 1)
}

func New(depth int, comHud chan string) *Ai {
	ai := new(Ai)
	ai.initChans()
	ai.depth = depth
	ai.comHud = comHud
	ai.SetHeur(heur)
	return ai
}

func heur(node *[boardSize][boardSize]int8, player int8) (count float64) {
	for y := int8(0); y < boardSize; y++ {
		for x := int8(0); x < boardSize; x++ {
			if node[y][x] == player {
				_, score := rules.CheckWin(y, x, node)
				count += float64(score)
			}
		}
	}
	return
}

func getMoves(node *[boardSize][boardSize]int8) (moves [][2]int8) {
	moves = make([][2]int8, 0, 19*19)
	for y := int8(0); y < boardSize; y++ {
		for x := int8(0); x < boardSize; x++ {
			if node[y][x] == 0 {
				moves = append(moves, [2]int8{y, x})
			}
		}
	}
	return moves
}

func (ai *Ai) miniMax(node *[boardSize][boardSize]int8, depth int, player int8) (best float64, y, x int8) {
	if depth == 0 /* or win or lose or no boardfull*/ {
		//heuristic
		return ai.heurF(node, player), 2, 2
	}
	if player == ai.p1 {
		best = 0.0
		moves := getMoves(node)
		for _, move := range moves {
			node[move[0]][move[1]] = player
			newBest, _, _ := ai.miniMax(node, depth-1, ai.p2)
			if newBest > best {
				best, y, x = newBest, move[0], move[1]
			}
			node[move[0]][move[1]] = 0
		}
	} else if player == ai.p2 {
		best = math.MaxFloat64
		moves := getMoves(node)
		for _, move := range moves {
			node[move[0]][move[1]] = player
			newBest, _, _ := ai.miniMax(node, depth-1, ai.p1)
			if newBest < best {
				best, y, x = newBest, move[0], move[1]
			}
			node[move[0]][move[1]] = 0
		}
	}
	return
}

func (ai *Ai) Think(board [boardSize][boardSize]int8, depth int) (move [2]int8) {
	ai.comHud <- "starting"
	_, y, x := ai.miniMax(&board, depth, ai.p1)
	move[0] = y
	move[1] = x
	return
}

var log chan string

func (ai *Ai) SetHeur(f func(*[boardSize][boardSize]int8, int8) float64) {
	ai.heurF = f
}

func (ai *Ai) Start() {
	rand.Seed(time.Now().UnixNano())
	log = ai.comHud
	c := <-conf.Conf
	ai.comHud <- "got conf"
	ai.p1 = c.P1
	ai.p2 = c.P2
	for {
		b := <-ai.ComIn
		move := ai.Think(b, ai.depth)
		ai.ComOut <- move
	}
}
