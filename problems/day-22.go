package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 22 -
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
	"alexi.ch/aoc2021/lib/types"
)

type Day22Instr struct {
	turnOn                 bool
	x1, x2, y1, y2, z1, z2 int
}

type Day22Cuboids map[types.Point]bool

type Day22 struct {
	solution1 int
	solution2 uint64

	instructions []Day22Instr
	cuboids      Day22Cuboids
}

func (p *Day22) GetName() string {
	return "xxx"
}

func (p *Day22) Init() {
	// Read input
	// input := lib.ParseGroupMatch(lib.ReadInputLines("input/day22-test.txt"), `(.*) x=(.*)\.\.(.*),y=(.*)\.\.(.*),z=(.*)\.\.(.*)`)
	input := lib.ParseGroupMatch(lib.ReadInputLines("input/day22-input.txt"), `(.*) x=(.*)\.\.(.*),y=(.*)\.\.(.*),z=(.*)\.\.(.*)`)
	p.instructions = make([]Day22Instr, 0, len(input))
	p.cuboids = make(Day22Cuboids)
	for _, i := range input {
		instr := Day22Instr{
			x1: lib.ToInt(i[2]),
			x2: lib.ToInt(i[3]),
			y1: lib.ToInt(i[4]),
			y2: lib.ToInt(i[5]),
			z1: lib.ToInt(i[6]),
			z2: lib.ToInt(i[7]),
		}
		if i[1] == "on" {
			instr.turnOn = true
		}
		p.instructions = append(p.instructions, instr)

	}
}

func (p *Day22) Run1() {
	for _, instr := range p.instructions {
		// minX := lib.WithinInt(instr.x1, -50, 50)
		// maxX := lib.WithinInt(instr.x2, -50, 50)
		// minY := lib.WithinInt(instr.y1, -50, 50)
		// maxY := lib.WithinInt(instr.y2, -50, 50)
		// minZ := lib.WithinInt(instr.z1, -50, 50)
		// maxZ := lib.WithinInt(instr.z2, -50, 50)

		for z := instr.z1; z <= instr.z2; z++ {
			if z < -50 || z > 50 {
				continue
			}
			for y := instr.y1; y <= instr.y2; y++ {
				if y < -50 || y > 50 {
					continue
				}
				for x := instr.x1; x <= instr.x2; x++ {
					if x < -50 || x > 50 {
						continue
					}
					p.cuboids[types.Point{X: x, Y: y, Z: z}] = instr.turnOn
				}
			}
		}
	}
	countOn := 0
	for _, on := range p.cuboids {
		if on == true {
			countOn++
		}
	}
	p.solution1 = countOn
}

func (p *Day22) Run2() {
	p.solution2 = 0
}

func (p *Day22) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day22) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
