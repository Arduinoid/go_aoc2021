package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayTwo() {
	fmt.Println("Day Two - Sub Controls")
	pos := positionCalc(string(readInput(2)))
	fmt.Printf("Part 1 final calculated position value is %d\n", pos)
	aPos := positionAdvCalc(string(readInput(2)))
	fmt.Printf("Part 2 final calculated position value is %d\n", aPos)
}

func positionCalc(input string) int {
	var h int
	var d int
	insts := parseControls(input)
	for _, inst := range insts {
		switch inst.move {
		case "forward":
			h = h + inst.amount
		case "up":
			d = d - inst.amount
		case "down":
			d = d + inst.amount
		}
	}
	return h * d
}

func positionAdvCalc(input string) int {
	var (
		h   int
		d   int
		aim int
	)
	insts := parseControls(input)
	for _, inst := range insts {
		switch inst.move {
		case "forward":
			h = h + inst.amount
			if aim != 0 {
				d = inst.amount*aim + d
			}
		case "up":
			aim = aim - inst.amount
		case "down":
			aim = aim + inst.amount
		}
	}
	return h * d
}

func parseControls(in string) []inst {
	var instructions []inst
	list := strings.Split(in, "\n")
	for i := range list {
		s := strings.Split(list[i], " ")
		num, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, inst{move: s[0], amount: num})
	}
	return instructions
}

type inst struct {
	move   string
	amount int
}
