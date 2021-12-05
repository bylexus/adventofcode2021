package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 05 - Hydrothermal Venture
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
)

type Day05Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Day05 struct {
	input     []Day05Line
	solution1 int64
	solution2 int64

	coordmap map[int]int
}

func (p *Day05) GetName() string {
	return "AoC 2021 - Day 5 - Hydrothermal Venture"
}

func (p *Day05) Init() {
	// Read input
	// lines := lib.ParseGroupMatch(lib.ReadInputLines("input/day05-test.txt"), `^([0-9]+),([0-9]+)\s+->\s+([0-9]+),([0-9]+)`)
	lines := lib.ParseGroupMatch(lib.ReadInputLines("input/day05-input.txt"), `^([0-9]+),([0-9]+)\s+->\s+([0-9]+),([0-9]+)`)
	p.input = make([]Day05Line, len(lines))
	for i, line := range lines {
		p.input[i] = Day05Line{
			x1: lib.ToInt(line[1]),
			y1: lib.ToInt(line[2]),
			x2: lib.ToInt(line[3]),
			y2: lib.ToInt(line[4]),
		}
	}
}

func (p *Day05) Run1() {
	// run only horizontal/vertial lines:
	var overlapCount int64 = 0
	p.coordmap = make(map[int]int)
	for _, line := range p.input {
		if line.x1 != line.x2 && line.y1 != line.y2 {
			// skip if no h/v line:
			continue
		}
		xInc := 0
		yInc := 0
		// check how we have to count for each direction:
		if line.x1 < line.x2 {
			xInc = 1
		} else if line.x1 > line.x2 {
			xInc = -1
		}
		if line.y1 < line.y2 {
			yInc = 1
		} else if line.y1 > line.y2 {
			yInc = -1
		}
		x := line.x1
		y := line.y1
		for {
			key := x*10000 + y
			count, present := p.coordmap[key]
			// as soon as we have an overlap already, count it:
			if count == 1 {
				overlapCount++
			}
			if present == true {
				p.coordmap[key]++
			} else {
				p.coordmap[key] = 1
			}
			if x == line.x2 && y == line.y2 {
				break
			}
			x += xInc
			y += yInc
		}
	}
	p.solution1 = overlapCount
}

func (p *Day05) Run2() {
	// run all lines
	var overlapCount int64 = 0
	p.coordmap = make(map[int]int)
	for _, line := range p.input {
		xInc := 0
		yInc := 0
		// check how we have to count for each direction:
		if line.x1 < line.x2 {
			xInc = 1
		} else if line.x1 > line.x2 {
			xInc = -1
		}
		if line.y1 < line.y2 {
			yInc = 1
		} else if line.y1 > line.y2 {
			yInc = -1
		}
		x := line.x1
		y := line.y1
		for {
			key := x*10000 + y
			count, present := p.coordmap[key]
			// as soon as we have an overlap already, count it:
			if count == 1 {
				overlapCount++
			}
			if present == true {
				p.coordmap[key]++
			} else {
				p.coordmap[key] = 1
			}
			if x == line.x2 && y == line.y2 {
				break
			}
			x += xInc
			y += yInc
		}
	}
	p.solution2 = overlapCount
}

func (p *Day05) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day05) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
