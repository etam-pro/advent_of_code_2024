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
	filept := nextFile(disk, len(disk)-1)

	for filept >= 0 {
		filesize := getFileSize(disk, filept)
		spacept := nextSpace(disk, filesize)

		if spacept >= 0 && spacept < filept {
			for i := 0; i < filesize; i++ {
				disk[spacept+i] = disk[filept-i]
				disk[filept-i] = SPACE
			}
		}

		filept = nextFile(disk, filept-filesize)
	}

	return disk
}

func nextSpace(disk []int, size int) int {
	for i := 0; i < len(disk); i++ {
		if disk[i] == SPACE {
			fit := true
			for j := 0; j < size; j++ {
				if i+j >= len(disk) || disk[i+j] != SPACE {
					fit = false
					break
				}
			}

			if fit {
				return i
			}
		}
	}

	return -1
}

func nextFile(disk []int, pt int) int {
	for i := pt; i >= 0; i-- {
		if disk[i] != SPACE {
			return i
		}
	}

	return -1
}

func getFileSize(disk []int, pt int) int {
	size := 1

	for i := pt - 1; i >= 0; i-- {
		if disk[i] == disk[pt] {
			size++
		} else {
			return size
		}

	}

	return size
}

func calc(disk []int) int {
	total := 0

	for i, c := range disk {
		if c == SPACE {
			continue
		}
		total += i * c
	}

	return total
}
