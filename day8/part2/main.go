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

// quick calculation of antinodes by applying the same x/y detas between 2 nodes on both directions
// ..........
// ...#......  <-- antinode 1
// ..........
// ....a.....
// ..........
// .....a....
// ..........
// ......#...  <-- antinode 2
// ..........
//
// repeat the same process on both ends until out of bounds
func markAntinodes(antennas map[string][]Coordinates, _map [][]string) {
	for _, nodes := range antennas {
		for i, node := range nodes {
			for j := i + 1; j < len(nodes); j++ {
				node2 := nodes[j]

				dx := node2.X - node.X
				dy := node2.Y - node.Y

				// antinode 1 and onward
				var nextx, nexty int = node.X - dx, node.Y - dy
				for inBound(_map, nextx, nexty) {
					_map[nexty][nextx] = "#"
					nextx -= dx
					nexty -= dy
				}

				// antiode 2 and onward
				nextx, nexty = node2.X+dx, node2.Y+dy
				for inBound(_map, nextx, nexty) {
					_map[nexty][nextx] = "#"
					nextx += dx
					nexty += dy
				}
			}

			// Mark the node as antinode too if it has more than one connection
			if len(nodes) > 1 {
				_map[node.Y][node.X] = "#"
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
