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
	robots := []*Robot{}

	utils.ReadLines("../data", func(line string) {
		input := utils.ParseLineInts(line)

		x, y, dx, dy := input[0], input[1], input[2], input[3]
		robots = append(robots, &Robot{X: x, Y: y, DX: dx, DY: dy})
	})

	run(robots, 100)
	// print(robots)

	fmt.Println(safetyFactor(robots))
}

func run(robots []*Robot, steps int) {
	for i := 0; i < steps; i++ {
		for _, r := range robots {
			r.Move()
		}
	}
}

func safetyFactor(robots []*Robot) int {
	quadrants := map[int]int{}

	for _, r := range robots {
		quadrants[r.Quadrant()]++
	}

	factor := quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]

	return factor
}

func print(robots []*Robot) {
	for y := 0; y < height; y++ {
		row := ""
		for x := 0; x < width; x++ {
			count := 0

			for _, robot := range robots {
				if robot.X == x && robot.Y == y {
					count++
				}
			}

			if count == 0 {
				row = fmt.Sprintf("%s.", row)
			} else {
				row = fmt.Sprintf("%s%d", row, count)
			}
		}
		fmt.Println(row)
	}
}
