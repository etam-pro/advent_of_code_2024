package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

type Tracker map[[2]int]bool

func main() {
	_map := [][]int{}

	utils.ReadLines("../data", func(line string) {
		_line := strings.Split(line, "")

		row := []int{}
		for _, cell := range _line {
			num, _ := strconv.Atoi(cell)
			row = append(row, num)
		}
		_map = append(_map, row)
	})

	trailheads := getTrailheads(_map)

	total := 0
	for _, th := range trailheads {
		tracker := map[[2]int]bool{}
		total += calc(_map, th, tracker)
	}

	fmt.Println(total)
}

func getTrailheads(_map [][]int) [][]int {
	trailheads := [][]int{}

	for y, row := range _map {
		for x, cell := range row {
			if cell == 0 {
				trailheads = append(trailheads, []int{x, y})
			}
		}
	}
	return trailheads
}

func calc(_map [][]int, xy []int, tracker Tracker) int {
	x, y := xy[0], xy[1]

	if _map[y][x] == 9 {
		if tracker[[2]int{x, y}] {
			return 0
		} else {
			tracker[[2]int{x, y}] = true
			return 1
		}
	}

	next := [][]int{}

	if isInBound(_map, x+1, y) && _map[y][x+1]-_map[y][x] == 1 {
		next = append(next, []int{x + 1, y})
	}

	if isInBound(_map, x-1, y) && _map[y][x-1]-_map[y][x] == 1 {
		next = append(next, []int{x - 1, y})
	}

	if isInBound(_map, x, y+1) && _map[y+1][x]-_map[y][x] == 1 {
		next = append(next, []int{x, y + 1})
	}

	if isInBound(_map, x, y-1) && _map[y-1][x]-_map[y][x] == 1 {
		next = append(next, []int{x, y - 1})
	}

	if len(next) == 0 {
		return 0
	}

	total := 0
	for _, n := range next {
		total += calc(_map, n, tracker)
	}

	return total
}

func isInBound(_map [][]int, x, y int) bool {
	return y >= 0 && y < len(_map) && x >= 0 && x < len(_map[y])
}
