package day10

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type commandType string

const (
	noop commandType = "noop"
	addX commandType = "addx"
)

type command struct {
	commandType commandType
	val         int
}

func parse(inputFileName string) ([]command, error) {
	readFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var commands = []command{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		switch commandType(line[0]) {
		case noop:
			commands = append(commands, command{noop, 0})
		case addX:
			val, err := strconv.Atoi(line[1])
			if err != nil {
				return nil, err
			}
			commands = append(commands, command{addX, val})
		}
	}

	return commands, nil
}
