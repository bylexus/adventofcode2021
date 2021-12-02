# Advent of Code 2021

This repo contains my solutions and a diary for [Advent of Code 2021](https://adventofcode.com/2021/). This year I chose GO as my language: I don't know GO, so
I take this as opportunity to learn it (a bit at least).

## Preparations

Today I set up a generic structure for all problems: Each problem will be defined as a struct that implements the following interface:

```go
type AocProblem interface {
	Init()                  // Called by main before Run1/2. Here you can read the input data.
	Run1()                  // calcs the solution for problem 1
	Run2()                  // calcs the solution for problem 2
	GetSolution1() string   // returns a string that represents solution 1
	GetSolution2() string   // returns a string that represents solution 2
	GetName() string        // a title of the problem
}
```

The main program then initializes the struct for each day, and runs `Init()`, `Run1()`, `Run2()`,
then outputs the soution including the duration.

The main program gets called with the day to run:

```shell
$ go run main.go day01
```

## Day 01 - Sonar Sweep

As always, this one is just for getting us started - I used it to fine-tune my code
structure and setting.

A simple looping problem, my machine ran both problem in practically unmeasurable time - 
1ms for problem 1, 2ms for problem 2, if the measurement is more or less correct.


## Day 02 - Dive!

Again, a simple loop-and-count problem, almost easier than the first. Nothing special here.
Needed it to figure out how to use regular expressions in Go, and to pack the values into a data struct,
that took me longer than the problem itself :-)

