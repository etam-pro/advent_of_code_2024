package main

import (
	"fmt"
	"strings"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	_map := [][]string{}

	utils.ReadLines("../data", func(line string) {
		_map = append(_map, strings.Split(line, ""))
	})

	x, y := getGuardPosition(_map)

	ob := false
	for !ob {
		_map, x, y, ob = move(_map, x, y)
	}

	fmt.Println(count(_map))
}

func move(_map [][]string, x, y int) ([][]string, int, int, bool) {
	direction := _map[y][x]
	_map[y][x] = "x"

	if outOfBounds(_map, x, y, direction) {
		return _map, x, y, true
	}

	var newX, newY int
	var newDirection string

	switch direction {
	case "^":
		if _map[y-1][x] == "#" {
			newDirection = ">"
			newX = x + 1
			newY = y
		} else {
			newX = x
			newY = y - 1
			newDirection = "^"
		}
	case "v":
		if _map[y+1][x] == "#" {
			newDirection = "<"
			newX = x - 1
			newY = y
		} else {
			newX = x
			newY = y + 1
			newDirection = "v"
		}
	case "<":
		if _map[y][x-1] == "#" {
			newDirection = "^"
			newX = x
			newY = y - 1
		} else {
			newX = x - 1
			newY = y
			newDirection = "<"
		}
	case ">":
		if _map[y][x+1] == "#" {
			newDirection = "v"
			newX = x
			newY = y + 1
		} else {
			newX = x + 1
			newY = y
			newDirection = ">"
		}
	}

	_map[newY][newX] = newDirection

	return _map, newX, newY, false
}

func outOfBounds(_map [][]string, x, y int, direction string) bool {
	var ob bool

	switch direction {
	case "^":
		ob = y-1 < 0
	case "v":
		ob = y+1 >= len(_map)
	case "<":
		ob = x-1 < 0
	case ">":
		ob = x+1 >= len(_map[0])
	}

	return ob
}

func getDirection(guard string) string {
	switch guard {
	case "v":
		return "down"
	case "^":
		return "up"
	case "<":
		return "left"
	case ">":
		return "right"
	}

	return ""
}

func getGuardPosition(_map [][]string) (int, int) {
	for y, row := range _map {
		for x, cell := range row {
			if cell == "^" || cell == "v" || cell == "<" || cell == ">" {
				return x, y
			}
		}
	}

	return -1, -1
}

func count(_map [][]string) int {
	count := 0

	for _, row := range _map {
		for _, cell := range row {
			if cell == "x" {
				count++
			}
		}
	}

	return count
}

func printMap(_map [][]string) {
	for _, row := range _map {
		fmt.Println(strings.Join(row, ""))
	}
}
