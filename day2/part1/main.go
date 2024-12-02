package main

import (
	"fmt"
	"math"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	reports := [][]int{}
	utils.ReadLines("./data", func(line string) {
		reports = append(reports, utils.ParseLineInts(line))
	})

	total := 0
	for _, report := range reports {
		if isSafe(report) {
			total += 1
		}
	}

	fmt.Println(total)
}

func isSafe(report []int) bool {
	changes := map[int]int{}

	for i, level := range report {
		if i == len(report)-1 {
			continue
		}

		change := level - report[i+1]

		if change > 0 {
			changes[1] += 1
		}

		if change < 0 {
			changes[-1] += 1
		}

		if change == 0 || (changes[1] > 0 && changes[-1] > 0) || math.Abs(float64(change)) > 3 {
			return false
		}
	}

	return true
}
