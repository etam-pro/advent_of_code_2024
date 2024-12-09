package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SPACE = -1

func main() {
	diskmap := getDiskMap("../data")

	var disk []int = []int{}

	isSpace := false
	fileID := 0
	for _, c := range diskmap {
		for i := 0; i < c; i++ {
			if isSpace {
				disk = append(disk, SPACE)
			} else {
				disk = append(disk, fileID)
			}
		}

		if !isSpace {
			fileID++
		}

		isSpace = !isSpace
	}

	disk = compact(disk)
	total := calc(disk)

	fmt.Println(total)
}

func getDiskMap(filePath string) []int {
	bytes, _ := os.ReadFile(filePath)
	diskmapStr := strings.Split(string(bytes), "")

	diskmap := []int{}

	for _, c := range diskmapStr {
		num, _ := strconv.Atoi(c)
		diskmap = append(diskmap, num)
	}

	return diskmap
}

func compact(disk []int) []int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == SPACE {
			disk = disk[:len(disk)-1]
			continue
		}

		si := nextSpace(disk)

		if si == -1 {
			break
		}

		disk[si] = disk[i]
		disk[i] = SPACE

		disk = disk[:len(disk)-1]
	}

	return disk
}

func nextSpace(disk []int) int {
	for i, c := range disk {
		if c == SPACE {
			return i
		}
	}

	return -1
}

func calc(disk []int) int {
	total := 0

	for i, c := range disk {
		total += i * c
	}

	return total
}
