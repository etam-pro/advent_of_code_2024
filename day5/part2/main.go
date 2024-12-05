package main

import (
	"fmt"
	"math"

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
		if !isValid(update, rules) {
			corrected := fix(update, rules)
			index := getMiddle(corrected)
			total += corrected[index]
		}
	}

	fmt.Println(total)
}

func isValid(update []int, rules map[int][]int) bool {
	for i, page := range update {
		for _, before := range update[:i] {
			if _, ok := rules[page]; ok {
				if contain(rules[page], before) {
					return false
				}
			}
		}
	}

	return true
}

func getMiddle(update []int) int {
	return int(math.Ceil(float64(len(update) / 2)))
}

func fix(update []int, rules map[int][]int) []int {
	newUpdate := make([]int, len(update))
	copy(newUpdate, update)

	for !isValid(newUpdate, rules) {
		for page, pagerules := range rules {
			if utils.IndexOf(newUpdate, page) == -1 {
				continue
			}

			for _, rule := range pagerules {
				if utils.IndexOf(newUpdate, rule) == -1 {
					continue
				}

				if utils.IndexOf(newUpdate, page) > utils.IndexOf(newUpdate, rule) {
					pi := utils.IndexOf(newUpdate, page)
					pri := utils.IndexOf(newUpdate, rule)
					newUpdate[pi] = rule
					newUpdate[pri] = page
				}
			}
		}
	}

	return newUpdate
}
