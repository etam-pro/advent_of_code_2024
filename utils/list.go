package utils

import (
	"fmt"
	"strings"
)

func Contain(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func ContainString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func IndexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}

func PrintMap(_map [][]string) {
	for _, row := range _map {
		fmt.Println(strings.Join(row, ""))
	}
}
