package main

import (
	"fmt"
	"strings"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

type Region struct {
	Type    string
	Visited bool
}

func main() {
	_map := [][]*Region{}

	utils.ReadLines("../data", func(line string) {
		plots := strings.Split(line, "")
		row := []*Region{}

		for _, plot := range plots {
			row = append(row, &Region{Type: plot, Visited: false})
		}

		_map = append(_map, row)
	})

	cost := 0
	for y := range _map {
		for x := range _map[y] {
			region := _map[y][x]

			if !region.Visited {
				cost += scanArea(_map, x, y)
			}
		}
	}

	fmt.Println(cost)
}

func scanArea(_map [][]*Region, x int, y int) int {

	queue := [][]int{{x, y}}

	perimeters := 0
	area := 0

	for len(queue) > 0 {
		coords := queue[0]
		x, y := coords[0], coords[1]
		queue = queue[1:]

		if _map[y][x].Visited {
			continue
		}

		area++

		region := _map[y][x]
		region.Visited = true

		if x-1 < 0 || _map[y][x-1].Type != region.Type {
			perimeters++
		} else {
			if !_map[y][x-1].Visited {
				queue = append(queue, []int{x - 1, y})
			}
		}

		if x+1 >= len(_map[0]) || _map[y][x+1].Type != region.Type {
			perimeters++
		} else {
			if !_map[y][x+1].Visited {
				queue = append(queue, []int{x + 1, y})
			}
		}

		if y-1 < 0 || _map[y-1][x].Type != region.Type {
			perimeters++
		} else {
			if !_map[y-1][x].Visited {
				queue = append(queue, []int{x, y - 1})
			}
		}

		if y+1 >= len(_map) || _map[y+1][x].Type != region.Type {
			perimeters++
		} else {
			if !_map[y+1][x].Visited {
				queue = append(queue, []int{x, y + 1})
			}
		}
	}

	return perimeters * area
}
