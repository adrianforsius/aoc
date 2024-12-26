package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Computer struct {
	A int64
	B int64
	C int64

	Program []int64
	Pointer int
	O       []int64
}

func (c *Computer) Quine() int64 {
	return 0
}

func (c *Computer) Run() {
	for c.Pointer < len(c.Program)-1 {

		opCode := c.Program[c.Pointer]
		operand := c.Program[c.Pointer+1]

		// fmt.Print64ln("opp", opCode, "operand", operand, "out", c.O)
		fmt.Printf("%#v\n", c)
		// if len(c.Program) > 2 {
		// 	o, rand, rest := c.Program[0], c.Program[0], c.Program[2:]
		// 	c.Program = rest
		// 	operand = rand
		// 	opCode = o
		// } else {
		// 	opCode = c.Program[0]
		// 	operand = c.Program[1]
		// 	c.Program = []int64{}
		// }

		switch opCode {
		case 0:
			c.Adv(operand)
			c.Pointer += 2
		case 1:
			c.Blx(operand)
			c.Pointer += 2
		case 2:
			c.Bst(operand)
			c.Pointer += 2
		case 3:
			if c.A == 0 {
				c.Pointer += 2
				continue
			}
			c.Pointer = int(operand)
			// c.Jnz(operand)
		case 4:
			c.Bxc(operand)
			c.Pointer += 2
		case 5:
			c.Out(operand)
			c.Pointer += 2
		case 6:
			c.Bdv(operand)
			c.Pointer += 2
		case 7:
			c.Cdv(operand)
			c.Pointer += 2
		}
	}
}

func (c *Computer) ToRegister(operand int64) int64 {
	switch operand {
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	}
	return operand
}

func pow(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func (c *Computer) Adv(operand int64) {
	// fmt.Println("adv combo", c.ToRegister(operand), "c.A", c.A)
	c.A >>= c.ToRegister(operand)
}

func (c *Computer) Blx(operand int64) {
	c.B = c.B ^ operand
}

func (c *Computer) Bst(operand int64) {
	c.B = c.ToRegister(operand) % 8
}

func (c *Computer) Jnz(operand int64) {
}

func (c *Computer) Bxc(operand int64) {
	c.B = c.B ^ c.C
}

func (c *Computer) Out(operand int64) {
	// fmt.Print64ln("out", c.ToRegister(operand)%8)
	c.O = append(c.O, c.ToRegister(operand)%8)
}

func (c *Computer) Bdv(operand int64) {
	// c.B = (c.A / pow(2, c.ToRegister(operand)))
	c.B = c.A >> c.ToRegister(operand)
}

func (c *Computer) Cdv(operand int64) {
	// c.C = (c.A / pow(2, c.ToRegister(operand)))
	c.C = c.A >> c.ToRegister(operand)
}

func Parse(in string) *Computer {
	parts := strings.Split(in, "\n\n")
	regLine := strings.Split(parts[0], "\n")
	aStr := strings.Split(regLine[0], "Register A: ")
	a, _ := strconv.ParseInt(aStr[1], 10, 64)
	bStr := strings.Split(regLine[1], "Register B: ")
	b, _ := strconv.ParseInt(bStr[1], 10, 64)
	cStr := strings.Split(regLine[2], "Register C: ")
	c, _ := strconv.ParseInt(cStr[1], 10, 64)

	var prog []int64
	p := strings.Split(parts[1], "Program: ")
	for _, s := range strings.Split(p[1], ",") {
		i, _ := strconv.ParseInt(s, 10, 64)
		prog = append(prog, i)
	}

	return &Computer{a, b, c, prog, 0, []int64{}}
}

func Day2(computer *Computer) int64 {
	// computer.Run()
	return computer.Quine()
}
