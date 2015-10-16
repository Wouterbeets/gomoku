package ai

import (
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

type ai struct {
	ComOut chan [2]int8
	ComIn  chan [boardSize][boardSize]int8
}

func (ai *ai) initChans() {
	ai.ComOut = make(chan [2]int8, 1)
	ai.ComIn = make(chan [boardSize][boardSize]int8, 1)
}

func New() *ai {
	ai := new(ai)
	ai.initChans()
	return ai
}

func (ai *ai) Think(board [boardSize][boardSize]int8) (move [2]int8) {
	move[0] = int8(rand.Intn(boardSize))
	move[1] = int8(rand.Intn(boardSize))
	return
}

func (ai *ai) Start() {
	rand.Seed(time.Now().UnixNano())
	for {
		board := <-ai.ComIn
		ai.ComOut <- ai.Think(board)
	}
}
