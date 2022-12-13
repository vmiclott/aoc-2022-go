package day10

import "fmt"

type cpu struct {
	cycle           int
	x               int
	signalStrengths []int
	crt             crt
}

func (cpu *cpu) tick() {
	if debug {
		fmt.Println("tick")
	}
	cpu.crt.tick(cpu.cycle, cpu.x)
	cpu.cycle++
	if cpu.cycle%40 == 20 {
		cpu.signalStrengths = append(cpu.signalStrengths, cpu.cycle*cpu.x)
		if debug {
			fmt.Printf("signal strength at %d: %d x %d = %d\n", cpu.cycle, cpu.cycle, cpu.x, cpu.cycle*cpu.x)
		}
	}
}

func (cpu *cpu) noop() {
	if debug {
		fmt.Println("noop")
	}
	cpu.tick()
}

func (cpu *cpu) addX(val int) {
	if debug {
		fmt.Printf("addX %d\n", val)
	}
	cpu.tick()
	cpu.tick()
	cpu.x += val
}

func (cpu *cpu) totalSignalStrength() int {
	sum := 0
	for _, signalStrength := range cpu.signalStrengths {
		sum += signalStrength
	}
	return sum
}
