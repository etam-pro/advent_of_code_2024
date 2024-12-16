package main

import (
	"fmt"
	"strings"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

const (
	E = 0
	S = 1
	W = 2
	N = 3
)

type Tracker map[[3]int]int

func NewTracker() Tracker {
	return Tracker{}
}

func (t Tracker) Add(x, y, dir, score int) Tracker {
	if t[[3]int{x, y, dir}] == 0 || t[[3]int{x, y, dir}] > score {
		t[[3]int{x, y, dir}] = score
	}
	return t
}

func (t Tracker) Get(x, y, dir int) int {
	return t[[3]int{x, y, dir}]
}

func (t Tracker) Visited(x, y, dir int) bool {
	_, ok := t[[3]int{x, y, dir}]
	return ok
}

type Map struct {
	MinScore  int
	BestPaths [][][2]int

	Grid [][]string
}

func NewMap(grid [][]string) *Map {
	return &Map{
		MinScore:  -1,
		BestPaths: [][][2]int{},
		Grid:      grid,
	}
}

func (m *Map) GetStart() (int, int) {
	for y, r := range m.Grid {
		for x, c := range r {
			if c == "S" {
				return x, y
			}
		}
	}

	return -1, -1
}

func (m *Map) Run(tracker Tracker) {
	x, y := m.GetStart()
	queue := [][]interface{}{{x, y, E, 0, tracker, [][2]int{}}}

	for len(queue) > 0 {
		input := queue[0]
		queue = queue[1:]

		x, y := input[0].(int), input[1].(int)
		dir := input[2].(int)
		score := input[3].(int)
		tracker := input[4].(Tracker)
		path := input[5].([][2]int)

		tracker = tracker.Add(x, y, dir, score)

		if m.Grid[y][x] == "E" {
			// Found a new best path, replace
			if m.MinScore == -1 || m.MinScore > score {
				m.MinScore = score
				endPath := copyPath(path)
				endPath = append(endPath, [2]int{x, y})
				m.BestPaths = [][][2]int{endPath}
			}

			// Found another path with the same score, append
			if m.MinScore == score {
				endPath := copyPath(path)
				endPath = append(endPath, [2]int{x, y})
				m.BestPaths = append(m.BestPaths, path)
			}
			continue
		}

		// If the current score is higher than the best score, skip
		if tracker.Get(x, y, dir) < score {
			continue
		}

		tox, toy := getTo(x, y, dir)
		if m.Grid[toy][tox] != "#" {
			newPath := copyPath(path)
			newPath = append(path, [2]int{x, y})
			queue = append(queue, []interface{}{tox, toy, dir, score + 1, tracker, newPath})
		}

		cdir := clockwise(dir)
		queue = append(queue, []interface{}{x, y, cdir, score + 1000, tracker, copyPath(path)})

		ccdir := counterClockwise(dir)
		queue = append(queue, []interface{}{x, y, ccdir, score + 1000, tracker, copyPath(path)})
	}
}

func getTo(x, y, dir int) (int, int) {
	tox, toy := x, y

	switch dir {
	case E:
		tox++
	case S:
		toy++
	case W:
		tox--
	case N:
		toy--
	}

	return tox, toy
}

func clockwise(dir int) int {
	return (dir + 1) % 4
}

func counterClockwise(dir int) int {
	return (dir + 3) % 4
}

func copyPath(path [][2]int) [][2]int {
	newPath := [][2]int{}
	newPath = append(newPath, path...)
	return newPath
}

func countPath(paths [][][2]int) int {
	_map := make(map[[2]int]bool)
	for _, path := range paths {
		for _, p := range path {
			_map[p] = true
		}
	}

	total := 0
	for range _map {
		total++
	}

	return total
}

func main() {
	grid := [][]string{}

	utils.ReadLines("../data", func(line string) {
		grid = append(grid, strings.Split(line, ""))
	})

	_map := NewMap(grid)
	_map.Run(NewTracker())

	fmt.Println(countPath(_map.BestPaths))
}
