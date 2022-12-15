package day12

import (
	"fmt"
)

type node struct {
	r        rune
	distance int
	row      int
	col      int
}

func (n node) isReachableNeighbor(neighbor node) bool {
	a, b := n.r, neighbor.r
	if a == 'S' {
		a = 'a'
	}
	if b == 'S' {
		b = 'a'
	}
	if a == 'E' {
		a = 'z'
	}
	if b == 'E' {
		b = 'z'
	}

	return b-a > -2
}

var neighborPositions [4][2]int = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func (n node) getNeighbors(g graph) []*node {
	row, col := n.row, n.col
	res := []*node{}
	for _, pos := range neighborPositions {
		if g.isInBounds(row+pos[0], col+pos[1]) {
			neighbor := g[row+pos[0]][col+pos[1]]
			if n.isReachableNeighbor(*neighbor) {
				res = append(res, neighbor)
			}
		}
	}
	return res
}

type graph [][]*node

func (g graph) print() {
	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			fmt.Print(string(g[row][col].r))
		}
		fmt.Println()
	}
}

func (g graph) isInBounds(row int, col int) bool {
	return 0 <= row && row < len(g) && 0 <= col && col < len(g[row])
}
