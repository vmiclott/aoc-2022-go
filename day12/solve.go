package day12

import (
	"container/heap"
	"fmt"
)

var debug bool

func solve(inputFileName string, endRunes map[rune]bool) (int, error) {
	g, pq, err := parse(inputFileName)
	if err != nil {
		return 0, err
	}
	if debug {
		g.print()
	}
	visited := [][]string{}
	for row := range g {
		visited = append(visited, []string{})
		for range g[row] {
			visited[row] = append(visited[row], ".")
		}
	}
	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*item[*node]).value
		visited[n.row][n.col] = string(n.r)
		for _, neighbor := range n.getNeighbors(g) {
			if neighbor.distance > n.distance+1 {
				neighbor.distance = n.distance + 1
				heap.Push(&pq, &item[*node]{value: neighbor, priority: neighbor.distance})
			}
			if endRunes[neighbor.r] {
				return neighbor.distance, nil
			}
		}
	}
	if debug {
		for row := 0; row < len(visited); row++ {
			for col := 0; col < len(visited[row]); col++ {
				fmt.Print(visited[row][col])
			}
			fmt.Println()
		}
	}
	return 0, fmt.Errorf("can't reach start")
}

func Solve(inputFileName string, d bool) error {
	debug = d

	// part 1
	res, err := solve(inputFileName, map[rune]bool{'S': true})
	if err != nil {
		return err
	}
	fmt.Printf("Part one: %d\n", res)

	// part 2
	res, err = solve(inputFileName, map[rune]bool{'a': true, 'S': true})
	if err != nil {
		return err
	}
	fmt.Printf("Part two: %d\n", res)

	return nil
}
