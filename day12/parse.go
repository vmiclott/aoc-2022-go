package day12

import (
	"bufio"
	"container/heap"
	"io"
	"math"
	"os"
)

func parse(inputFileName string) (graph, priorityQueue[*node], error) {
	readFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, nil, err
	}
	defer readFile.Close()

	reader := bufio.NewReader(readFile)

	row := 0
	col := 0

	g := graph{}
	g = append(g, []*node{})
	pq := priorityQueue[*node]{}
	heap.Init(&pq)

	for {
		if r, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, nil, err
			}
		} else if r == '\n' {
			g = append(g, []*node{})
			row++
			col = 0
		} else {
			n := &node{r: r, distance: math.MaxInt, row: row, col: col}
			if r == 'E' {
				n.distance = 0
				heap.Push(&pq, &item[*node]{value: n, priority: 0})
			}
			g[row] = append(g[row], n)
			col++
		}
	}

	return g, pq, nil
}
