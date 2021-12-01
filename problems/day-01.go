package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 01 - Sonar Sweep
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
)

type Day01 struct {
	input     []int64
	solution1 int64
	solution2 int64
}

func (p *Day01) GetName() string {
	return "AoC 2021 - Day 1 - Sonar Sweep"
}

func (p *Day01) Init() {
	// Read input
	// p.input = lib.ParseIntLines(lib.ReadInputLines("input/day01-test.txt"))
	p.input = lib.ParseIntLines(lib.ReadInputLines("input/day01-input.txt"))

}

func (p *Day01) Run1() {
	var prev int64 = p.input[0]
	var counter int64 = 0
	for i := 1; i < len(p.input); i++ {
		if p.input[i] > prev {
			counter++
		}
		prev = p.input[i]
	}
	p.solution1 = counter
}

func (p *Day01) Run2() {
	var prev int64 = p.input[0] + p.input[1] + p.input[2]
	var counter int64 = 0
	for i := 1; i < len(p.input)-2; i++ {
		win_sum := p.input[i] + p.input[i+1] + p.input[i+2]
		if win_sum > prev {
			counter++
		}
		prev = win_sum
	}
	p.solution2 = counter
}

func (p *Day01) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day01) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
