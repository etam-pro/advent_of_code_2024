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
			grids := strings.Split(line, "")
			row := []string{}
			for _, g := range grids {
				if g == "@" {
					row = append(row, []string{"@", "."}...)
				} else if g == "O" {
					row = append(row, []string{"[", "]"}...)
				} else {
					row = append(row, []string{g, g}...)
				}
			}
			_map = append(_map, row)
		} else {
			steps = append(steps, strings.Split(line, "")...)
		}
	})

	x, y := findRobot(_map)

	utils.PrintMap(_map)
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
			if _map[y][x] == "[" {
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

func moveUp(_map [][]string, x, y int) (int, int) {
	tox, toy := x, y

	toy--

	if _map[toy][tox] == "[" {
		moveUp(_map, tox, toy)
		moveUp(_map, tox+1, toy)
	}

	if _map[toy][tox] == "]" {
		moveUp(_map, tox, toy)
		moveUp(_map, tox-1, toy)
	}

	_map[toy][tox] = _map[y][x]
	_map[y][x] = "."

	return tox, toy
}

func moveDown(_map [][]string, x, y int) (int, int) {
	tox, toy := x, y

	toy++

	if _map[toy][tox] == "[" {
		moveDown(_map, tox, toy)
		moveDown(_map, tox+1, toy)
	}

	if _map[toy][tox] == "]" {
		moveDown(_map, tox, toy)
		moveDown(_map, tox-1, toy)
	}

	_map[toy][tox] = _map[y][x]
	_map[y][x] = "."

	return tox, toy
}

func moveRight(_map [][]string, x, y int) (int, int) {
	tox := x + 1

	if _map[y][tox] != "." {
		moveRight(_map, tox, y)
	}

	_map[y][tox] = _map[y][x]
	_map[y][x] = "."

	return tox, y
}

func moveLeft(_map [][]string, x, y int) (int, int) {
	tox := x - 1

	if _map[y][tox] != "." {
		moveLeft(_map, tox, y)
	}

	_map[y][tox] = _map[y][x]
	_map[y][x] = "."

	return tox, y
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

func canMoveUp(_map [][]string, x, y int) bool {
	tox, toy := x, y

	toy--

	if _map[toy][tox] == "[" {
		return canMoveUp(_map, tox, toy) && canMoveUp(_map, tox+1, toy)
	}

	if _map[toy][tox] == "]" {
		return canMoveUp(_map, tox, toy) && canMoveUp(_map, tox-1, toy)
	}

	if _map[toy][tox] == "#" {
		return false
	}

	return true
}

func canMoveDown(_map [][]string, x, y int) bool {
	tox, toy := x, y

	toy++

	if _map[toy][tox] == "[" {
		return canMoveDown(_map, tox, toy) && canMoveDown(_map, tox+1, toy)
	}

	if _map[toy][tox] == "]" {
		return canMoveDown(_map, tox, toy) && canMoveDown(_map, tox-1, toy)
	}

	if _map[toy][tox] == "#" {
		return false
	}

	return true
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
