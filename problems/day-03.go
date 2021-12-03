package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 03 - Binary Diagnostic
// ----------------------------------------------------------------------------

import (
	"fmt"
	"math"
	"strconv"

	"alexi.ch/aoc2021/lib"
)

type Day03Input struct {
}

type Day03 struct {
	input     []string
	maxLen    int
	oneBits   []int
	zeroBits  []int
	solution1 int64
	solution2 int64
}

func (p *Day03) GetName() string {
	return "AoC 2021 - Day 3 - Binary Diagnostic"
}

func (p *Day03) Init() {
	// Read input
	// p.input = lib.ReadInputLines("input/day03-test.txt")
	p.input = lib.ReadInputLines("input/day03-input.txt")
	p.maxLen = 0
	for _, line := range p.input {

		p.maxLen = int(math.Max(float64(p.maxLen), float64(len(line))))
	}
}

//
// Count the ones and zeroes per position in a list of bit strings.
// So `ones[0]` contains the nr of 1-bits at the first (highest) position,
// `zeroes[7]` contains the nr of 0-bits at the 8th position in the list
//
// returns ones and zeroes as array
func (p *Day03) countBits(input []string) ([]int, []int) {
	ones := make([]int, p.maxLen)
	zeroes := make([]int, p.maxLen)
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			if line[i] == '0' {
				zeroes[i]++
			} else {
				ones[i]++
			}
		}
	}
	return ones, zeroes

}

func (p *Day03) Run1() {
	ones, zeroes := p.countBits(p.input)

	gamma := ""
	epsilon := ""
	for i := 0; i < len(ones); i++ {
		if ones[i] > zeroes[i] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}
	p.solution1 = gammaInt * epsilonInt
}

func (p *Day03) Run2() {
	oxygenList := p.input[:]
	co2List := p.input[:]

	// Filter the input list as long as there are > 1 element:
	// each inner loop filters the list by the actual bit position,
	// then for the next loop, the bit position is increased to the
	// right (next) bit.

	// Oxygen list: keep more significant entries:
	bitNr := 0
	for {
		newList := make([]string, 0)
		ones, zeroes := p.countBits(oxygenList)
		for _, line := range oxygenList {
			if line[bitNr] == '1' && ones[bitNr] >= zeroes[bitNr] {
				newList = append(newList, line)
			}
			if line[bitNr] == '0' && zeroes[bitNr] > ones[bitNr] {
				newList = append(newList, line)
			}
		}
		oxygenList = newList
		bitNr += 1
		if len(oxygenList) == 1 {
			break
		}
	}

	// CO2 list: keep less significant entries:
	bitNr = 0
	for {
		newList := make([]string, 0)
		ones, zeroes := p.countBits(co2List)
		for _, line := range co2List {
			if line[bitNr] == '0' && zeroes[bitNr] <= ones[bitNr] {
				newList = append(newList, line)
			}
			if line[bitNr] == '1' && ones[bitNr] < zeroes[bitNr] {
				newList = append(newList, line)
			}
		}
		co2List = newList
		bitNr += 1
		if len(co2List) == 1 {
			break
		}
	}

	oxyVal, err := strconv.ParseInt(oxygenList[0], 2, 64)
	if err != nil {
		panic(err)
	}
	coVal, err := strconv.ParseInt(co2List[0], 2, 64)
	if err != nil {
		panic(err)
	}

	p.solution2 = oxyVal * coVal
}

func (p *Day03) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day03) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
