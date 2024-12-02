package main

import (
	"fmt"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	left := []int{}
	right := []int{}

	utils.ReadLines("./data", func(line string) {
		if line == "" {
			return
		}

		ints := utils.ParseLineInts(line)
		left = append(left, ints[0])
		right = append(right, ints[1])
	})

	score := 0
	for _, num := range left {
		score += similarityScore(num, right)
	}

	fmt.Println(score)
}

func similarityScore(num int, list []int) int {
	apperances := count(num, list)
	return num * apperances
}

func count(num int, list []int) int {
	total := 0
	for _, n := range list {
		if num == n {
			total += 1
		}
	}

	return total
}
