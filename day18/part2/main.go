package main

import (
	"fmt"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

const (
	width  = 71
	height = 71
	bytes  = 1024
)

type Tracker map[[2]int]int

func NewTracker() Tracker {
	return Tracker{}
}

func copyTracker(t Tracker) Tracker {
	newTracker := NewTracker()
	for k, v := range t {
		newTracker[k] = v
	}
	return newTracker
}

type Map struct {
	Grid [][]string
}

func (m *Map) FindPath(start, end []int) int {
	fromx, fromy := start[0], start[1]
	tox, toy := end[0], end[1]

	queue := [][]interface{}{{fromx, fromy, 0}}

	visited := map[[2]int]int{}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		x, y, steps := item[0].(int), item[1].(int), item[2].(int)

		if x == tox && y == toy {
			return steps
		}

		if x < 0 || x >= width || y < 0 || y >= height {
			continue
		}

		if m.Grid[y][x] == "#" {
			continue
		}

		if visited[[2]int{x, y}] != 0 && visited[[2]int{x, y}] <= steps {
			continue
		}

		visited[[2]int{x, y}] = steps

		queue = append(queue, []interface{}{x + 1, y, steps + 1, visited})
		queue = append(queue, []interface{}{x - 1, y, steps + 1, visited})
		queue = append(queue, []interface{}{x, y + 1, steps + 1, visited})
		queue = append(queue, []interface{}{x, y - 1, steps + 1, visited})
	}

	return -1
}

func main() {
	xys := [][2]int{}

	utils.ReadLines("../data", func(line string) {
		xy := utils.ParseLineInts(line)
		xys = append(xys, [2]int{xy[0], xy[1]})
	})

	grid := [][]string{}
	for i := 0; i < height; i++ {
		row := []string{}
		for j := 0; j < width; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	for _, xy := range xys[:bytes] {
		grid[xy[1]][xy[0]] = "#"
	}

	_map := &Map{Grid: grid}

	x, y := -1, -1
	for _, xy := range xys[bytes:] {
		_map.Grid[xy[1]][xy[0]] = "#"
		steps := _map.FindPath([]int{0, 0}, []int{width - 1, height - 1})

		if steps == -1 {
			x, y = xy[0], xy[1]
			break
		}
	}

	fmt.Printf("%d,%d\n", x, y)
}
