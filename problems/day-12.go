package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 12 - xxx
// ----------------------------------------------------------------------------

import (
	"fmt"
	"strings"

	"alexi.ch/aoc2021/lib"
)

type Day12Edge struct {
	cave1 string
	cave2 string
}

type Day12Cave struct {
	name            string
	remainingVisits int
	initialVisits   int
	big             bool
}

type Day12 struct {
	solution1 int
	solution2 int

	edges []*Day12Edge
	caves map[string]*Day12Cave
}

func (p *Day12) GetName() string {
	return "AoC 2021 - Day 12 - xxx"
}

func (p *Day12) Init() {
	// Read input
	// lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day12-test.txt"), "-")
	// lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day12-test2.txt"), "-")
	// lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day12-test3.txt"), "-")
	lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day12-input.txt"), "-")

	p.edges = make([]*Day12Edge, 0)
	p.caves = make(map[string]*Day12Cave)

	for _, line := range lines {
		cave1 := Day12Cave{name: line[0], remainingVisits: 1, initialVisits: 1}
		cave2 := Day12Cave{name: line[1], remainingVisits: 1, initialVisits: 1}
		p.caves[line[0]] = &cave1
		p.caves[line[1]] = &cave2
		if strings.ToUpper(cave1.name) == cave1.name {
			cave1.big = true
		}
		if strings.ToUpper(cave2.name) == cave2.name {
			cave2.big = true
		}
		p.edges = append(p.edges, &Day12Edge{cave1: cave1.name, cave2: cave2.name})
	}
}

func (p *Day12) findNextCaves(cave *Day12Cave) []string {
	// a cave is possible if:
	// - it's the end cave
	// - it's a big cave
	// - it's a non-visited small cave
	caves := make([]string, 0)
	for _, edge := range p.edges {
		var other string
		if edge.cave1 == cave.name {
			other = edge.cave2
		}
		if edge.cave2 == cave.name {
			other = edge.cave1
		}
		otherCave, present := p.caves[other]
		if present == true {
			if otherCave.name == "start" {
				// never visit start again
				continue
			}
			if otherCave.name == "end" {
				caves = append(caves, otherCave.name)
			} else if otherCave.big == true {
				caves = append(caves, otherCave.name)
			} else if otherCave.remainingVisits > 0 {
				caves = append(caves, otherCave.name)
			}
		}
	}
	return caves
}

func (p *Day12) resetAllCaves() {
	for _, cave := range p.caves {
		cave.remainingVisits = 1
		cave.initialVisits = 1
	}
}

func (p *Day12) resetCaves(caves []string) {
	for _, cave := range caves {
		// p.caves[cave].remainingVisits = p.caves[cave].initialVisits
		p.caves[cave].remainingVisits = 1
		// p.caves[cave].remainingVisits++
		// p.caves[cave].remainingVisits = lib.MinInt(p.caves[cave].remainingVisits, p.caves[cave].initialVisits)
	}
}

// takes a cave, walks through the end,
// and returns an array of paths (array of array of cave names)
// possible to the end
func (p *Day12) walk(cave *Day12Cave) [][]string {
	cave.remainingVisits--
	nextCaves := p.findNextCaves(cave)
	paths := make([][]string, 0)
	for _, next := range nextCaves {
		nextCave := p.caves[next]
		if next == "end" {
			paths = append(paths, []string{"end"})
		} else {
			subpaths := p.walk(nextCave)
			// fmt.Printf("Subpaths: %#v\n", subpaths)
			if len(subpaths) > 0 {
				for _, s := range subpaths {
					s = append([]string{nextCave.name}, s...)
					p.resetCaves(s)
					paths = append(paths, s)
				}
			}
		}
	}
	return paths
}

func (p *Day12) printPath(paths []string) {
	fmt.Print("start, ")
	for _, path := range paths {
		fmt.Printf("%v, ", path)
	}
	fmt.Println()
}

func (p *Day12) Run1() {
	start := p.caves["start"]
	paths := p.walk(start)
	// for i, path := range paths {
	// 	fmt.Printf("%v: ", i+1)
	// 	p.printPath(path)
	// }

	p.solution1 = len(paths)
}

func (p *Day12) Run2() {
	// do it for each small cave,
	// with another small cave as 2 times visible:
	allPaths := make([][]string, 0)
	nrOfEdges := len(p.edges)
	for _, cave := range p.caves {
		p.resetAllCaves()
		if cave.big == false && cave.name != "start" && cave.name != "end" {
			virtualCave := Day12Cave{name: "xxxxxxxx", remainingVisits: 1, initialVisits: 1}
			caveEdges := make([]*Day12Edge, 0)
			for _, edge := range p.edges {
				if edge.cave1 == cave.name {
					caveEdges = append(caveEdges, &Day12Edge{cave1: virtualCave.name, cave2: edge.cave2})
				} else if edge.cave2 == cave.name {
					caveEdges = append(caveEdges, &Day12Edge{cave1: edge.cave1, cave2: virtualCave.name})
				}
			}
			p.edges = append(p.edges, caveEdges...)
			p.caves[virtualCave.name] = &virtualCave

			start := p.caves["start"]
			paths := p.walk(start)
			for i, path := range paths {
				for j, part := range path {
					if part == virtualCave.name {
						paths[i][j] = cave.name
					}
				}
			}
			allPaths = append(allPaths, paths...)

			// remove additional edges:
			p.edges = p.edges[:nrOfEdges]
			delete(p.caves, virtualCave.name)
		}
	}

	pathMap := make(map[string][]string)
	for _, path := range allPaths {
		key := fmt.Sprintf("%#v", path)
		pathMap[key] = path
	}

	p.solution2 = len(pathMap)
}

func (p *Day12) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day12) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
