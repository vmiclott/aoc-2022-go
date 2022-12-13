package day10

import "fmt"

var debug bool

func solve(commands []command) {
	cpu := cpu{x: 1}
	for _, command := range commands {
		switch command.commandType {
		case noop:
			cpu.noop()
		case addX:
			cpu.addX(command.val)
		}
	}

	fmt.Printf("Part one: %v\n", cpu.totalSignalStrength())
	fmt.Printf("Part two:\n")
	cpu.crt.draw()
}

func Solve(inputFileName string, d bool) error {
	debug = d
	commands, err := parse(inputFileName)
	if err != nil {
		return err
	}
	solve(commands)
	return nil
}
