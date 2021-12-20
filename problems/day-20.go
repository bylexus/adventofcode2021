package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 20 - Trench Map
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
	"alexi.ch/aoc2021/lib/types"
)

type Day20 struct {
	solution1 uint64
	solution2 uint64
	input     []string
	algo      []byte
	image     map[types.Point]byte
	minX      int
	maxX      int
	minY      int
	maxY      int
	emptyVal  byte
}

func (p *Day20) GetName() string {
	return "Trench Map"
}

func (p *Day20) Init() {
	// Read input
	// input := lib.ReadInputLines("input/day20-test.txt")
	input := lib.ReadInputLines("input/day20-input.txt")
	p.algo = make([]byte, 0, len(input[0]))
	for _, r := range input[0] {
		if r == '.' {
			p.algo = append(p.algo, 0)
		} else {
			p.algo = append(p.algo, 1)

		}
	}

	p.minX = 0
	p.minY = 0
	p.maxY = len(input[1:]) - 1
	p.maxX = len(input[1]) - 1
	p.image = make(map[types.Point]byte)
	p.emptyVal = 0

	for y, line := range input[1:] {
		for x, r := range line {
			if r == '.' {
				p.image[types.Point{X: x, Y: y}] = 0
			} else {
				p.image[types.Point{X: x, Y: y}] = 1
			}
		}
	}
}

func (p *Day20) getAt(x, y int) byte {
	val, present := p.image[types.Point{X: x, Y: y}]
	if present != true {
		return p.emptyVal
	} else {
		return val
	}
}

func (p *Day20) printImg() {
	for y := p.minY; y <= p.maxY; y++ {
		for x := p.minX; x <= p.maxX; x++ {
			val := p.image[types.Point{X: x, Y: y}]
			if val == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (p *Day20) calcNewVal(x, y int) byte {
	nr := 0
	// create binary nr from 9-patch around x/y:
	for dy := y - 1; dy <= y+1; dy++ {
		for dx := x - 1; dx <= x+1; dx++ {
			val := p.getAt(dx, dy)
			if val == 0 {
				nr = (nr << 1) | 0
			} else {
				nr = (nr << 1) | 1
			}
		}
	}
	return p.algo[nr]
}

func (p *Day20) calcRound() {
	newImg := make(map[types.Point]byte)
	p.minX--
	p.minY--
	p.maxX++
	p.maxY++
	for y := p.minY; y <= p.maxY; y++ {
		for x := p.minX; x <= p.maxX; x++ {
			newVal := p.calcNewVal(x, y)
			newImg[types.Point{X: x, Y: y}] = newVal
		}
	}
	p.image = newImg
	if p.emptyVal == 0 {
		p.emptyVal = p.algo[0]
	} else {
		p.emptyVal = p.algo[len(p.algo)-1]
	}
}

func (p *Day20) calcOnPixels() int {
	pixels := 0
	for _, r := range p.image {
		if r == 1 {
			pixels++
		}
	}
	return pixels
}

func (p *Day20) Run1() {
	for i := 0; i < 2; i++ {
		p.calcRound()
	}

	p.solution1 = uint64(p.calcOnPixels())
}

func (p *Day20) Run2() {
	// only calc 48 rounds (2 already done in part 1)
	for i := 0; i < 48; i++ {
		p.calcRound()
	}

	p.solution2 = uint64(p.calcOnPixels())
}

func (p *Day20) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day20) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
