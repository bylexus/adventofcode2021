package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 21 -
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
)

type Day21 struct {
	solution1 uint64
	solution2 uint64
	playerPos []uint64
	playerSum []uint64
	actPlayer int
	die       uint64
}

func (p *Day21) GetName() string {
	return "xxx"
}

func (p *Day21) Init() {
	// Read input
	// input := lib.SplitLinesByRegex(lib.ReadInputLines("input/day21-test.txt"), "position: ")
	input := lib.SplitLinesByRegex(lib.ReadInputLines("input/day21-input.txt"), "position: ")
	p.die = 1
	p.playerPos = make([]uint64, 2, 2)
	p.playerSum = make([]uint64, 2, 2)
	p.actPlayer = 0
	p.playerPos[0] = uint64(lib.ToInt(input[0][1]))
	p.playerPos[1] = uint64(lib.ToInt(input[1][1]))
	fmt.Printf("%v\n", p.playerPos)
}

func (p *Day21) Run1() {
	dieCount := 0
	for {
		forward := (p.die-1)%100 + 1 + (p.die+1-1)%100 + 1 + (p.die+2-1)%100 + 1
		p.die = (p.die+3-1)%100 + 1
		dieCount += 3
		p.playerPos[p.actPlayer] = (p.playerPos[p.actPlayer]+forward-1)%10 + 1
		p.playerSum[p.actPlayer] += p.playerPos[p.actPlayer]

		if p.playerSum[p.actPlayer] >= 1000 {
			break
		}
		p.actPlayer = (p.actPlayer + 1) % 2
	}

	p.solution1 = uint64(dieCount * int(lib.MinUInt64(p.playerSum[0], p.playerSum[1])))
}

func (p *Day21) Run2() {
	p.solution2 = 0
}

func (p *Day21) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day21) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
