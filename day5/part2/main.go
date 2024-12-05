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

func fix(update []int, rules map[int][]int) []int {
	newUpdate := make([]int, len(update))
	copy(newUpdate, update)

	// Keep swapping until it is valid
	for !isValid(newUpdate, rules) {
		// go through all the rules
		for page, pagerules := range rules {
			// If the page is not in the update, skip
			if utils.IndexOf(newUpdate, page) == -1 {
				continue
			}

			for _, rule := range pagerules {
				// If the rule is not in the update, skip
				if utils.IndexOf(newUpdate, rule) == -1 {
					continue
				}

				// If it is a violation, swap the pages
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
