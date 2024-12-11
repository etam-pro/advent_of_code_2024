package main

import (
	"fmt"
	"strconv"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	arrangement := []int{}

	utils.ReadLines("../data", func(line string) {
		arrangement = utils.ParseLineInts(line)
	})

	for i := 0; i < 25; i++ {
		new := []int{}
		for _, num := range arrangement {
			if num == 0 {
				new = append(new, 1)
			} else if isEvenDigit(num) {
				new = append(new, split(num)...)
			} else {
				new = append(new, num*2024)
			}
		}

		arrangement = new
	}

	fmt.Println(len(arrangement))
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
