package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 15 - Chiton
// ----------------------------------------------------------------------------

import (
	"fmt"
	"strconv"

	"alexi.ch/aoc2021/lib"
)

type Day15Location struct {
	x         int
	y         int
	value     int
	visited   bool
	totalRisk int
}

type Day15 struct {
	solution1 int
	solution2 int

	cave   map[lib.Point]*Day15Location
	width  int
	height int

	queue map[lib.Point]*Day15Location
}

func (p *Day15) GetName() string {
	return "AoC 2021 - Day 15 - Chiton"
}

func (p *Day15) Init() {
	// Read input
	// lines := lib.ReadInputLines("input/day15-test.txt")
	lines := lib.ReadInputLines("input/day15-input.txt")

	p.cave = make(map[lib.Point]*Day15Location)
	p.queue = make(map[lib.Point]*Day15Location)

	// create cave nodes from input
	for y, line := range lines {
		if p.height < y+1 {
			p.height = y + 1
		}
		for x, r := range line {
			if p.width < x+1 {
				p.width = x + 1
			}
			val, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			loc := Day15Location{
				x:         x,
				y:         y,
				value:     val,
				visited:   false,
				totalRisk: -1,
			}
			p.cave[lib.Point{X: x, Y: y}] = &loc
		}
	}

}

func (p *Day15) printCave() {
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			loc := p.cave[lib.Point{X: x, Y: y}]
			fmt.Printf("%v", loc.value)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (p *Day15) printTotRiskMap() {
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			fmt.Printf("%4d", p.cave[lib.Point{X: x, Y: y}].totalRisk)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (p *Day15) getNeighbours(node *Day15Location) []*Day15Location {
	neighbours := make([]*Day15Location, 0)
	// top
	n, ok := p.cave[lib.Point{X: node.x, Y: node.y - 1}]
	if ok == true {
		neighbours = append(neighbours, n)
	}
	// right
	n, ok = p.cave[lib.Point{X: node.x + 1, Y: node.y}]
	if ok == true {
		neighbours = append(neighbours, n)
	}
	// bottom
	n, ok = p.cave[lib.Point{X: node.x, Y: node.y + 1}]
	if ok == true {
		neighbours = append(neighbours, n)
	}
	// left
	n, ok = p.cave[lib.Point{X: node.x - 1, Y: node.y}]
	if ok == true {
		neighbours = append(neighbours, n)
	}
	return neighbours
}

// pops the entry with the lowest total risk from the queue
func (p *Day15) popQueue() *Day15Location {
	minRisk := -1
	var minEntry *Day15Location
	for _, n := range p.queue {
		if minRisk == -1 || n.totalRisk < minRisk {
			minRisk = n.totalRisk
			minEntry = n
		}
	}
	if minEntry == nil {
		panic("Oops! cannot happen")
	}

	// pop and return:
	delete(p.queue, lib.Point{X: minEntry.x, Y: minEntry.y})
	return minEntry
}

// Here, Dijkstra does its magic :-)
// classical dijkstra, no strings attached
func (p *Day15) examineNode(node *Day15Location) {
	node.visited = true

	neighbours := p.getNeighbours(node)
	for _, n := range neighbours {
		// calc new risk value:
		// cost for the route is the next node's risk value,
		// while the "distance" (total risk) to this node is
		// the sum of previous paths (totalRisk) + the node's value.
		// If the new total risk value is lower thatn the already stored one,
		// we have found a lower-risk route.
		sumRisk := n.value + node.totalRisk
		if n.totalRisk == -1 || sumRisk < n.totalRisk {
			n.totalRisk = sumRisk
		}
		// enqueue, if not yet visited:
		if n.visited == false {
			p.queue[lib.Point{X: n.x, Y: n.y}] = n
		}
	}

	// now work on the queue until it's empty:
	// process unvisited nodes in nearest order:
	for {
		if len(p.queue) == 0 {
			break
		}
		// nextEl is the element with the lowest path until now
		nextEl := p.popQueue()
		p.examineNode(nextEl)
	}
}

// Seems to be a straight-forward dijkstra.
// Let's see.
func (p *Day15) Run1() {
	startNode := p.cave[lib.Point{X: 0, Y: 0}]
	startNode.totalRisk = 0
	p.examineNode(startNode)
	p.solution1 = p.cave[lib.Point{X: p.width - 1, Y: p.height - 1}].totalRisk
}

// yes, straight-forward dijkstra so far.
// now we just have to enlarge the cave system:
func (p *Day15) Run2() {
	// Reset existing cave
	p.queue = make(map[lib.Point]*Day15Location)
	for _, n := range p.cave {
		n.totalRisk = -1
		n.visited = false

	}
	// expand 4x in both dirs
	for ty := 0; ty < 5; ty++ {
		for tx := 0; tx < 5; tx++ {
			if tx == 0 && ty == 0 {
				continue
			}
			// copy grid to new location
			for y := 0; y < p.height; y++ {
				for x := 0; x < p.width; x++ {
					newNode := Day15Location{
						x:         tx*p.width + x,
						y:         ty*p.height + y,
						totalRisk: -1,
						visited:   false,
						value:     p.cave[lib.Point{X: x, Y: y}].value + tx + ty,
					}
					if newNode.value > 9 {
						newNode.value = 1 + (newNode.value % 10)
					}
					p.cave[lib.Point{X: newNode.x, Y: newNode.y}] = &newNode
				}
			}
		}
	}
	p.height = 5 * p.height
	p.width = 5 * p.width

	// and run again:
	startNode := p.cave[lib.Point{X: 0, Y: 0}]
	startNode.totalRisk = 0
	p.examineNode(startNode)
	p.solution2 = p.cave[lib.Point{X: p.width - 1, Y: p.height - 1}].totalRisk
}

func (p *Day15) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day15) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
