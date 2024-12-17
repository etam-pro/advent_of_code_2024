package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/etam-pro/advent_of_code_2024/utils"
)

const (
	OPCODE_adv = 0
	OPCODE_bxl = 1
	OPCODE_bst = 2
	OPCODE_jnz = 3
	OPCODE_bxc = 4
	OPCODE_out = 5
	OPCODE_bdv = 6
	OPCODE_cdv = 7
)

type Computer struct {
	A       int
	B       int
	C       int
	Pointer int
	Program []int
	Output  []int
}

func (c *Computer) Run() {
	for c.Pointer < len(c.Program)-1 {
		op, operand := c.Program[c.Pointer], c.Program[c.Pointer+1]
		c.calc(op, operand)

		if op != OPCODE_jnz {
			c.Pointer += 2
		}
	}
}

func (c *Computer) calc(op, operand int) {
	switch op {
	case OPCODE_adv:
		c.PerfomAdv(operand)
	case OPCODE_bxl:
		c.PerfomBxl(operand)
	case OPCODE_bst:
		c.PerfomBst(operand)
	case OPCODE_jnz:
		c.PerfomJnz(operand)
	case OPCODE_bxc:
		c.PerfomBxc(operand)
	case OPCODE_out:
		c.PerfomOut(operand)
	case OPCODE_bdv:
		c.PerfomBdv(operand)
	case OPCODE_cdv:
		c.PerfomCdv(operand)
	}
}

func (c *Computer) PerfomAdv(operand int) {
	numerator := c.A
	dinominator := int(math.Pow(2, float64(c.ComboOperand(operand))))

	c.A = numerator / dinominator
}

func (c *Computer) PerfomBxl(operand int) {
	c.B = c.B ^ operand
}

func (c *Computer) PerfomBst(operand int) {
	c.B = c.ComboOperand(operand) % 8
}

func (c *Computer) PerfomJnz(operand int) {
	if c.A == 0 {
		c.Pointer += 2
		return
	}

	c.Pointer = operand
}

func (c *Computer) PerfomBxc(operand int) {
	c.B = c.B ^ c.C
}

func (c *Computer) PerfomOut(operand int) {
	c.Output = append(c.Output, c.ComboOperand(operand)%8)
}

func (c *Computer) PerfomBdv(operand int) {
	numerator := c.A
	dinominator := int(math.Pow(2, float64(c.ComboOperand(operand))))

	c.B = numerator / dinominator
}

func (c *Computer) PerfomCdv(operand int) {
	numerator := c.A
	dinominator := int(math.Pow(2, float64(c.ComboOperand(operand))))

	c.C = numerator / dinominator
}

func (c *Computer) ComboOperand(val int) int {
	switch val {
	case 0, 1, 2, 3:
		return val
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	case 7:
		panic("Attempt to access reserved operand")
	default:
		panic("Invalid operand")
	}
}

func (c *Computer) PrintOutput() {
	strs := []string{}
	for _, val := range c.Output {
		strs = append(strs, strconv.Itoa(val))
	}
	fmt.Println(strings.Join(strs, ","))
}

func main() {
	inputs := []int{}
	utils.ReadLines("../data", func(line string) {
		inputs = append(inputs, utils.ParseLineInts(line)...)
	})

	a, b, c, program := inputs[0], inputs[1], inputs[2], inputs[3:]

	computer := &Computer{A: a, B: b, C: c, Pointer: 0, Program: program}
	computer.Run()
	computer.PrintOutput()
}
