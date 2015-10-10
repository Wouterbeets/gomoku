package rules

import (
	"errors"
)

const (
	boardSize = 19
)

func Check(board *[boardSize][boardSize]int8) error {
	//this is where we check the board to see if no rules need to be treated like
	//win
	//double three
	//placed on top of other stone
	//etc
	return errors.New("fuck")
}
