package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 10 - Syntax Scoring
// ----------------------------------------------------------------------------

import (
	"fmt"
	"sort"

	"alexi.ch/aoc2021/lib"
)

type Day10 struct {
	input            [][]string
	solution1        uint64
	solution2        uint64
	lines            []string
	incompleteStacks [][]rune
}

func (p *Day10) GetName() string {
	return "Syntax Scoring"
}

func (p *Day10) Init() {
	// Read input
	// p.lines = lib.ReadInputLines("input/day10-test.txt")
	p.lines = lib.ReadInputLines("input/day10-input.txt")
}

func (p *Day10) Run1() {
	// build a stack - push to stack when an OPENING char appears,
	// pop when a CLOSING char appears. The popped char must
	// be the corresponding opening char.
	openCloseMap := map[rune]rune{'<': '>', '(': ')', '{': '}', '[': ']'}
	illegalChars := make([]rune, 0)
Check:
	for _, line := range p.lines {
		stack := make([]rune, 0)
		for _, input := range line {
			for opener := range openCloseMap {
				if input == opener {
					stack = append(stack, input)
				}
			}
			for opener, closer := range openCloseMap {
				if input == closer {
					popped := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					if popped != opener {
						illegalChars = append(illegalChars, input)
						continue Check
					}
				}
			}
		}
		p.incompleteStacks = append(p.incompleteStacks, stack)
	}
	var sum uint64 = 0
	for _, chr := range illegalChars {
		switch chr {
		case ')':
			sum += 3
		case ']':
			sum += 57
		case '}':
			sum += 1197
		case '>':
			sum += 25137
		}
	}
	p.solution1 = (sum)
}

func (p *Day10) Run2() {
	scores := make([]int, 0)

	for _, stack := range p.incompleteStacks {

		var sum int = 0
		for i := len(stack) - 1; i >= 0; i-- {
			sum *= 5
			opener := stack[i]
			switch opener {
			case '(':
				sum += 1
			case '[':
				sum += 2
			case '{':
				sum += 3
			case '<':
				sum += 4
			}
		}
		scores = append(scores, sum)
	}
	sort.Ints(scores)
	p.solution2 = uint64(scores[len(scores)/2])
}

func (p *Day10) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day10) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
