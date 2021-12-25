package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 25 - Sea Cucumber
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
	"alexi.ch/aoc2021/lib/types"
)

type Day25Map map[types.Point]rune

type Day25 struct {
	solution1 int
	solution2 uint64

	groundMap Day25Map
	mapWidth  int
	mapHeight int
}

func (p *Day25) GetName() string {
	return "Sea Cucumber"
}

func (p *Day25) Init() {
	// Read input
	// input := lib.ReadInputLines("input/day25-test.txt")
	// input := lib.ReadInputLines("input/day25-test1.txt")
	input := lib.ReadInputLines("input/day25-input.txt")

	p.groundMap = make(Day25Map)

	p.mapHeight = len(input)
	p.mapWidth = len(input[0])
	for y, line := range input {
		for x, r := range line {
			if r == '>' || r == 'v' {
				// skip emtpy spaces, they will not be stored in the map
				p.groundMap[types.Point{X: x, Y: y}] = r
			}
		}
	}
}

func (p *Day25) moveCucumbers() (haveMoved bool) {
	newMap := make(Day25Map)

	// process right-heading cucumbers:
	for pos, c := range p.groundMap {
		if c == 'v' {
			newMap[pos] = c
			continue
		}
		nextPos := types.Point{X: (pos.X + 1) % p.mapWidth, Y: pos.Y}
		_, exists := p.groundMap[nextPos]
		if exists == true {
			// next pos is not free, stay:
			newMap[pos] = c
		} else {
			// next pos is free, move:
			haveMoved = true
			newMap[nextPos] = c
		}
	}

	// copy back double-buffer map:
	p.groundMap = newMap
	newMap = make(Day25Map)

	// process down-heading cucumbers:
	for pos, c := range p.groundMap {
		if c == '>' {
			newMap[pos] = c
			continue
		}
		nextPos := types.Point{Y: (pos.Y + 1) % p.mapHeight, X: pos.X}
		_, exists := p.groundMap[nextPos]
		if exists == true {
			// next pos is not free, stay:
			newMap[pos] = c
		} else {
			// next pos is free, move:
			haveMoved = true
			newMap[nextPos] = c
		}
	}

	// copy back double-buffer map:
	p.groundMap = newMap

	return
}

func (p *Day25) printMap() {
	for y := 0; y < p.mapHeight; y++ {
		for x := 0; x < p.mapWidth; x++ {
			pos := types.Point{X: x, Y: y}
			c, exists := p.groundMap[pos]
			if exists == true {
				fmt.Printf("%v", string(c))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (p *Day25) Run1() {
	counter := 0
	for {
		haveMoved := p.moveCucumbers()
		counter++
		if !haveMoved {
			break
		}
	}
	p.solution1 = counter
}

func (p *Day25) Run2() {
	p.solution2 = 0
}

func (p *Day25) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day25) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
