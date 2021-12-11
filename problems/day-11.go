package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 11 - xxx
// ----------------------------------------------------------------------------

import (
	"fmt"
	"strconv"

	"alexi.ch/aoc2021/lib"
)

type Day11Octopuses [][]int

type Day11 struct {
	solution1 uint64
	solution2 uint64
	octopuses Day11Octopuses
	dirs      []lib.Point
	flashes   uint64
}

func (p *Day11) GetName() string {
	return "AoC 2021 - Day 11 - xxx"
}

func (p *Day11) Init() {
	// Read input
	// lines := lib.ReadInputLines("input/day11-test-small.txt")
	// lines := lib.ReadInputLines("input/day11-test.txt")
	lines := lib.ReadInputLines("input/day11-input.txt")
	p.dirs = make([]lib.Point, 8)

	p.dirs[0] = lib.Point{X: -1, Y: -1}
	p.dirs[1] = lib.Point{X: 0, Y: -1}
	p.dirs[2] = lib.Point{X: 1, Y: -1}
	p.dirs[3] = lib.Point{X: -1, Y: 0}
	p.dirs[5] = lib.Point{X: 1, Y: 0}
	p.dirs[4] = lib.Point{X: -1, Y: 1}
	p.dirs[6] = lib.Point{X: 0, Y: 1}
	p.dirs[7] = lib.Point{X: 1, Y: 1}

	p.octopuses = make(Day11Octopuses, len(lines))
	for i := range lines {
		line := lines[i]
		p.octopuses[i] = make([]int, len(line))
		for idx, r := range line {
			nr, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			p.octopuses[i][idx] = nr
		}
	}

	fmt.Printf("octopuses: %#v\n", p.octopuses)
}

func (p *Day11) increase() {
	for y := 0; y < len(p.octopuses); y++ {
		for x := 0; x < len(p.octopuses[y]); x++ {
			p.octopuses[y][x] += 1
		}
	}
}

func (p *Day11) flashSingle(x, y int) {
	o := p.octopuses[y][x]
	if o == 0 {
		// already in the "flashed" state, do not flash again
		return
	}
	p.flashes++
	p.octopuses[y][x] = 0 // mark as "flashed"
	// flash all 8 adjacent octies:
	for _, d := range p.dirs {
		dx := x + d.X
		dy := y + d.Y
		if dx >= 0 && dy >= 0 && dy < len(p.octopuses) && dx < len(p.octopuses[dy]) && p.octopuses[dy][dx] > 0 {
			p.octopuses[dy][dx]++
		}
	}
}

func (p *Day11) flash() {
	for {
		runAgain := false
		for y := 0; y < len(p.octopuses); y++ {
			for x := 0; x < len(p.octopuses[y]); x++ {
				if p.octopuses[y][x] > 9 {
					runAgain = true
					p.flashSingle(x, y)
				}
			}
		}
		if runAgain == false {
			break
		}
	}
}

func (p *Day11) checkAllFlashed() bool {
	for y := 0; y < len(p.octopuses); y++ {
		for x := 0; x < len(p.octopuses[y]); x++ {
			if p.octopuses[y][x] != 0 {
				return false
			}
		}
	}

	return true
}

func (p *Day11) print() {
	for y := 0; y < len(p.octopuses); y++ {
		for x := 0; x < len(p.octopuses[y]); x++ {
			fmt.Printf("%v ", p.octopuses[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (p *Day11) Run1() {
	runcounter := 0
	p.print()
	for {
		runcounter++
		p.increase()
		p.flash()
		// p.print()
		// fmt.Println("============================")

		if runcounter >= 100 {
			break
		}
	}
	p.solution1 = p.flashes
}

func (p *Day11) Run2() {
	p.Init()
	runcounter := 0
	p.flashes = 0
	p.print()
	for {
		runcounter++
		p.increase()
		p.flash()
		if p.checkAllFlashed() {
			break
		}
		// p.print()
		// fmt.Println("============================")

		// if runcounter >= 100 {
		// 	break
		// }
	}
	p.solution2 = uint64(runcounter)
}

func (p *Day11) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day11) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
