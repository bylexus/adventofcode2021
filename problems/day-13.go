package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 13 - Transparent Origami
// ----------------------------------------------------------------------------

import (
	"fmt"
	"strconv"
	"strings"

	"alexi.ch/aoc2021/lib"
)

type Day13Map map[lib.Point]lib.Point

type Day13 struct {
	solution1 int
	solution2 string

	pointMap Day13Map
	folds    []lib.Point

	maxX int
	maxY int
}

func (p *Day13) GetName() string {
	return "AoC 2021 - Day 13 - Transparent Origami"
}

func (p *Day13) Init() {
	// Read input
	// lines := lib.ReadInputLines("input/day13-test.txt")
	lines := lib.ReadInputLines("input/day13-input.txt")

	p.pointMap = make(Day13Map)
	p.folds = make([]lib.Point, 0)

	for _, line := range lines {
		if len(line) > 0 {
			points := strings.Split(line, ",")
			if len(points) == 2 {
				x, err := strconv.Atoi(points[0])
				if err != nil {
					panic(err)
				}
				y, err := strconv.Atoi(points[1])
				if err != nil {
					panic(err)
				}
				newPoint := lib.Point{X: x, Y: y}
				p.pointMap[newPoint] = newPoint

				p.maxX = lib.MaxInt(p.maxX, x)
				p.maxY = lib.MaxInt(p.maxY, y)
			}
			foldInfo := strings.Split(line, "y=")
			if len(foldInfo) == 2 {
				y, err := strconv.Atoi(foldInfo[1])
				if err != nil {
					panic(err)
				}
				p.folds = append(p.folds, lib.Point{X: 0, Y: y})
			}
			foldInfo = strings.Split(line, "x=")
			if len(foldInfo) == 2 {
				x, err := strconv.Atoi(foldInfo[1])
				if err != nil {
					panic(err)
				}
				p.folds = append(p.folds, lib.Point{X: x, Y: 0})
			}
		}
	}
}

func (p *Day13) printMap() string {
	var out string = ""
	maxX := 0
	maxY := 0
	for _, point := range p.pointMap {
		maxX = lib.MaxInt(maxX, point.X)
		maxY = lib.MaxInt(maxY, point.Y)
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			_, defined := p.pointMap[lib.Point{X: x, Y: y}]
			if defined {
				out += fmt.Sprint("█")
			} else {
				out += fmt.Sprint(" ")
			}
		}
		out += fmt.Sprintln()
	}
	out += fmt.Sprintln()
	return out
}

func (p *Day13) fold(foldPoint lib.Point) {
	for key, point := range p.pointMap {
		if foldPoint.X > 0 && point.X > foldPoint.X {
			delete(p.pointMap, key)
			point.X = 2*foldPoint.X - point.X
		} else if foldPoint.Y > 0 && point.Y > foldPoint.Y {
			delete(p.pointMap, key)
			point.Y = 2*foldPoint.Y - point.Y
		}
		p.pointMap[point] = point
	}
}

func (p *Day13) Run1() {
	p.fold(p.folds[0])
	p.solution1 = len(p.pointMap)
}

func (p *Day13) Run2() {
	for _, fold := range p.folds[1:] {
		p.fold(fold)
	}
	fmt.Println(p.printMap())
	p.solution2 = "\n"
}

func (p *Day13) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day13) GetSolution2() string {
	return fmt.Sprintf("\n%v\n", p.solution2)
}
