package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 17 - Trick Shot
// ----------------------------------------------------------------------------

import (
	"fmt"
	"math"

	"alexi.ch/aoc2021/lib"
)

type Day17 struct {
	solution1 int
	solution2 int

	targetMinX int
	targetMinY int
	targetMaxX int
	targetMaxY int
}

func (p *Day17) GetName() string {
	return "AoC 2021 - Day 17 - Trick Shot"
}

func (p *Day17) Init() {
	// Read input
	// lines := lib.ReadInputLines("input/day17-test.txt")
	lines := lib.ReadInputLines("input/day17-input.txt")
	data := lib.ParseGroupMatch(lines, `=(-?[0-9]+)\.\.(-?[0-9]+).*=(-?[0-9]+)\.\.(-?[0-9]+)`)
	p.targetMinX = lib.ToInt(data[0][1])
	p.targetMaxX = lib.ToInt(data[0][2])
	p.targetMinY = lib.ToInt(data[0][3])
	p.targetMaxY = lib.ToInt(data[0][4])
}

// Just brute-force it here: try different start x/y velocities,
// and calculate each step until we are under the target area or beyond.
// can do both solutions in one run.
func (p *Day17) Run1() {
	maxY := math.MinInt
	maxStartYVel := 500 // luky punch:  I hope more is not needed?
	hitCount := 0
	for startYvel := p.targetMinY; startYvel <= maxStartYVel; startYvel++ {
		for startXvel := 1; startXvel <= p.targetMaxX; startXvel++ {
			xVel := startXvel
			yVel := startYvel
			actX := 0
			actY := 0
			localMaxY := math.MinInt
			for {
				if actY < p.targetMinY || actX > p.targetMaxX {
					break
				}
				actX = actX + xVel
				actY = actY + yVel
				if actY > localMaxY {
					localMaxY = actY
				}

				if actX >= p.targetMinX && actX <= p.targetMaxX && actY >= p.targetMinY && actY <= p.targetMaxY {
					// in target area, found.
					hitCount++
					if localMaxY > maxY {
						maxY = localMaxY
					}
					break
				}

				if xVel != 0 {
					xVel--
				}
				yVel--
			}
		}
	}
	p.solution1 = maxY
	p.solution2 = hitCount
}

func (p *Day17) Run2() {
	// pass - already solved in run 1
}

func (p *Day17) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day17) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
