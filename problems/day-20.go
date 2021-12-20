package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 20 - Trench Map
// ----------------------------------------------------------------------------

import (
	"fmt"
	"strconv"

	"alexi.ch/aoc2021/lib"
	"alexi.ch/aoc2021/lib/types"
)

type Day20 struct {
	solution1 uint64
	solution2 uint64
	input     []string
	algo      string
	image     map[types.Point]rune
	minX      int
	maxX      int
	minY      int
	maxY      int
	emptyVal  rune
}

func (p *Day20) GetName() string {
	return "Trench Map"
}

func (p *Day20) Init() {
	// Read input
	// input := lib.ReadInputLines("input/day20-test.txt")
	input := lib.ReadInputLines("input/day20-input.txt")

	p.algo = input[0]

	p.minX = 0
	p.minY = 0
	p.maxY = len(input[1:]) - 1
	p.maxX = len(input[1]) - 1
	p.image = make(map[types.Point]rune)
	p.emptyVal = '.'

	for y, line := range input[1:] {
		for x, r := range line {
			p.image[types.Point{X: x, Y: y}] = r
		}
	}
}

func (p *Day20) getAt(x, y int) rune {
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
			fmt.Print(string(p.image[types.Point{X: x, Y: y}]))
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (p *Day20) calcNewVal(x, y int) rune {
	binstr := ""
	// create binary nr from 9-patch around x/y:
	for dy := y - 1; dy <= y+1; dy++ {
		for dx := x - 1; dx <= x+1; dx++ {
			val := p.getAt(dx, dy)
			if val == '.' {
				binstr += "0"
			} else {
				binstr += "1"
			}
		}
	}
	nr, err := strconv.ParseInt(binstr, 2, 16)
	if err != nil {
		panic(err)
	}
	return rune(p.algo[nr])
}

func (p *Day20) calcRound() {
	newImg := make(map[types.Point]rune)
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
	if p.emptyVal == '.' {
		p.emptyVal = rune(p.algo[0])
	} else {
		p.emptyVal = rune(p.algo[len(p.algo)-1])
	}
}

func (p *Day20) calcOnPixels() int {
	pixels := 0
	for _, r := range p.image {
		if r == '#' {
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
