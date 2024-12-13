package main

import (
	"fmt"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

func main() {
	machines := []Machine{}

	cur := NewMachine()
	inputType := 1
	utils.ReadLines("../data", func(line string) {
		if line == "" {
			machines = append(machines, *cur)
			cur = NewMachine()
			inputType = 1
			return
		}

		input := utils.ParseLineInts(line)

		if inputType == 1 {
			cur.A = Button{X: input[0], Y: input[1]}
		}

		if inputType == 2 {
			cur.B = Button{X: input[0], Y: input[1]}
		}

		if inputType == 3 {
			cur.Prize = input
		}

		inputType++
	})

	total := 0
	for _, m := range machines {
		total += m.Calc()
	}

	fmt.Println(total)
}

type Button struct {
	X int
	Y int
}

type Machine struct {
	A         Button
	B         Button
	Prize     []int
	MinTokens int
}

func NewMachine() *Machine {
	return &Machine{MinTokens: -1}
}

func (m *Machine) Calc() int {
	total := 0

	var numbA, numbB int
	solved := false

	// optimize by A
	numbA, numbB, solved = solve(m.A.X, m.A.Y, m.B.X, m.B.Y, m.Prize[0], m.Prize[1])
	if solved {
		tokens := numbA*3 + numbB
		if total == 0 || tokens < total {
			total = tokens
		}
	}

	numbA, numbB, solved = solve(m.A.Y, m.A.X, m.B.Y, m.B.X, m.Prize[1], m.Prize[0])
	if solved {
		tokens := numbA*3 + numbB
		if total == 0 || total < total {
			total = tokens
		}
	}

	// optimize by B
	numbB, numbA, solved = solve(m.B.X, m.B.Y, m.A.X, m.A.Y, m.Prize[0], m.Prize[1])
	if solved {
		tokens := numbA*3 + numbB
		if total == 0 || tokens < total {
			total = tokens
		}
	}

	numbB, numbA, solved = solve(m.B.Y, m.B.X, m.A.Y, m.A.X, m.Prize[1], m.Prize[0])
	if solved {
		tokens := numbA*3 + numbB
		if total == 0 || tokens < total {
			total = tokens
		}
	}

	return total
}

func solve(ax, ay, bx, by, targetx, check int) (int, int, bool) {
	numA := targetx / ax
	numB := (targetx - ax*numA) / bx

	remaining := targetx - ax*numA - bx*numB

	for {
		if numA < 0 {
			return -1, -1, false
		}

		numA -= 1
		numB = (targetx - ax*numA) / bx
		remaining = targetx - ax*numA - bx*numB

		if remaining == 0 {
			if numA*ay+numB*by == check {
				break
			}
		}
	}

	if numA > 100 || numB > 100 {
		return -1, -1, false
	}

	return numA, numB, true
}
