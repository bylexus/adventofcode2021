package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 08 - Seven Segment Search
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
)

type Day08DigitMap map[string]int64
type Day08 struct {
	input     [][]string
	solution1 uint64
	solution2 uint64
}

func (p *Day08) GetName() string {
	return "AoC 2021 - Day 8 - Seven Segment Search"
}

func (p *Day08) Init() {
	// Read input
	p.input = make([][]string, 0)

	// lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day08-test.txt"), `\s+`)
	lines := lib.SplitLinesByRegex(lib.ReadInputLines("input/day08-input.txt"), `\s+`)
	for _, line := range lines {
		res := line[:10]
		res = append(res, line[11:]...)
		p.input = append(p.input, res)
	}
}

func mapKnownDigits(digitMap Day08DigitMap, inputs []string) {
	for i := 0; i < 10; i++ {
		digit := inputs[i]
		switch len(digit) {
		// 2-segment: 1
		case 2:
			digitMap[digit] = 1
		// 3-segment: 7
		case 3:
			digitMap[digit] = 7
		// 4-segment: 4
		case 4:
			digitMap[digit] = 4
		// 7-segment: 8
		case 7:
			digitMap[digit] = 8
		}
	}
}

func (p *Day08) Run1() {
	knownCounter := 0
	for _, line := range p.input {
		digitMap := make(Day08DigitMap)
		// fill in known digits:
		mapKnownDigits(digitMap, line)
		// count known digits:
		for _, output := range line[10:] {
			// just use known lengths:
			switch len(output) {
			case 2:
				knownCounter++
			case 3:
				knownCounter++
			case 4:
				knownCounter++
			case 7:
				knownCounter++
			}
		}
	}

	p.solution1 = uint64(knownCounter)
}

/**
Idea Solution 2:
I simply rely on statistics:
Each digit from 0-9 uses a certain number of single segments.
For all 10 digits, we create a segment statistic: How many times
is a single segment used for all digits?

I then create a sum for each digit:
- get the occurence number of for all segments for a digit
- sum up those usage number

--> it appears that this usage numer is UNIQUE for all of the 10 digits!
so we can simply identify the single output digit string by calculating
the statistics as described above, and done!

Example:

Digit 1 uses 2 segments (#2 and #5, see below). Those 2 segments appear as follows
in all 10 digits:

* #2: 8 times for all 10 digits
* #5: 9 times for all 10 digits
* Adds up to 8 + 9 = 17 for the digit "1"

This sum is unique for all 10 digits, and can be calculated beforehand.

Then I do the same for the 4 output digits, find the corresponding sum in the
pre-calculated stats, and got the real digit number for each digit.

As an example, the output digit string "ac" is interpreted as follows:

"a" and "c" are different segments of a 2-segment digit, but we don't know which.
But we DO know that for all input digit strings, "a" appears 37 times, while "c"
appears 34 times. The sum of these appearance per segment is 71,
so "ac" must be the digit "1" (sum matches).


For easier identification, I number each digit's segment:
     0
    dddd
1  e    a 2
   e    a
3   ffff
4  g    b 5
   g    b
    cccc
	  6
*/
func (p *Day08) Run2() {
	var totalSum int64 = 0

	// Pre-calculations:

	// create a segment map for each digit: (which digit uses which segments)
	numberSegmentMap := make([][]int, 10)
	numberSegmentMap[0] = []int{0, 1, 2, 4, 5, 6}
	numberSegmentMap[1] = []int{2, 5}
	numberSegmentMap[2] = []int{0, 2, 3, 4, 6}
	numberSegmentMap[3] = []int{0, 2, 3, 5, 6}
	numberSegmentMap[4] = []int{1, 2, 3, 5}
	numberSegmentMap[5] = []int{0, 1, 3, 5, 6}
	numberSegmentMap[6] = []int{0, 1, 3, 4, 5, 6}
	numberSegmentMap[7] = []int{0, 2, 5}
	numberSegmentMap[8] = []int{0, 1, 2, 3, 4, 5, 6}
	numberSegmentMap[9] = []int{0, 1, 2, 3, 5, 6}

	// count the segments per digit:
	// segmentNrCount map(digit => nr of segments)
	segmentNrCount := make(map[int]int)
	for _, segments := range numberSegmentMap {
		for _, s := range segments {
			segmentNrCount[s]++
		}
	}
	// calc segment appeareance for all digits
	// segmentOccurenceSumPerDigit map(digit => sum of all segment appearances)
	// this is a unique sum per digit
	segmentOccurenceSumPerDigit := make([]int, 10)
	for d, segs := range numberSegmentMap {
		sum := 0
		for _, segnr := range segs {
			sum += segmentNrCount[segnr]
		}
		segmentOccurenceSumPerDigit[d] = sum
	}

	// now, process each input
	for _, line := range p.input {

		// Step 1: count how many times a segment appears:
		segmentCount := make(map[rune]int)
		for _, digit := range line[:10] {
			for _, c := range digit {
				segmentCount[c]++
			}
		}

		// now we can simply lookup the nr of occurence sums in the real digit map for each output digit:
		var realNumber int64 = 0
		for _, outputDigit := range line[10:] {
			sum := 0
			// calc segment appearance sum from each segment in the actual digit
			for _, seqnr := range outputDigit {
				sum += segmentCount[seqnr]
			}
			// find matching real digit:
			for digit, dSum := range segmentOccurenceSumPerDigit {
				if dSum == sum {
					realNumber = (realNumber*10 + int64(digit))
					break
				}
			}
		}
		totalSum += realNumber
	}

	p.solution2 = uint64(totalSum)
}

func (p *Day08) array_contains(check int, arr []int) bool {
	for _, nr := range arr {
		if check == nr {
			return true
		}
	}
	return false
}

func (p *Day08) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day08) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
