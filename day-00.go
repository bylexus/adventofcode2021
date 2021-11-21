package main

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 0 - setup tool chain
// This just implements the last year's day 1 riddle, to set up all the
// needed tools.
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
)

func solution1(input []int64) int64 {

	for o, line := range input {
		for i := o + 1; i < len(input); i++ {
			if line+input[i] == 2020 {
				return line * input[i]
			}
		}
	}
	return 0
}

func solution2(input []int64) int64 {
	for o, line := range input {
		for i := o + 1; i < len(input); i++ {
			for j := i + 1; j < len(input); j++ {
				if line+input[i]+input[j] == 2020 {
					return line * input[i] * input[j]
				}
			}
		}
	}
	return 0
}

func main() {
	lib.OutputTitle("Day 00 - Day 1 from AoC 2020 for starters")

	lines := lib.ParseIntLines(lib.ReadInputLines("input/day-00.txt"))

	var res1 int64
	duration := lib.MeasureTime(func() {
		res1 = solution1(lines)
	})
	lib.OutputSolution(1, duration, res1)

	var res2 int64
	duration = lib.MeasureTime(func() {
		res2 = solution2(lines)
	})
	lib.OutputSolution(2, duration, res2)

	input := [2]string{"Hello => this, is, a complicated, string", "and, another, much more simple"}
	res := lib.SplitLinesByRegex(input[:], `\s*,\s*`)
	for _, l := range res {
		fmt.Printf("Regex test: 0: %v, 1: %v\n", l[0], l[1])
	}
}
