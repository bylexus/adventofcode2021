package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 09 - Smoke Basin
// ----------------------------------------------------------------------------

import (
	"fmt"
	"sort"
	"strconv"

	"alexi.ch/aoc2021/lib"
)

type Day09LowPoints map[uint64]uint64

type Day09 struct {
	input     [][]string
	solution1 uint64
	solution2 uint64

	heatmap   [][]uint64
	lowpoints Day09LowPoints
}

func (p *Day09) GetName() string {
	return "Smoke Basin"
}

func (p *Day09) Init() {
	// Read input
	// lines := lib.ReadInputLines("input/day09-test.txt")
	lines := lib.ReadInputLines("input/day09-input.txt")
	p.heatmap = make([][]uint64, len(lines))
	for y, line := range lines {
		p.heatmap[y] = make([]uint64, len(line))
		for x, height := range line {
			h, err := strconv.Atoi(string(height))
			if err != nil {
				panic(err)
			}
			p.heatmap[y][x] = uint64(h)
		}
	}
}

func (p *Day09) isLowPoint(x int, y int, val uint64, hmap [][]uint64) bool {
	// top
	if y > 0 && hmap[y-1][x] <= val {
		return false
	}
	// right
	if x < len(hmap[y])-1 && hmap[y][x+1] <= val {
		return false
	}
	// bottom
	if y < len(hmap)-1 && hmap[y+1][x] <= val {
		return false
	}
	// left
	if x > 0 && hmap[y][x-1] <= val {
		return false
	}

	return true
}

func (p *Day09) Run1() {
	p.lowpoints = make(Day09LowPoints, 0)
	for y := 0; y < len(p.heatmap); y++ {
		for x := 0; x < len(p.heatmap[y]); x++ {
			act := p.heatmap[y][x]
			if p.isLowPoint(x, y, act, p.heatmap) {
				p.lowpoints[uint64(y*10000+x)] = act
			}
		}
	}
	var riskSum uint64 = 0
	for _, p := range p.lowpoints {
		riskSum += p + 1
	}
	p.solution1 = riskSum
}

// Flood-fills a pool "bottom up": start from a pool (lowest location), then
// flood-fill (depth-first search) all adjacent tiles until the top is reached.
func (p *Day09) calcPool(pool uint64, x int, y int, heatmap [][]uint64, pointBasinMap map[uint64]uint64) {
	pointKey := uint64(10000*y + x)
	actVal := heatmap[y][x]
	if actVal == 9 {
		panic("Oops: 9 should not be processed")
	}
	// mark with pool, also marks visited:
	pointBasinMap[pointKey] = pool

	// flood in all directions possible:
	// up
	if y > 0 {
		checkX := x
		checkY := y - 1
		checkKey := uint64(10000*checkY + checkX)
		_, defined := pointBasinMap[checkKey]
		checkVal := heatmap[checkY][checkX]
		if defined == false && checkVal >= actVal && checkVal < 9 {
			p.calcPool(pool, checkX, checkY, heatmap, pointBasinMap)
		}
	}

	// left
	if x > 0 {
		checkX := x - 1
		checkY := y
		checkKey := uint64(10000*checkY + checkX)
		_, defined := pointBasinMap[checkKey]
		checkVal := heatmap[checkY][checkX]
		if defined == false && checkVal >= actVal && checkVal < 9 {
			p.calcPool(pool, checkX, checkY, heatmap, pointBasinMap)
		}
	}
	// bottom
	if y < len(heatmap)-1 {
		checkX := x
		checkY := y + 1
		checkKey := uint64(10000*checkY + checkX)
		_, defined := pointBasinMap[checkKey]
		checkVal := heatmap[checkY][checkX]
		if defined == false && checkVal >= actVal && checkVal < 9 {
			p.calcPool(pool, checkX, checkY, heatmap, pointBasinMap)
		}
	}
	// right
	if x < len(heatmap[y])-1 {
		checkX := x + 1
		checkY := y
		checkKey := uint64(10000*checkY + checkX)
		_, defined := pointBasinMap[checkKey]
		checkVal := heatmap[checkY][checkX]
		if defined == false && checkVal >= actVal && checkVal < 9 {
			p.calcPool(pool, checkX, checkY, heatmap, pointBasinMap)
		}
	}
}

// Approach:
// I start from each Basin, and flood it upwards, until the peak is reached :-)
// this uses a flood-fill / depth-first search approach.
func (p *Day09) Run2() {
	// we map each point to a basin nr. The basin nr is the x/y coord of the low point it belongs to.
	// [point(10000*y+x) => pool
	pointBasinMap := make(map[uint64]uint64)

	// we start with the lowpoints itself, which form a basin each:
	for pool := range p.lowpoints {
		// pool ist the pool ident, while lowpoint is the height of the basin bottom
		x := int(pool % 10000)
		y := int(pool / 10000)
		pointBasinMap[pool] = pool
		// start flooding:
		p.calcPool(pool, x, y, p.heatmap, pointBasinMap)
	}

	// calc pools and pool sizes:
	poolSizes := make(map[uint64]int64)
	for _, pool := range pointBasinMap {
		poolSizes[pool]++
	}
	sizesOnly := make([]int, 0)
	for _, size := range poolSizes {
		sizesOnly = append(sizesOnly, int(size))
	}
	sort.Ints(sizesOnly)
	var total uint64 = 1
	for _, size := range sizesOnly[len(sizesOnly)-3:] {
		total *= uint64(size)
	}

	p.solution2 = uint64(total)
}

func (p *Day09) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day09) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
