package lib

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadInputLines(filename string) []string {
	const maxCapacity = 2 ^ 16

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
		line := scanner.Text()
		result = append(result, line)
	}

	return result
}

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

func SplitLinesByRegex(lines []string, pattern string) [][]string {
	res := make([][]string, len(lines))
	re := regexp.MustCompile(pattern)
	for i, line := range lines {
		split := re.Split(line, -1)
		res[i] = split
	}
	return res
}
