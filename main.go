package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"alexi.ch/aoc2021/lib"
	"alexi.ch/aoc2021/problems"
)

type Problems map[string]problems.AocProblem

type Durations struct {
	duration1 time.Duration
	duration2 time.Duration
}

func runProblem(problem problems.AocProblem) (duration Durations) {
	problem.Init()
	duration.duration1 = lib.MeasureTime(problem.Run1)
	duration.duration2 = lib.MeasureTime(problem.Run2)
	return
}

func runAll(probs Problems) {
	wg := sync.WaitGroup{}
	var totalRuntime1 time.Duration
	var totalRuntime2 time.Duration
	var totalRuntime time.Duration
	for _, problem := range probs {
		wg.Add(1)
		go func(p problems.AocProblem) {
			p.Init()
			d := runProblem(p)
			fmt.Printf("%v done.\n", p.GetName())
			fmt.Printf("* Solution 1: after %v: %v", d.duration1, p.GetSolution1())
			fmt.Printf("* Solution 2: after %v: %v", d.duration2, p.GetSolution2())
			totalRuntime1 += d.duration1
			totalRuntime2 += d.duration2
			totalRuntime += d.duration1 + d.duration2
			wg.Done()
		}(problem)
		fmt.Printf("%v started\n", problem.GetName())
	}
	wg.Wait()

	fmt.Printf("Total runtime Problems 1: %v\n", totalRuntime1)
	fmt.Printf("Total runtime Problems 2: %v\n", totalRuntime2)
	fmt.Printf("TOTAL runtime over all  : %v\n", totalRuntime)
}

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the problem name (e.g. 'day01') as parameter.")
	}

	problemName := os.Args[1]

	problemMap := make(Problems)
	problemMap["prepare01"] = &problems.Prepare01{}
	problemMap["prepare02"] = &problems.Prepare02{}

	problemMap["day01"] = &problems.Day01{}
	problemMap["day02"] = &problems.Day02{}
	problemMap["day03"] = &problems.Day03{}
	problemMap["day04"] = &problems.Day04{}
	problemMap["day05"] = &problems.Day05{}
	problemMap["day06"] = &problems.Day06{}
	problemMap["day07"] = &problems.Day07{}
	problemMap["day08"] = &problems.Day08{}
	problemMap["day09"] = &problems.Day09{}
	problemMap["day10"] = &problems.Day10{}
	problemMap["day11"] = &problems.Day11{}
	problemMap["day12"] = &problems.Day12{}
	problemMap["day13"] = &problems.Day13{}
	problemMap["day14"] = &problems.Day14{}
	problemMap["day15"] = &problems.Day15{}
	problemMap["day16"] = &problems.Day16{}
	problemMap["day17"] = &problems.Day17{}
	problemMap["day18"] = &problems.Day18{}

	if problemName == "all" {
		runAll(problemMap)
	} else {
		problem, defined := problemMap[problemName]
		if defined == true {
			lib.OutputTitle(problem.GetName())
			solution := runProblem(problem)
			lib.OutputSolution(1, solution.duration1, problem.GetSolution1())
			lib.OutputSolution(2, solution.duration2, problem.GetSolution2())
		} else {
			panic("Oops - Problem not found - is it defined in main?")
		}
	}
}
