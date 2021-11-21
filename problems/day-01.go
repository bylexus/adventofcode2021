package problems

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

type Day01 struct {
	input     []int64
	solution1 int64
	solution2 int64
}

func (p *Day01) GetName() string {
	return "AoC 2021 - Day 1"
}

func (p *Day01) Init() {
	// Read input
	p.input = lib.ParseIntLines(lib.ReadInputLines("input/day-01.txt"))

}

func (p *Day01) Run1() {
	p.solution1 = 0
}

func (p *Day01) Run2() {
	p.solution2 = 0
}

func (p *Day01) GetSolution1() string {
	return fmt.Sprintf("%v", p.solution1)
}

func (p *Day01) GetSolution2() string {
	return fmt.Sprintf("%v", p.solution2)
}
