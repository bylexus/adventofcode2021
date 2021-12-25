package lib

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//
// Reads a file into a slice of lines, and returns a slice of strings.
func ReadInputLines(filename string) []string {
	maxCapacity := 65536

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Cannot open file " + filename)
	}
	defer file.Close()
	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			result = append(result, line)
		}
	}

	return result
}

// Takes an slice of strings, and returns a slice of ints,
// taking each line as an int number
func ParseIntLines(lines []string) []int64 {
	res := make([]int64, len(lines))
	for i, line := range lines {
		convValue, err := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		if err != nil {
			panic(err)
		}
		res[i] = convValue
	}
	return res
}

// Takes an slice of strings, and returns a slice of uint64,
// taking each line as an unsigned int number
func ParseUIntLines(lines []string) []uint64 {
	res := make([]uint64, len(lines))
	for i, line := range lines {
		convValue, err := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		if err != nil {
			panic(err)
		}
		res[i] = uint64(convValue)
	}
	return res
}

//
// Takes a slice of strings, split each line by the given regex,
// and returns a slice of split slices (array of arrays).
func SplitLinesByRegex(lines []string, pattern string) [][]string {
	res := make([][]string, 0)
	re := regexp.MustCompile(pattern)
	for _, line := range lines {
		split := re.Split(line, -1)
		if len(split) > 0 {
			res = append(res, split)
		}
	}
	return res
}

func SplitStringByRegex(line string, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.Split(line, -1)
}

// Takes a slice of strings, and runs a regex.FindSubmatch for each line,
// returning all group matches for each line (so an array of arrays)
func ParseGroupMatch(lines []string, pattern string) [][]string {
	res := make([][]string, 0)
	re := regexp.MustCompile(pattern)
	for _, line := range lines {
		split := re.FindStringSubmatch(line)
		if len(split) > 0 {
			res = append(res, split)
		}
	}
	return res
}

func ToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func WithinInt(a, min, max int) int {
	if a < min {
		return min
	}
	if a > max {
		return max
	}
	return a
}

func MaxUInt64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func MinUInt64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func ArrayContains(check interface{}, arr []interface{}) bool {
	for _, nr := range arr {
		if check == nr {
			return true
		}
	}
	return false
}

// StackOverflow: https://stackoverflow.com/a/43004689
func AsBits(val uint64, nrOfBits int) []uint8 {
	bits := []uint8{}
	for i := 0; i < nrOfBits; i++ {
		bits = append([]uint8{uint8(val & 0x1)}, bits...)
		val = val >> 1
	}
	return bits
}

func BitsToUint64(bits []uint8) uint64 {
	var nr uint64 = 0
	for _, bit := range bits {
		nr = (nr << 1) | uint64(bit)
	}
	return nr
}
