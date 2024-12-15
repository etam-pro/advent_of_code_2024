package main

import (
	"fmt"
	"strings"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	_map := [][]string{}

	steps := []string{}

	isParsingMap := true
	utils.ReadLines("../data", func(line string) {
		if line == "" {
			isParsingMap = false
		}

		if isParsingMap {
			_map = append(_map, strings.Split(line, ""))
		} else {
			steps = append(steps, strings.Split(line, "")...)
		}
	})

	x, y := findRobot(_map)

	for _, step := range steps {
		x, y = move(_map, x, y, step)
	}

	total := 0
	for _, box := range findBoxes(_map) {
		total += box[0] + box[1]*100
	}

	fmt.Println(total)
}

func findRobot(_map [][]string) (int, int) {
	for y := range _map {
		for x := range _map[y] {
			if _map[y][x] == "@" {
				return x, y
			}
		}
	}
	return -1, -1
}

func findBoxes(_map [][]string) [][]int {
	boxes := [][]int{}
	for y := range _map {
		for x := range _map[y] {
			if _map[y][x] == "O" {
				boxes = append(boxes, []int{x, y})
			}
		}
	}
	return boxes
}

func move(_map [][]string, x, y int, step string) (int, int) {
	can := canMove(_map, x, y, step)
	if !can {
		return x, y
	}

	switch step {
	case ">":
		return moveRight(_map, x, y)
	case "<":
		return moveLeft(_map, x, y)
	case "^":
		return moveUp(_map, x, y)
	case "v":
		return moveDown(_map, x, y)
	}

	return x, y
}

func moveRight(_map [][]string, x, y int) (int, int) {
	tox := x + 1
	for i := tox; i > x; i-- {
		_map[y][i] = _map[y][i-1]
	}

	_map[y][x] = "."

	return x + 1, y
}

func moveLeft(_map [][]string, x, y int) (int, int) {
	tox := x - 1
	for i := tox; i < x; i++ {
		_map[y][i] = _map[y][i+1]
	}

	_map[y][x] = "."

	return x - 1, y
}

func moveUp(_map [][]string, x, y int) (int, int) {
	toy := y - 1
	for i := toy; i < y; i++ {
		_map[i][x] = _map[i+1][x]
	}

	_map[y][x] = "."

	return x, y - 1
}

func moveDown(_map [][]string, x, y int) (int, int) {
	toy := y + 1
	for i := toy; i > y; i-- {
		_map[i][x] = _map[i-1][x]
	}

	_map[y][x] = "."

	return x, y + 1
}

func canMove(_map [][]string, x int, y int, step string) bool {
	switch step {
	case ">":
		return canMoveRight(_map, x, y)
	case "<":
		return canMoveLeft(_map, x, y)
	case "^":
		return canMoveUp(_map, x, y)
	case "v":
		return canMoveDown(_map, x, y)
	}

	return false
}

func canMoveRight(_map [][]string, x int, y int) bool {
	for i := x; i < len(_map[y]); i++ {
		if _map[y][i] == "#" {
			break
		}

		if _map[y][i] == "." {
			return true
		}
	}

	return false
}

func canMoveLeft(_map [][]string, x int, y int) bool {
	for i := x; i >= 0; i-- {
		if _map[y][i] == "#" {
			break
		}

		if _map[y][i] == "." {
			return true
		}
	}

	return false
}

func canMoveUp(_map [][]string, x int, y int) bool {
	for i := y; i >= 0; i-- {
		if _map[i][x] == "#" {
			break
		}

		if _map[i][x] == "." {
			return true
		}
	}

	return false
}

func canMoveDown(_map [][]string, x int, y int) bool {
	for i := y; i < len(_map); i++ {
		if _map[i][x] == "#" {
			break
		}

		if _map[i][x] == "." {
			return true
		}
	}

	return false
}
