package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func dayOne() {
	input := parseDayOne()
	count := countIncrease(input)
	winCount := countWindowedIncrease(input)
	fmt.Println("Day One - Depth Increase Measurements")
	fmt.Printf("Part 1\nthere were %d increases in depth\n", count)
	fmt.Printf("Part 2\nthere were %d increases in windowed depth\n", winCount)
}

func countIncrease(input []int) int {
	var count int
	var lastInput int

	for i, depth := range input {
		if i == 0 {
			lastInput = depth
			continue
		}

		if notLess(depth, lastInput) {
			count++
		}
		lastInput = depth
	}
	return count
}

func countWindowedIncrease(input []int) int {
	var count int
	var lastInput int
	var length int = len(input)
	var window int = 3

	for i := range input {
		if i+window > length {
			return count
		}
		current := input[i] + input[i+1] + input[i+2]
		if i == 0 {
			lastInput = current
			continue
		}

		if notLess(current, lastInput) {
			count++
		}
		lastInput = current
	}
	return count
}

func notLess(a, b int) bool {
	return a > b
}

func parseDayOne() []int {
	var result []int
	b := readInput(1)
	bs := bytes.Fields(b)
	for i := range bs {
		n, err := strconv.Atoi(string(bs[i]))
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}
	return result
}
