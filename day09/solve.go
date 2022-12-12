package day09

import "fmt"

var debug bool

func solve(moves []move, ropeLength int) int {
	visited := map[position]bool{[2]int{0, 0}: true}
	var rope rope = createRope(ropeLength)

	for _, move := range moves {
		for i := 0; i < move.distance; i++ {
			moveRope(move.direction, rope)
			visited[rope[ropeLength-1]] = true
		}
	}

	return len(visited)
}

func Solve(inputFileName string, d bool) error {
	debug = d
	moves, err := parse(inputFileName)
	if err != nil {
		return err
	}
	fmt.Printf("Part one: %v\n", solve(moves, 2))
	fmt.Printf("Part two: %v\n", solve(moves, 10))
	return nil
}
