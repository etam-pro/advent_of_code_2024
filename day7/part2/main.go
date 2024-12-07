package main

import (
	"fmt"
	"strconv"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

type Eval struct {
	Target  int
	Nums    []int
	IsValid bool
}

func (e *Eval) Perform(acc int, op int, nums []int) int {
	// Early finish if the result is already found
	if e.IsValid {
		return acc
	}

	if len(nums) == 0 {
		// check if the final result is the target
		if acc == e.Target {
			e.IsValid = true
		}
		return acc
	}

	acc = calc(acc, op, nums[0])

	// Early exit if the evaluation is not finished but accumulated value is already greater than the target
	if acc > e.Target {
		return acc
	}

	// check all next operations
	// +
	total1 := e.Perform(acc, 0, nums[1:])
	// *
	total2 := e.Perform(acc, 1, nums[1:])
	// ||
	total3 := e.Perform(acc, 2, nums[1:])

	if total1 == e.Target || total2 == e.Target || total3 == e.Target {
		return e.Target
	}

	return acc
}

func main() {
	calibrations := [][]int{}

	utils.ReadLines("../data", func(line string) {
		calibrations = append(calibrations, utils.ParseLineInts(line))
	})

	total := 0
	for _, calibration := range calibrations {
		val := calibration[0]
		nums := calibration[1:]

		eval := &Eval{Target: val, Nums: nums}
		// check all possible starting operations
		eval.Perform(nums[0], 0, nums[1:])
		eval.Perform(nums[0], 1, nums[1:])
		eval.Perform(nums[0], 2, nums[1:])

		if eval.IsValid {
			total += val
		}
	}

	fmt.Println(total)
}

func calc(a, op, b int) int {
	switch op {
	case 0:
		return a + b
	case 1:
		return a * b
	case 2:
		num, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))

		if err != nil {
			panic(err)
		}

		return num
	}

	return 0
}
