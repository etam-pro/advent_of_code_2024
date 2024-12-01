package utils

import (
	"bufio"
	"os"
)

func ReadLines(path string, action func(line string)) {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()

		action(str)
	}
}
