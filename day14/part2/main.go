package main

import (
	"fmt"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

const (
	// width  = 11
	// height = 7

	width  = 101
	height = 103
)

type Robot struct {
	X  int
	Y  int
	DX int
	DY int
}

func (r *Robot) Move() {
	r.X += r.DX
	r.Y += r.DY

	if r.X < 0 {
		r.X += width
	}

	if r.X >= width {
		r.X -= width
	}

	if r.Y < 0 {
		r.Y += height
	}

	if r.Y >= height {
		r.Y -= height
	}
}

func (r *Robot) Quadrant() int {
	if r.X == width/2 || r.Y == height/2 {
		return 0
	}

	if r.X < width/2 && r.Y < height/2 {
		return 1
	}

	if r.X > width/2 && r.Y < height/2 {
		return 2
	}

	if r.X < width/2 && r.Y > height/2 {
		return 3
	}

	return 4
}

func main() {
	_map := initMap()
	robots := []*Robot{}

	utils.ReadLines("../data", func(line string) {
		input := utils.ParseLineInts(line)

		x, y, dx, dy := input[0], input[1], input[2], input[3]
		robots = append(robots, &Robot{X: x, Y: y, DX: dx, DY: dy})
		_map[[2]int{x, y}] += 1
	})

	secs := run(robots, _map)
	print(_map)
	fmt.Println(secs)
}

func move(r *Robot, _map map[[2]int]int) {
	_map[[2]int{r.X, r.Y}] -= 1
	r.Move()
	_map[[2]int{r.X, r.Y}] += 1
}

func run(robots []*Robot, _map map[[2]int]int) int {
	seconds := 0

	for {
		seconds++

		for _, r := range robots {
			move(r, _map)
		}

		if check(_map) {
			break
		}
	}

	return seconds
}

func check(_map map[[2]int]int) bool {
	adjacents := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if _map[[2]int{x, y}] > 1 {
				return false
			}

			if x+1 >= width {
				continue
			}

			if _map[[2]int{x, y}] == 1 && _map[[2]int{x + 1, y}] == 1 {
				adjacents++
			}
		}
	}

	return adjacents > 200
}

func safetyFactor(_map map[[2]int]int) int {
	quadrants := map[int]int{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			quadrants[quadrant(x, y)] += _map[[2]int{x, y}]
		}
	}

	factor := quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]

	return factor
}

func print(_map map[[2]int]int) {

	for y := 0; y < height; y++ {
		row := ""
		for x := 0; x < width; x++ {
			count := _map[[2]int{x, y}]

			if count == 0 {
				row = fmt.Sprintf("%s.", row)
			} else {
				row = fmt.Sprintf("%s%d", row, count)
			}
		}
		fmt.Println(row)
	}
}

func initMap() map[[2]int]int {
	_map := map[[2]int]int{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_map[[2]int{x, y}] = 0
		}
	}

	return _map
}

func quadrant(x, y int) int {
	if x == width/2 || y == height/2 {
		return 0
	}

	if x < width/2 && y < height/2 {
		return 1
	}

	if x > width/2 && y < height/2 {
		return 2
	}

	if x < width/2 && y > height/2 {
		return 3
	}

	return 4
}
