package main

import (
	"fmt"
	"strings"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

type Coordinates struct {
	X int
	Y int
}

func main() {
	_map := [][]string{}

	utils.ReadLines("../data", func(line string) {
		_map = append(_map, strings.Split(line, ""))
	})

	antennas := getAntennas(_map)
	markAntinodes(antennas, _map)

	total := calcAntinodes(_map)

	utils.PrintMap(_map)
	fmt.Println(total)
}

func getAntennas(_map [][]string) map[string][]Coordinates {
	antennas := map[string][]Coordinates{}

	for y, row := range _map {
		for x, cell := range row {
			if isAntenna(cell) {
				antennas[cell] = append(antennas[cell], Coordinates{x, y})
			}
		}
	}
	return antennas
}

func isAntenna(cell string) bool {
	return cell != "."
}

func markAntinodes(antennas map[string][]Coordinates, _map [][]string) {
	for _, nodes := range antennas {
		for i, node := range nodes {
			for j := i + 1; j < len(nodes); j++ {
				node2 := nodes[j]

				dx := node2.X - node.X
				dy := node2.Y - node.Y

				if inBound(_map, node.X-dx, node.Y-dy) {
					_map[node.Y-dy][node.X-dx] = "#"
				}

				if inBound(_map, node2.X+dx, node2.Y+dy) {
					_map[node2.Y+dy][node2.X+dx] = "#"
				}
			}
		}
	}
}

func inBound(_map [][]string, x, y int) bool {
	return x >= 0 && x < len(_map[0]) && y >= 0 && y < len(_map)
}

func calcAntinodes(_map [][]string) int {
	total := 0

	for _, row := range _map {
		for _, cell := range row {
			if cell == "#" {
				total++
			}
		}
	}

	return total
}
