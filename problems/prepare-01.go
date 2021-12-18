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

type Prepare01 struct {
	input     []int64
	solution1 int64
	solution2 int64
}

func (p *Prepare01) GetName() string {
	return "Preparations - AoC 2020 - Day 1"
}

func (p *Prepare01) Init() {
	// Read input
	p.input = lib.ParseIntLines(lib.ReadInputLines("input/prepare-01.txt"))

}

func (p *Prepare01) Run1() {
	p.solution1 = 0
	for o, line := range p.input {
		for i := o + 1; i < len(p.input); i++ {
			if line+p.input[i] == 2020 {
				p.solution1 = line * p.input[i]
				return
			}
		}
	}
}

func (p *Prepare01) Run2() {
	p.solution2 = 0
	for o, line := range p.input {
		for i := o + 1; i < len(p.input); i++ {
			for j := i + 1; j < len(p.input); j++ {
				if line+p.input[i]+p.input[j] == 2020 {
					p.solution2 = line * p.input[i] * p.input[j]
					return
				}
			}
		}
	}
}

func (p *Prepare01) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Prepare01) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
