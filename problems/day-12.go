package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 12 - Passage Pathing
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
	name    string
	visited bool
	big     bool
}

type Day12 struct {
	solution1 int
	solution2 int

	edges []*Day12Edge
	caves map[string]*Day12Cave
}

func (p *Day12) GetName() string {
	return "Passage Pathing"
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
		cave1 := Day12Cave{name: line[0], visited: false}
		cave2 := Day12Cave{name: line[1], visited: false}
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

// finds the neighbour caves of the given cave
// which are still open for visit.
// a cave is still visitable if:
// - it's the end cave
// - it's a big cave
// - it's a non-visited small cave
func (p *Day12) findNextCaves(cave *Day12Cave) []string {
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
			} else if otherCave.visited == false {
				caves = append(caves, otherCave.name)
			}
		}
	}
	return caves
}

func (p *Day12) resetAllCaves() {
	for _, cave := range p.caves {
		cave.visited = false
	}
}

func (p *Day12) resetCaves(caves []string) {
	for _, cave := range caves {
		p.caves[cave].visited = false
	}
}

// takes a cave, walks through the end,
// and returns an array of paths (array of array of cave names)
// possible to the end
// This is a recursive graph walk algorithm, with a twist:
// before walking each sub-path, we have to reset the
// caves from the last sub-path (visit status), as they will
// possibly be visited again.
func (p *Day12) walk(cave *Day12Cave) [][]string {
	cave.visited = true
	nextCaves := p.findNextCaves(cave)
	paths := make([][]string, 0)
	for _, next := range nextCaves {
		nextCave := p.caves[next]
		if next == "end" {
			paths = append(paths, []string{"end"})
		} else {
			subpaths := p.walk(nextCave)
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

func (p *Day12) Run1() {
	start := p.caves["start"]
	paths := p.walk(start)
	p.solution1 = len(paths)
}

func (p *Day12) Run2() {
	// for the 2nd part, I use a little hack:
	// for each small cave, I re-run the evaluation again and collect the paths.
	// In each run, I add a copy of a single small cave, which then acts as a
	// "2nd visit" cave. So instead of adopting the walking algorithm, I simply
	// trick it for a single cave.
	allPaths := make([][]string, 0)
	nrOfEdges := len(p.edges)

	for _, cave := range p.caves {
		if cave.big == false && cave.name != "start" && cave.name != "end" {
			// for each small cave, we run the graph walk, but with an additional
			// small cave
			p.resetAllCaves()

			// creating a copy of the actual small cave, including edges that point to it:
			// we give it a unique name, so that it can be addressed in the cave map:
			virtualCave := Day12Cave{name: "xxxxxxxx", visited: false}
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

			// copy done, now do the graph walk:
			start := p.caves["start"]
			paths := p.walk(start)

			// now the returned paths contain the virtual name instead of the real (originam) name.
			// we have to replace the virtual name in the paths with the original one:
			for i, path := range paths {
				for j, part := range path {
					if part == virtualCave.name {
						paths[i][j] = cave.name
					}
				}
			}
			// add the paths from this run to our path memory:
			allPaths = append(allPaths, paths...)

			// remove additional edges and delete the virtual node for the next run:
			p.edges = p.edges[:nrOfEdges]
			delete(p.caves, virtualCave.name)
		}
	}

	// now we have to filter duplicate paths: I simply hash the paths and add it to a map,
	// to create a "set" of paths:
	pathMap := make(map[string][]string)
	for _, path := range allPaths {
		key := fmt.Sprintf("%#v", path)
		pathMap[key] = path
	}

	// voilà:
	p.solution2 = len(pathMap)
}

func (p *Day12) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day12) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
