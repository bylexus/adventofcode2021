package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 02 - Dive!
// ----------------------------------------------------------------------------

import (
	"fmt"
	"strconv"

	"alexi.ch/aoc2021/lib"
)

type Day02Input struct {
	Dir    string
	Amount int
}

type Day02 struct {
	input     []Day02Input
	solution1 int64
	solution2 int64
}

func (p *Day02) GetName() string {
	return "Dive!"
}

func (p *Day02) Init() {
	// Read input
	// lines := lib.ParseGroupMatch(lib.ReadInputLines("input/day02-test.txt"), `^([a-zA-Z]+)\s+([0-9]+)`)
	lines := lib.ParseGroupMatch(lib.ReadInputLines("input/day02-input.txt"), `^([a-zA-Z]+)\s+([0-9]+)`)
	p.input = make([]Day02Input, len(lines))
	for i, line := range lines {
		amount, err := strconv.Atoi(line[2])
		if err != nil {
			panic("Could not convert input")
		}
		p.input[i] = Day02Input{Dir: line[1], Amount: amount}
	}

}

func (p *Day02) Run1() {
	var x int64 = 0
	var depth int64 = 0
	for _, data := range p.input {
		if data.Dir == "forward" {
			x += int64(data.Amount)
		}
		if data.Dir == "up" {
			depth -= int64(data.Amount)
		}
		if data.Dir == "down" {
			depth += int64(data.Amount)
		}
	}
	p.solution1 = x * depth
}

func (p *Day02) Run2() {
	var x int64 = 0
	var depth int64 = 0
	var aim int64 = 0
	for _, data := range p.input {
		if data.Dir == "forward" {
			x += int64(data.Amount)
			depth += aim * int64(data.Amount)
		}
		if data.Dir == "up" {
			aim -= int64(data.Amount)
		}
		if data.Dir == "down" {
			aim += int64(data.Amount)
		}
	}
	p.solution2 = x * depth
}

func (p *Day02) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day02) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
