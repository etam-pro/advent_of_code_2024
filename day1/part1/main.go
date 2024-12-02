package main

import (
	"fmt"
	"math"
	"sort"

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

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	dist := 0.0
	for i, l := range left {
		dist += math.Abs(float64(right[i] - l))
	}

	fmt.Println(int64(dist))
}
