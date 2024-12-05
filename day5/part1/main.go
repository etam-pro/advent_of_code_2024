package main

import (
	"fmt"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	rules := map[int][]int{}
	updates := [][]int{}

	isParsingRules := true

	utils.ReadLines("../data", func(line string) {
		if line == "" {
			isParsingRules = false
			return
		}

		input := utils.ParseLineInts(line)

		if isParsingRules {
			rules[input[0]] = append(rules[input[0]], input[1])
		} else {
			updates = append(updates, input)
		}
	})

	total := 0
	for _, update := range updates {
		if isValid(update, rules) {
			index := getMiddle(update)
			total += update[index]
		}
	}

	fmt.Println(total)
}

func isValid(update []int, rules map[int][]int) bool {
	// 1. Going through every page in the update
	for i, page := range update {
		// 2. check the pages before the current page to see if they are in the rules
		for _, before := range update[:i] {
			// 3. If they are, it is a violation and hence invalid
			if _, ok := rules[page]; ok {
				if utils.Contain(rules[page], before) {
					return false
				}
			}
		}
	}

	return true
}

func getMiddle(update []int) int {
	return len(update) / 2
}
