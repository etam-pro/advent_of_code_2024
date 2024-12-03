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
	matcher := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches := matcher.FindAllString(input, -1)

	total := 0
	for _, match := range matches {
		nums := utils.ParseLineInts(match)
		total += nums[0] * nums[1]
	}

	fmt.Println(total)
}
