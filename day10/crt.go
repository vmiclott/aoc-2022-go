package day10

import "fmt"

const NUM_ROWS = 6
const NUM_COLS = 40

type pixel string

const (
	dark pixel = "."
	lit  pixel = "#"
)

type crt [NUM_COLS][NUM_ROWS]pixel

func (crt *crt) draw() {
	for row := 0; row < NUM_ROWS; row++ {
		for col := 0; col < NUM_COLS; col++ {
			fmt.Print(crt[col][row])
		}
		fmt.Println()
	}
}

func (crt *crt) tick(cycle int, x int) error {
	col := cycle % NUM_COLS
	row := cycle / NUM_COLS
	if row >= NUM_ROWS || col >= NUM_COLS {
		return fmt.Errorf("cannot add another pixel, crt is full")
	}

	if abs(x-col) < 2 {
		crt[col][row] = lit
	} else {
		crt[col][row] = dark
	}
	return nil
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
