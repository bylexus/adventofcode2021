package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 06 - Lanternfish
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
)

type Day06 struct {
	input     []uint64
	solution1 uint64
	solution2 uint64
}

func (p *Day06) GetName() string {
	return "Lanternfish"
}

func (p *Day06) Init() {
	// Read input
	// lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day06-test.txt"), `,`)
	lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day06-input.txt"), `,`)
	p.input = lib.ParseUIntLines(lines[0])

}

func calcAges(input []uint64) []uint64 {
	first := input[0]
	input = input[1:]
	input[6] += first
	input = append(input, first)
	return input
}

// Idea:
// we keep a map of nr of fishes per "age" (per day):
// each day, there is a list of ages and how many fishes are _that_old.
// each day, this sums rotate:
// - fishes with count 0 got added to the count of day 6 (will be re-started)
// - rotate the list (as each fish must be decrease to 0)
// - at the same time, those fishes produce the same amount of new fishes, so add them to the end, too
func (p *Day06) Run1() {
	ages := make([]uint64, 9)
	for _, d := range p.input {
		ages[d] += 1
	}

	for i := 0; i < 80; i++ {
		ages = calcAges(ages)
	}
	var sum uint64 = 0
	for _, age := range ages {
		sum += age
	}
	p.solution1 = sum
}

func (p *Day06) Run2() {
	ages := make([]uint64, 9)
	for _, d := range p.input {
		ages[d] += 1
	}

	for i := 0; i < 256; i++ {
		ages = calcAges(ages)
	}
	var sum uint64 = 0
	for _, age := range ages {
		sum += age
	}
	p.solution2 = sum
}

func (p *Day06) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day06) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
