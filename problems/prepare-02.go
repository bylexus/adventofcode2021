package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 0 - setup tool chain
// This just implements the last year's day 2 riddle, to set up all the
// needed tools.
// ----------------------------------------------------------------------------

import (
	"fmt"
	"strconv"
	"strings"

	"alexi.ch/aoc2021/lib"
)

type Data struct {
	min      int
	max      int
	letter   string
	password string
}

type Prepare02 struct {
	input     []Data
	solution1 int64
	solution2 int64
}

func (p *Prepare02) GetName() string {
	return "Preparations - AoC 2020 - Day 2"
}

func (p *Prepare02) Init() {
	// Read input
	lines := lib.ParseGroupMatch(lib.ReadInputLines("input/prepare-02.txt"), `^([0-9]+)-([0-9]+)\s+([a-zA-z]):\s+([a-zA-Z]+)`)
	p.input = make([]Data, 0)
	for _, data := range lines {
		min, err := strconv.Atoi(data[1])
		if err != nil {
			panic("Cannot parse min value: " + data[1])
		}
		max, err := strconv.Atoi(data[2])
		if err != nil {
			panic("Cannot parse max value: " + data[2])
		}
		p.input = append(p.input, Data{
			min:      min,
			max:      max,
			letter:   data[3],
			password: data[4],
		})
	}
}

func (p *Prepare02) Run1() {
	p.solution1 = 0
	for _, rule := range p.input {
		count := strings.Count(rule.password, rule.letter)
		if count >= rule.min && count <= rule.max {
			p.solution1++
		}
	}
}

func (p *Prepare02) Run2() {
	p.solution2 = 0
	for _, rule := range p.input {
		if rule.password[rule.min-1:rule.min] == rule.letter && rule.password[rule.max-1:rule.max] == rule.letter {
			continue
		}
		if rule.password[rule.min-1:rule.min] == rule.letter || rule.password[rule.max-1:rule.max] == rule.letter {
			p.solution2++
		}
	}
}

func (p *Prepare02) GetSolution1() string {
	return fmt.Sprintf("%v", p.solution1)
}

func (p *Prepare02) GetSolution2() string {
	return fmt.Sprintf("%v", p.solution2)
}
