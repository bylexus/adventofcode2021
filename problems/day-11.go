package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 11 - Dumbo Octopus
// ----------------------------------------------------------------------------

import (
	"fmt"
	"strconv"

	"alexi.ch/aoc2021/lib"
	"alexi.ch/aoc2021/lib/types"
)

type Day11Octopuses map[types.Point]int

type Day11 struct {
	solution1  uint64
	solution2  uint64
	octopuses  Day11Octopuses
	origPusses Day11Octopuses
	dirs       []types.Point
	flashes    uint64
}

func (p *Day11) GetName() string {
	return "Dumbo Octopus"
}

func (p *Day11) Init() {
	// Read input
	// lines := lib.ReadInputLines("input/day11-test-small.txt")
	// lines := lib.ReadInputLines("input/day11-test.txt")
	lines := lib.ReadInputLines("input/day11-input.txt")
	p.dirs = make([]types.Point, 8)

	p.dirs[0] = types.Point{X: -1, Y: -1}
	p.dirs[1] = types.Point{X: 0, Y: -1}
	p.dirs[2] = types.Point{X: 1, Y: -1}
	p.dirs[3] = types.Point{X: -1, Y: 0}
	p.dirs[5] = types.Point{X: 1, Y: 0}
	p.dirs[4] = types.Point{X: -1, Y: 1}
	p.dirs[6] = types.Point{X: 0, Y: 1}
	p.dirs[7] = types.Point{X: 1, Y: 1}

	p.octopuses = make(Day11Octopuses)
	p.origPusses = make(Day11Octopuses)
	for i := range lines {
		line := lines[i]
		for idx, r := range line {
			nr, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			key := types.Point{X: idx, Y: i}
			p.octopuses[key] = nr
			p.origPusses[key] = nr
		}
	}
}

func (p *Day11) increase() {
	for key := range p.octopuses {
		p.octopuses[key]++
	}
}

func (p *Day11) flashSingle(key types.Point) {
	if p.octopuses[key] == 0 {
		// already in the "flashed" state, do not flash again
		return
	}
	p.flashes++
	p.octopuses[key] = 0 // mark as "flashed"
	// flash all 8 adjacent octies:
	point := key
	for _, d := range p.dirs {
		dx := point.X + d.X
		dy := point.Y + d.Y
		dKey := types.Point{X: dx, Y: dy}
		val, present := p.octopuses[dKey]
		if present == true && val > 0 {
			p.octopuses[dKey]++
		}
	}
}

func (p *Day11) flash() {
	for {
		runAgain := false
		for key := range p.octopuses {
			if p.octopuses[key] > 9 {
				runAgain = true
				p.flashSingle(key)
			}
		}
		if runAgain == false {
			break
		}
	}
}

func (p *Day11) checkAllFlashed() bool {
	for key := range p.octopuses {
		if p.octopuses[key] != 0 {
			return false
		}
	}

	return true
}

func (p *Day11) Run1() {
	runcounter := 0
	for {
		runcounter++
		p.increase()
		p.flash()

		if runcounter >= 100 {
			break
		}
	}
	p.solution1 = p.flashes
}

func (p *Day11) Run2() {
	// reset octopusses:
	p.octopuses = p.origPusses
	runcounter := 0
	p.flashes = 0
	for {
		runcounter++
		p.increase()
		p.flash()
		if p.checkAllFlashed() {
			break
		}
	}
	p.solution2 = uint64(runcounter)
}

func (p *Day11) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day11) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
