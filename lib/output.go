package lib

import (
	"fmt"
	"sort"
	"time"

	"alexi.ch/aoc2021/lib/types"
)

func AnsiBold(input string) string {
	return fmt.Sprintf("\033[31;1;4m%v\033[0m", input)
}

func Highlight(input string) string {
	return fmt.Sprintf("\033[1;4m%v\033[0m", input)
}

func OutputTitle(input string) {
	outStr := fmt.Sprintf("%v", input)
	fmt.Printf("\033[1;4m%v\033[0m\n\n", outStr)
}

func OutputSolution(nr int, duration time.Duration, solution string) {
	fmt.Printf("Solution %v: %v\n", nr, Highlight(solution))
	fmt.Printf("Took: %v ms\n\n", float64(duration.Nanoseconds())/1000.0/1000.0)
}

func MeasureTime(f func()) time.Duration {
	start := time.Now()
	f()
	end := time.Since(start)
	return end
}

func OutputResultList(list []types.AoCResult, total time.Duration) {
	fmt.Printf(
		"%-6s|%-35s|%15s|%15s|%15s\n",
		"Nr.",
		"Title",
		"Time Part 1",
		"Time Part 2",
		"Total time",
	)
	fmt.Print(
		"------|-----------------------------------|---------------|---------------|---------------\n",
	)
	var total1 time.Duration = 0
	var total2 time.Duration = 0
	var runTotal time.Duration = 0

	sort.Slice(list, func(i, j int) bool {
		return list[i].Key < list[j].Key
	})

	for _, line := range list {
		total1 += line.TimeSolution1
		total2 += line.TimeSolution2
		runTotal += line.TimeSolution1 + line.TimeSolution2
		name := line.Problem.GetName()
		if len(name) > 29 {
			name = name[:len(name)-6] + "..."
		}
		fmt.Printf(
			"%-6s|%-35s|%15s|%15s|%15s\n",
			line.Key,
			name,
			line.TimeSolution1,
			line.TimeSolution2,
			line.TimeSolution1+line.TimeSolution2,
		)
	}
	fmt.Print(
		"------------------------------------------|---------------|---------------|---------------\n",
	)
	fmt.Printf(
		"%-42s|%15s|%15s|%15s\n",
		"Total time",
		total1,
		total2,
		runTotal,
	)
	fmt.Printf("\nTotal run time (parallel runs): %s\n\n", total)
}
