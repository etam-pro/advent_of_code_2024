package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	data, _ := os.ReadFile("../data")
	input := string(data)
	matcher := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	matches := matcher.FindAllString(input, -1)

	total := 0
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}
		if match == "don't()" {
			enabled = false
			continue
		}
		if enabled == false {
			continue
		}
		nums := utils.ParseLineInts(match)
		total += nums[0] * nums[1]
	}

	fmt.Println(total)
}
