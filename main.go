package main

import (
	"os"

	"alexi.ch/aoc2021/lib"
	"alexi.ch/aoc2021/problems"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the problem name (e.g. 'day01') as parameter.")
	}

	problemName := os.Args[1]

	problemMap := make(map[string]problems.AocProblem)
	problemMap["prepare01"] = &problems.Prepare01{}
	problemMap["prepare02"] = &problems.Prepare02{}

	problemMap["day01"] = &problems.Day01{}

	problem, defined := problemMap[problemName]
	if defined == true {
		lib.OutputTitle(problem.GetName())
		problem.Init()

		duration := lib.MeasureTime(problem.Run1)
		lib.OutputSolution(1, duration, problem.GetSolution1())

		duration = lib.MeasureTime(problem.Run2)
		lib.OutputSolution(2, duration, problem.GetSolution2())
	}
}
