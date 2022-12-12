package day09

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type direction string

const (
	up    direction = "U"
	down  direction = "D"
	left  direction = "L"
	right direction = "R"
)

type move struct {
	direction direction
	distance  int
}

func parse(inputFileName string) ([]move, error) {
	readFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var moves = []move{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction := direction(line[0])
		distance, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}
		moves = append(moves, move{direction, distance})
	}

	return moves, nil
}
