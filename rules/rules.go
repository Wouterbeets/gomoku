package rules

import (
	"errors"
	"math"
)

const (
	UND = iota
	P1
	P2
	boardSize = 19
)

func diagUp(posY, posX int8, b *[boardSize][boardSize]int8) (int, error) {
	l, y, x, tile := 1, posY+1, posX-1, b[posY][posX]
	for ; y < boardSize && x > 0 && b[y][x] == tile; y, x = y+1, x-1 {
		l++
	}
	y, x = posY-1, posX+1
	for ; y >= 0 && x < boardSize && b[y][x] == tile; y, x = y-1, x+1 {
		l++
	}
	if l > 4 {
		return l, errors.New("win")
	}
	return l, nil
}

func diagDown(posY, posX int8, b *[boardSize][boardSize]int8) (int, error) {
	l, y, x, tile := 1, posY+1, posX+1, b[posY][posX]
	for ; y < boardSize && x < boardSize && b[y][x] == tile; y, x = y+1, x+1 {
		l++
	}
	y, x = posY-1, posX-1
	for ; y >= 0 && x >= 0 && b[y][x] == tile; y, x = y-1, x-1 {
		l++
	}
	if l > 4 {
		return l, errors.New("win")
	}
	return l, nil
}

func verti(posY, posX int8, b *[boardSize][boardSize]int8) (int, error) {
	l, y, x, tile := 1, posY+1, posX, b[posY][posX]
	for ; y < boardSize && b[y][x] == tile; y++ {
		l++
	}
	y, x = posY-1, posX
	for ; y >= 0 && b[y][x] == tile; y-- {
		l++
	}
	if l > 4 {
		return l, errors.New("win")
	}
	return l, nil
}

func horiz(posY, posX int8, b *[boardSize][boardSize]int8) (int, error) {
	l, y, x, tile := 1, posY, posX+1, b[posY][posX]
	for ; x < boardSize && b[y][x] == tile; x++ {
		l++
	}
	y, x = posY, posX-1
	for ; x >= 0 && b[y][x] == tile; x-- {
		l++
	}
	if l > 4 {
		return l, errors.New("win")
	}
	return l, nil
}

func checkDraw(b *[boardSize][boardSize]int8) error {
	for _, row := range b {
		for _, col := range row {
			if col == 0 {
				return nil
			}
		}
	}
	return errors.New("draw")
}

func CheckWin(y, x int8, b *[boardSize][boardSize]int8) (error, int) {
	checkDraw(b)
	sum := 0
	score := 0
	err := errors.New("")
	if score, err = horiz(y, x, b); err != nil {
		return err, math.MinInt64
	}
	sum += score
	if score, err = verti(y, x, b); err != nil {
		return err, math.MinInt64
	}
	sum += score
	if score, err = diagDown(y, x, b); err != nil {
		return err, math.MinInt64
	}
	sum += score
	if score, err = diagUp(y, x, b); err != nil {
		return err, math.MinInt64
	}
	sum += score
	return nil, sum
}

func CheckPos(y, x int8, b *[boardSize][boardSize]int8) error {
	err, _ := CheckWin(y, x, b)
	return err
}

func CheckOccupied(y, x int8, b *[boardSize][boardSize]int8) error {
	if b[y][x] != UND {
		return errors.New("already occupied")
	}
	return nil
}
