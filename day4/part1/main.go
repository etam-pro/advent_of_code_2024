package main

import (
	"fmt"
	"strings"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	target := "XMAS"
	puzzle := [][]string{}
	utils.ReadLines("../data", func(line string) {
		puzzle = append(puzzle, strings.Split(line, ""))
	})

	total := 0
	for i, line := range puzzle {
		for j := range line {
			total += scan(puzzle, i, j, "up", target)
			total += scan(puzzle, i, j, "down", target)
			total += scan(puzzle, i, j, "left", target)
			total += scan(puzzle, i, j, "right", target)
			total += scan(puzzle, i, j, "up-left", target)
			total += scan(puzzle, i, j, "up-right", target)
			total += scan(puzzle, i, j, "down-left", target)
			total += scan(puzzle, i, j, "down-right", target)
		}
	}

	fmt.Println(total)
}

func scan(puzzle [][]string, i, j int, dir string, target string) int {
	var token string
	var err error

	switch dir {
	case "up":
		token, err = scanUp(puzzle, i, j)
	case "down":
		token, err = scanDown(puzzle, i, j)
	case "left":
		token, err = scanLeft(puzzle, i, j)
	case "right":
		token, err = scanRight(puzzle, i, j)
	case "up-left":
		token, err = scanUpLeft(puzzle, i, j)
	case "up-right":
		token, err = scanUpRight(puzzle, i, j)
	case "down-left":
		token, err = scanDownLeft(puzzle, i, j)
	case "down-right":
		token, err = scanDownRight(puzzle, i, j)
	}

	if err != nil {
		return 0
	}

	if token != target {
		return 0
	}

	fmt.Printf("Found %s at (%d, %d) in %s direction\n", target, i, j, dir)

	return 1
}

func scanRight(puzzle [][]string, i, j int) (string, error) {
	if j+3 >= len(puzzle[i]) {
		return "", fmt.Errorf("out of bounds")
	}

	token := strings.Join(puzzle[i][j:j+4], "")
	return token, nil
}

func scanLeft(puzzle [][]string, i, j int) (string, error) {
	if j-3 < 0 {
		return "", fmt.Errorf("out of bounds")
	}

	token := (reverse(puzzle[i][j-3 : j+1]))
	return strings.Join(token, ""), nil
}

func scanUp(puzzle [][]string, i, j int) (string, error) {
	if i-3 < 0 {
		return "", fmt.Errorf("out of bounds")
	}

	token := []string{puzzle[i][j], puzzle[i-1][j], puzzle[i-2][j], puzzle[i-3][j]}

	return strings.Join(token, ""), nil
}

func scanDown(puzzle [][]string, i, j int) (string, error) {
	if i+3 >= len(puzzle) {
		return "", fmt.Errorf("out of bounds")
	}

	token := []string{puzzle[i][j], puzzle[i+1][j], puzzle[i+2][j], puzzle[i+3][j]}

	return strings.Join(token, ""), nil
}

func scanUpRight(puzzle [][]string, i, j int) (string, error) {
	if i-3 < 0 || j+3 >= len(puzzle[i]) {
		return "", fmt.Errorf("out of bounds")
	}

	token := []string{puzzle[i][j], puzzle[i-1][j+1], puzzle[i-2][j+2], puzzle[i-3][j+3]}

	return strings.Join(token, ""), nil
}

func scanUpLeft(puzzle [][]string, i, j int) (string, error) {
	if i-3 < 0 || j-3 < 0 {
		return "", fmt.Errorf("out of bounds")
	}

	token := []string{puzzle[i][j], puzzle[i-1][j-1], puzzle[i-2][j-2], puzzle[i-3][j-3]}

	return strings.Join(token, ""), nil
}

func scanDownRight(puzzle [][]string, i, j int) (string, error) {
	if i+3 >= len(puzzle) || j+3 >= len(puzzle[i]) {
		return "", fmt.Errorf("out of bounds")
	}

	token := []string{puzzle[i][j], puzzle[i+1][j+1], puzzle[i+2][j+2], puzzle[i+3][j+3]}

	return strings.Join(token, ""), nil
}

func scanDownLeft(puzzle [][]string, i, j int) (string, error) {
	if i+3 >= len(puzzle) || j-3 < 0 {
		return "", fmt.Errorf("out of bounds")
	}

	token := []string{puzzle[i][j], puzzle[i+1][j-1], puzzle[i+2][j-2], puzzle[i+3][j-3]}

	return strings.Join(token, ""), nil
}

func reverse(chars []string) []string {
	revsered := make([]string, len(chars))
	for i, char := range chars {
		revsered[len(chars)-1-i] = char
	}

	return revsered
}
