package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 02 - xxxx
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
)

type Day02 struct {
	input     []int64
	solution1 int64
	solution2 int64
}

func (p *Day02) GetName() string {
	return "AoC 2021 - Day 2 - xx"
}

func (p *Day02) Init() {
	// Read input
	p.input = lib.ParseIntLines(lib.ReadInputLines("input/day02-test.txt"))
	// p.input = lib.ParseIntLines(lib.ReadInputLines("input/day02-input.txt"))

}

func (p *Day02) Run1() {
	p.solution1 = 0
}

func (p *Day02) Run2() {
	p.solution2 = 0
}

func (p *Day02) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day02) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
