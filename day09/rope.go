package day09

import "fmt"

type position [2]int
type rope []position

const (
	nowhere   direction = ""
	upLeft    direction = "UL"
	upRight   direction = "UR"
	downLeft  direction = "DL"
	downRight direction = "DR"
)

func createRope(length int) rope {
	var rope rope
	for i := 0; i < length; i++ {
		rope = append(rope, position{0, 0})
	}
	return rope
}

func getMoveDirection(from position, to position) direction {
	difference := position{to[0] - from[0], to[1] - from[1]}
	if difference[0] < -1 {
		if difference[1] > 0 {
			return upLeft
		}
		if difference[1] < 0 {
			return downLeft
		}
		return left
	}

	if difference[0] > 1 {
		if difference[1] > 0 {
			return upRight
		}
		if difference[1] < 0 {
			return downRight
		}
		return right
	}

	if difference[1] < -1 {
		if difference[0] > 0 {
			return downRight
		}
		if difference[0] < 0 {
			return downLeft
		}
		return down
	}

	if difference[1] > 1 {
		if difference[0] > 0 {
			return upRight
		}
		if difference[0] < 0 {
			return upLeft
		}
		return up
	}
	return nowhere
}

func movePosition(direction direction, pos *position) {
	switch direction {
	case up:
		pos[1]++
	case down:
		pos[1]--
	case left:
		pos[0]--
	case right:
		pos[0]++
	case upLeft:
		pos[0]--
		pos[1]++
	case upRight:
		pos[0]++
		pos[1]++
	case downLeft:
		pos[0]--
		pos[1]--
	case downRight:
		pos[0]++
		pos[1]--
	}

}

func moveRope(direction direction, rope rope) {
	movePosition(direction, &rope[0])

	for i := 1; i < len(rope); i++ {
		movePosition(getMoveDirection(rope[i], rope[i-1]), &rope[i])
		if debug {
			fmt.Println(rope)
		}
	}
}
