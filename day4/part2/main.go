package main

import (
	"fmt"
	"strings"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	target := "MAS"
	puzzle := [][]string{}
	utils.ReadLines("../sample", func(line string) {
		puzzle = append(puzzle, strings.Split(line, ""))
	})

	total := 0
	for i, line := range puzzle {
		for j := range line {
			if isX(puzzle, i, j, target) {
				total += 1
			}
		}
	}

	fmt.Println(total)
}

func isX(puzzle [][]string, i, j int, target string) bool {
	if isOutOfBound(puzzle, i, j) {
		return false
	}

	if puzzle[i][j] != "A" {
		return false
	}

	return scanLeftToRight(puzzle, i, j, target) && scanRightToLeft(puzzle, i, j, target)
}

func scanLeftToRight(puzzle [][]string, i, j int, target string) bool {
	token := []string{puzzle[i-1][j-1], puzzle[i][j], puzzle[i+1][j+1]}
	reversed := reverse(token)

	return strings.Join(token, "") == target || strings.Join(reversed, "") == target
}

func scanRightToLeft(puzzle [][]string, i, j int, target string) bool {
	token := []string{puzzle[i-1][j+1], puzzle[i][j], puzzle[i+1][j-1]}
	reversed := reverse(token)

	return strings.Join(token, "") == target || strings.Join(reversed, "") == target
}

func isOutOfBound(puzzle [][]string, i, j int) bool {
	if j-1 < 0 {
		return true
	}

	if j+1 >= len(puzzle[i]) {
		return true
	}

	if i-1 < 0 {
		return true
	}

	if i+1 >= len(puzzle) {
		return true
	}

	return false
}

func reverse(chars []string) []string {
	revsered := make([]string, len(chars))
	for i, char := range chars {
		revsered[len(chars)-1-i] = char
	}

	return revsered
}
