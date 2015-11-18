package main

import (
	"errors"
	"fmt"
	"genetic"
	"gomoku/ai"
	"gomoku/conf"
	"gomoku/rules"
	"math/rand"
	"nn"
	"time"
)

type Arena struct {
	ais          [2]*ai.Ai
	com1         chan string
	com2         chan string
	board        [19][19]int8
	p1           int8
	p2           int8
	indexAi1     int
	indexAi2     int
	gamesPlayed1 int
	gamesPlayed2 int
}

func (a *Arena) String() (str string) {
	for _, row := range a.board {
		str += fmt.Sprintln(row)
	}
	str += "\n"
	return str
}

func brainHeur(brain *nn.Net) func(*[19][19]int8, int8) []float64 {
	inp := make([]float64, 19*19)
	return func(board *[19][19]int8, player int8) []float64 {
		for y, row := range board {
			for x, col := range row {
				if col != 0 {
					if col == player {
						col = 1
					} else {
						col = 2
					}
				}
				inp[y*19+x] = float64(col)
			}
		}
		brain.In(inp)
		heur := brain.Out()
		return heur
	}
}

func newArena(brain1 *nn.Net, brain2 *nn.Net) *Arena {
	if brain1 == nil || brain2 == nil {
		fmt.Println("no brain in arena")
		return nil
	}
	conf := <-conf.Conf
	com1 := make(chan string)
	com2 := make(chan string)
	a := Arena{
		ais: [2]*ai.Ai{
			ai.New(conf.Depth, com1, conf.P1),
			ai.New(conf.Depth, com2, conf.P2),
		},
		com1: com1,
		com2: com2,
		p1:   conf.P1,
		p2:   conf.P2,
	}
	a.ais[0].SetHeur(brainHeur(brain1))
	go a.ais[0].Start()
	a.ais[1].SetHeur(brainHeur(brain2))
	go a.ais[1].Start()
	return &a
}

func (a *Arena) move(player int8) error {
	a.ais[player-1].ComIn <- a.board
	move := <-a.ais[player-1].ComOut
	if a.board[move[0]][move[1]] != 0 {
		y, x := rand.Intn(19), rand.Intn(19)
		for a.board[y][x] != 0 {
			y, x = rand.Intn(19), rand.Intn(19)
		}
		a.board[y][x] = player
		win, _ := rules.CheckWin(move[0], move[1], &a.board)
		if win != nil {
			return errors.New("draw")
		}
		return errors.New("space already occupied")
	}
	a.board[move[0]][move[1]] = player
	win, _ := rules.CheckWin(move[0], move[1], &a.board)
	if win != nil {
		return win
	}
	return nil
}

func (a *Arena) Fight(gen, i int) (int, int) {
	turns1 := 0
	turns2 := 0
	doubles1 := 0
	doubles2 := 0
	for {
		turns1++
		for gameState := a.move(1); gameState != nil; {
			errMsg := gameState.Error()
			if errMsg == "win" {
				return 361 - doubles1 - turns1, turns2 - doubles2
			} else if errMsg == "draw" {
				return turns1 - doubles1, turns2 - doubles2
			} else if errMsg == "space already occupied" {
				doubles1++
				gameState = nil
			}
		}
		//fmt.Println("player one genereation played", gen, "\n", a)
		turns2++
		for gameState := a.move(2); gameState != nil; {
			errMsg := gameState.Error()
			if errMsg == "win" {
				return turns1 - doubles1, 361 - turns2 - doubles2
			} else if errMsg == "draw" {
				return turns1 - doubles1, turns2 - doubles2
			} else if errMsg == "space already occupied" {
				doubles2++
				gameState = nil
			}
		}
	}
}

//this function is set as the gen package's fight function, it calculates the fitness of the ais
func ArenaFightFunc(ais []*gen.Ai, gen int) {
	arenas := make([]*Arena, len(ais))
	for i := 0; i < len(ais); i++ {
		for j := i + 1; j < len(ais); j++ {
			arenas[i] = newArena(ais[i].Net, ais[j].Net)
			arenas[i].indexAi1 = i
			arenas[i].indexAi2 = j
			score1, score2 := arenas[i].Fight(gen, i)
			fmt.Println(i, j, "fight", score1, score2)
			ais[arenas[i].indexAi1].Score += float64(score1)
			ais[arenas[i].indexAi1].GamesPlayed++
			ais[arenas[i].indexAi2].Score += float64(score2)
			ais[arenas[i].indexAi2].GamesPlayed++
		}
	}
	for _, ai := range ais {
		fmt.Print(ai.Score, " ")
		ai.Score /= ai.GamesPlayed
		fmt.Println(ai.Score, ai.GamesPlayed)
	}
	fmt.Println("")
}

func main() {

	c := conf.Config{
		Depth: 1,
		P1:    1,
		P2:    2,
	}
	conf.SetConf <- c
	pool := gen.CreatePool(10, 0.01, 10, 19*19, 50, 3, 19*19)
	pool.FightFunc = ArenaFightFunc
	pool.Evolve(300000, nil, nil)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
