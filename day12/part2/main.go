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
	sides := 0
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
		} else {
			if !_map[y][x-1].Visited {
				queue = append(queue, []int{x - 1, y})
			}
		}

		if x+1 >= len(_map[0]) || _map[y][x+1].Type != region.Type {
		} else {
			if !_map[y][x+1].Visited {
				queue = append(queue, []int{x + 1, y})
			}
		}

		if y-1 < 0 || _map[y-1][x].Type != region.Type {
		} else {
			if !_map[y-1][x].Visited {
				queue = append(queue, []int{x, y - 1})
			}
		}

		if y+1 >= len(_map) || _map[y+1][x].Type != region.Type {
		} else {
			if !_map[y+1][x].Visited {
				queue = append(queue, []int{x, y + 1})
			}
		}

		sides += countSides(_map, x, y)

	}

	return sides * area
}

func countSides(_map [][]*Region, x, y int) int {
	sides := 0

	neighboursX := []*Region{}
	neighboursY := []*Region{}

	outboundY := y-1 < 0 || y+1 >= len(_map) || (_map[y-1][x].Type != _map[y][x].Type) || (_map[y+1][x].Type != _map[y][x].Type)
	outboundX := x-1 < 0 || x+1 >= len(_map) || (_map[y][x-1].Type != _map[y][x].Type) || (_map[y][x+1].Type != _map[y][x].Type)

	if x-1 >= 0 && _map[y][x-1].Type == _map[y][x].Type {
		neighboursX = append(neighboursX, _map[y][x-1])
	}
	if x+1 < len(_map[0]) && _map[y][x+1].Type == _map[y][x].Type {
		neighboursX = append(neighboursX, _map[y][x+1])
	}

	if y-1 >= 0 && _map[y-1][x].Type == _map[y][x].Type {
		neighboursY = append(neighboursY, _map[y-1][x])
	}

	if y+1 < len(_map) && _map[y+1][x].Type == _map[y][x].Type {
		neighboursY = append(neighboursY, _map[y+1][x])
	}

	if len(neighboursX)+len(neighboursY) == 1 {
		sides += 2
		return sides
	}

	if len(neighboursX)+len(neighboursY) == 0 {
		sides += 4
		return sides
	}

	if outboundY && outboundX {
		sides++
	}

	if (x-1 >= 0 && _map[y][x-1].Type == _map[y][x].Type) && y-1 >= 0 && _map[y-1][x].Type == _map[y][x].Type && _map[y-1][x-1].Type != _map[y][x].Type {
		sides++
	}

	if (x+1 < len(_map[0]) && _map[y][x+1].Type == _map[y][x].Type) && y-1 >= 0 && _map[y-1][x].Type == _map[y][x].Type && _map[y-1][x+1].Type != _map[y][x].Type {
		sides++
	}

	if (x-1 >= 0 && _map[y][x-1].Type == _map[y][x].Type) && y+1 < len(_map) && _map[y+1][x].Type == _map[y][x].Type && _map[y+1][x-1].Type != _map[y][x].Type {
		sides++
	}

	if (x+1 < len(_map[0]) && _map[y][x+1].Type == _map[y][x].Type) && y+1 < len(_map) && _map[y+1][x].Type == _map[y][x].Type && _map[y+1][x+1].Type != _map[y][x].Type {
		sides++
	}

	return sides
}
