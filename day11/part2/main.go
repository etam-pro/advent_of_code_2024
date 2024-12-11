package main

import (
	"fmt"
	"strconv"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

type Stone struct {
	Num        int
	Split      int
	TargetIter int
}

func main() {
	arrangement := []int{}
	stones := []Stone{}

	utils.ReadLines("../data", func(line string) {
		arrangement = utils.ParseLineInts(line)
		for _, num := range arrangement {
			stones = append(stones, Stone{Num: num, TargetIter: 75})
		}
	})

	total := 0
	memo := map[[3]int]int{}
	for _, stone := range stones {
		total += stone.Solve(1, 0, stone.Num, memo)
	}

	fmt.Println(total)
}

func (s *Stone) Solve(acc int, iter int, num int, memo map[[3]int]int) int {
	if iter == s.TargetIter {
		return acc
	}

	if val, ok := memo[[3]int{acc, iter, num}]; ok {
		return val
	}

	if num == 0 {
		return s.Solve(acc, iter+1, 1, memo)
	}

	if isEvenDigit(num) {
		lr := split(num)
		// Splitting, left (original stone) new acc acount as 0 because it is the original branch
		//            right (new stone) new acc account as 1 because it is the new branch
		newAcc := acc + s.Solve(0, iter+1, lr[0], memo) + s.Solve(1, iter+1, lr[1], memo)
		memo[[3]int{acc, iter, num}] = newAcc
		return newAcc
	}

	return s.Solve(acc, iter+1, num*2024, memo)
}

func isEvenDigit(num int) bool {
	return len(strconv.Itoa(num))%2 == 0
}

func split(num int) []int {
	str := strconv.Itoa(num)
	half := len(str) / 2
	left, _ := strconv.Atoi(str[:half])
	right, _ := strconv.Atoi(str[half:])
	return []int{left, right}
}
