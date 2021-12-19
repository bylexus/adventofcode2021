package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 07 - The Treachery of Whales
// ----------------------------------------------------------------------------

import (
	"fmt"
	"math"

	"alexi.ch/aoc2021/lib"
)

type Day07 struct {
	input     []uint64
	inMin     uint64
	inMax     uint64
	solution1 uint64
	solution2 uint64
}

func (p *Day07) GetName() string {
	return "The Treachery of Whales"
}

func (p *Day07) Init() {
	// Read input
	// lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day07-test.txt"), `,`)
	lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day07-input.txt"), `,`)
	p.input = lib.ParseUIntLines(lines[0])
	for _, nr := range p.input {
		if nr < p.inMin {
			p.inMin = nr
		}
		if nr > p.inMax {
			p.inMax = nr
		}
	}
}

func (p *Day07) Run1() {
	var minFuel uint64 = math.MaxUint64
	var fueldiff uint64 = 0
	for pos := p.inMin; pos <= p.inMax; pos++ {
		fueldiff = 0
		for _, checkPos := range p.input {
			fueldiff += lib.MaxUInt64(pos, checkPos) - lib.MinUInt64(pos, checkPos)
		}
		if fueldiff < minFuel {
			minFuel = fueldiff
		}
	}
	p.solution1 = minFuel
}

func (p *Day07) Run2() {
	var minFuel uint64 = math.MaxUint64
	var fueldiff uint64 = 0
	for pos := p.inMin; pos <= p.inMax; pos++ {
		fueldiff = 0
		for _, checkPos := range p.input {
			// gausssche Summenformel: (n^2 + n) / 2
			// where n = diff(pos, checkpos)
			diff := lib.MaxUInt64(pos, checkPos) - lib.MinUInt64(pos, checkPos)
			fueldiff += (diff*diff + diff) / 2
		}
		if fueldiff < minFuel {
			minFuel = fueldiff
		}
	}

	p.solution2 = minFuel
}

func (p *Day07) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day07) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
