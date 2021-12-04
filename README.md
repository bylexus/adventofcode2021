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
1ns for problem 1, 2ns for problem 2, if the measurement is more or less correct.


## Day 02 - Dive!

Again, a simple loop-and-count problem, almost easier than the first. Nothing special here.
Needed it to figure out how to use regular expressions in Go, and to pack the values into a data struct,
that took me longer than the problem itself :-)

## Day 03 - Binary Diagnostic

Today bit operations where asked - as a first approach, I did it completely string-wise - which does NOT satisfy me...
So I will try it again, that time with proper binary operations (masking, and/oring etc..)

On a 2nd thought, it may be not much faster to do it bitwise: in the end, a string position access and char compare is
not so much slower than bit mask operations.

So my bot solutions took:

* Solution 1: 0.06ms
* Solution 2: 0.3ms

OK, that's fast enough :-)

## Day 04 - Giant Squid

That one was a bit of a loop orgy - loop over boards over rows over cols... so at least O(n^3), which is not that fancy,
but the input was not too much, so that's ok.

The good thing: Solution 1 could be re-used for solution 2, so part of it was already done.

Solutions took:

* Solution 1: 0.4ms
* Solution 2: 1.2ms

Pretty fast :-)