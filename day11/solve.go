package day11

import (
	"fmt"
)

var debug bool
var numberOfRounds int
var monkeysGetBored bool

var roundsToPrint = map[int]bool{
	0:    true,
	19:   true,
	999:  true,
	1999: true,
	2999: true,
	3999: true,
	4999: true,
	5999: true,
	6999: true,
	7999: true,
	8999: true,
	9999: true,
}

func solve(monkeys monkeys) int {
	for i := 0; i < numberOfRounds; i++ {
		if debug && roundsToPrint[i] {
			fmt.Printf("\nAFTER ROUND %d:\n", i+1)
		}
		for _, monkey := range monkeys {
			monkey.processAllItems(monkeys)
		}
		for _, monkey := range monkeys {
			if debug && roundsToPrint[i] {
				monkey.print()
			}
		}
	}
	return monkeys.monkeyBusiness()
}

func Solve(inputFileName string, d bool) error {
	debug = d

	// part 1
	mod = 1
	numberOfRounds = 20
	monkeysGetBored = true
	monkeys, err := parse(inputFileName)
	if err != nil {
		return err
	}
	res := solve(monkeys)
	fmt.Printf("Part one: %d\n", res)

	// part 2
	mod = 1
	numberOfRounds = 10000
	monkeysGetBored = false
	monkeys, err = parse(inputFileName)
	if err != nil {
		return err
	}
	res = solve(monkeys)
	fmt.Printf("Part two: %d\n", res)

	return nil
}
