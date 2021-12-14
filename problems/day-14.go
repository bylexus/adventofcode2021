package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 14 - Extended Polymerization
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
)

type Day14Rule struct {
	input  string
	insert string
}

type Day14 struct {
	solution1 uint64
	solution2 uint64

	template string
	rules    []Day14Rule
}

func (p *Day14) GetName() string {
	return "AoC 2021 - Day 14 - Extended Polymerization"
}

func (p *Day14) Init() {
	// Read input
	// lines := lib.ReadInputLines("input/day14-test.txt")
	lines := lib.ReadInputLines("input/day14-input.txt")
	p.template = lines[0]

	steps := lib.ParseGroupMatch(lines[1:], "(.*) -> (.*)")
	p.rules = make([]Day14Rule, 0)
	for _, step := range steps {
		p.rules = append(p.rules, Day14Rule{input: step[1], insert: step[2]})
	}
}

// The naive implementation with really inserting all chars do not work,
// as this problem grows exponentially.
// So we have to be clever:
//
// we count pairs, and how many chars are inserted for each rule while processing:
//
// start: NNCB --> forms pairs NN, NC, CB, each 1 time. Also, we have a char count of N = 2, C = 1, B = 1.
//
// now we process the rules, e.g.:
// CH -> B: Nothing happens: CH does not exist as a pair yet.
// NN -> C :
//      - For each existing pair NN, C is inserted (increase the C count NN times)
//      - Additionally, All NN pairs are broken up into 2 NEW pairs NC, CN, also NN times.
// After all rules are processed in one run, add the newly added pairs to the pair map.
// ... loop
// in the end we have a map of single chars => counts, voila.
func (p *Day14) calcRounds(rounds int) uint64 {

	// initialize pairs and char count from the given template:
	pairs := make(map[string]uint64)
	charCount := make(map[int]uint64)
	for i := 0; i < len(p.template)-1; i++ {
		pair := p.template[i : i+2]
		pairs[pair]++
		charCount[int(p.template[i])]++
	}
	charCount[int(p.template[len(p.template)-1])]++

	// Do n rounds:
	for i := 0; i < rounds; i++ {
		appendPairs := make(map[string]uint64) // keeps the new pairs, to be added after the run is complete
		// process all rules:
		for _, rule := range p.rules {
			pairCount, exists := pairs[rule.input]
			if exists == true {
				// existing pair:
				// increase the char count with n times the existing pair, and form 2 new pairs
				charCount[int(rule.insert[0])] += pairCount
				// break up existing pair into 2 new ones:
				// so there are no longer n pairs, but 2*n new pairs:
				appendPairs[rule.input[0:1]+rule.insert] += pairs[rule.input]
				appendPairs[rule.insert+rule.input[1:]] += pairs[rule.input]
				pairs[rule.input] = 0
			}
		}
		// add the newly inserted pairs to the existing ones
		for pair, nr := range appendPairs {
			pairs[pair] += nr
		}
	}

	// find min/max char count and return the result:
	var min uint64 = 0
	var max uint64 = 0
	for _, count := range charCount {
		if count > max {
			max = count
		}
		if min == 0 || min > count {
			min = count
		}
	}
	return max - min

}

func (p *Day14) Run1() {
	p.solution1 = p.calcRounds(10)
}

func (p *Day14) Run2() {
	p.solution2 = p.calcRounds(40)
}

func (p *Day14) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day14) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
