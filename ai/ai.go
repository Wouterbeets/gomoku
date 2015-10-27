package ai

import (
	"fmt"
	"gomoku/conf"
	"gomoku/rules"
	"math"
	"math/rand"
	//	"sort"
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
	player int8
	heurF  func(*[boardSize][boardSize]int8, int8) []float64
}

type Board [boardSize][boardSize]int8

func (b Board) String() (str string) {
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

func New(depth int, comHud chan string, player int8) *Ai {
	ai := new(Ai)
	ai.initChans()
	ai.depth = depth
	ai.comHud = comHud
	//ai.SetHeur(Heur)
	ai.player = player
	return ai

}
func Heur(node *[boardSize][boardSize]int8, player int8) (count float64) {
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

func moveInLim(lims []int8, occ [][2]int8) bool {
	for i, _ := range occ {
		_ = "breakpoint"
		if occ[i][0] >= lims[0] && occ[i][0] <= lims[1] && occ[i][1] >= lims[2] && occ[i][1] <= lims[3] {
			return true
		} else if occ[i][0] > lims[1] {
			return false
		}
	}
	return false
}

func closeToOccupied(y, x int8, occ [][2]int8) bool {
	lims := []int8{y - 2, y + 2, x - 2, x + 2}
	if y < 2 {
		lims[0] = 0
	} else if y > 17 {
		lims[1] = 19
	}
	if x < 2 {
		lims[2] = 0
	} else if x > 17 {
		lims[3] = 19
	}
	//i := sort.Search(len(occ), func(i int) bool { return occ[i][0] >= lims[0] })
	//j := sort.Search(len(occ), func(i int) bool { return occ[i][0] < lims[0] })
	//if j > i {
	if moveInLim(lims, occ) {
		return false
	}
	//}
	return true
}

func shuffle(arr [][2]int8) [][2]int8 {
	for i := range arr {
		ran := rand.Intn(len(arr))
		arr[i], arr[ran] = arr[ran], arr[i]
	}
	return arr
}

func getMoves(node *[boardSize][boardSize]int8) (moves [][2]int8) {
	moves = make([][2]int8, 0, 30)
	occ := make([][2]int8, 0, 30)
	for y := int8(0); y < boardSize; y++ {
		for x := int8(0); x < boardSize; x++ {
			if node[y][x] != 0 {
				occ = append(occ, [2]int8{y, x})
			}
		}
	}
	if len(occ) == 0 {
		return [][2]int8{{9, 9}}
	}
	for y := int8(0); y < boardSize; y++ {
		for x := int8(0); x < boardSize; x++ {
			if node[y][x] == 0 {
				if !closeToOccupied(y, x, occ) {
					moves = append(moves, [2]int8{y, x})
				}
			}
		}
	}
	return shuffle(moves)
}

func (ai *Ai) miniMax(node *[boardSize][boardSize]int8, depth int, player int8) (best float64, y, x int8) {
	if depth == 0 /* or win or lose or no boardfull*/ {
		//heuristic
		sf := ai.heurF(node, player)
		return sf[0], 0, 0
	}
	if player == ai.p1 {
		best = 0.0
		moves := getMoves(node)
		fmt.Println(len(moves))
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
		fmt.Println(len(moves))
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
	//ai.comHud <- "starting"
	//	_, y, x := ai.miniMax(&board, depth, ai.p1)
	moves := ai.heurF(&board, ai.p1)
	movesIndex := []int{}
	bestVal := 0.0
	for k, v := range moves {
		if v > bestVal {
			bestVal = v
			movesIndex = []int{k}
		}
		if v == bestVal {
			movesIndex = append(movesIndex, k)
		}
	}
	bestMove := movesIndex[rand.Intn(len(movesIndex))]
	//fmt.Printf("bestmove %+v %+v %+v", bestMove, bestMove/19, bestMove%19)
	move[0] = int8(bestMove / 19)
	move[1] = int8(bestMove % 19)
	return
}

var log chan string

func (ai *Ai) SetHeur(f func(*[boardSize][boardSize]int8, int8) []float64) {
	ai.heurF = f
}

func (ai *Ai) Start() {
	rand.Seed(time.Now().UnixNano())
	log = ai.comHud
	c := <-conf.Conf
	//ai.comHud <- "got conf"
	ai.p1 = c.P1
	ai.p2 = c.P2
	for {
		b := <-ai.ComIn
		move := ai.Think(b, ai.depth)
		ai.ComOut <- move
	}
}

func (ai *Ai) Kill() {
	close(ai.ComIn)
	close(ai.ComOut)
	close(ai.comHud)
}
