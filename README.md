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

## Day 05 - Hydrothermal Venture

The today's riddle was about drawing coordinates - as always with an indetermined size of the map I decided to store the coords
as map - coordinates as key, occurence count as value.

With that in mind it was pretty simple. For each line, I calculated the "increase" for each direction (x/y) for calculating the next coordinate,
then I simply loop until I reach the final coords. Works for both h/v and diagonal (45deg) lines.

Based on the string key generation, this solution was not that fast, so I ended at:

* Solution 1: 38ms
* Solution 2: 74ms

Maybe with a better key hash algo it would be faster?

--> OK, yes, definitifely:

I changed the coord key algorithm from string to a simple x*1000+y algo (as x/y never goes above 1000): Now the solutions are at:


* Solution 1: 10ms
* Solution 2: 22ms

## Day 06 - Lanternfish

This one took some thinking... It was clear from the first part that a brute-force solution will not work for the 2nd part - so
there need to be a more clever way to solve this... Here is my approach:

- We just count how many fishes per "age" exists --> we create an "age map": day --> nr or fishes
- each day, this list rotates and sums up:
	- fishes with count 0 got added to the count of day 6 (will be re-started)
	- rotate the list (as each fish must be decrease to 0):
		- sum of day 0 will be popped from the list --> now the list is shifted left 
		- the old day 0 sum will be added to day 6 --> they are restarted
		- at the same time, those fishes produce the same amount of new fishes, so add them to the end (day 8), too
--> just loop over all days, sum up, done :-)

With this approach, it was just a simple loop over 256 steps:

* Solution 1: 0.004ms
* Solution 2: 0.006ms

## Day 07 - The Treachery of Whales

OK, that was a very simple one - but it's a weekday, anyway :-)

Just a simple loop-and-diff riddle - for the 2nd part, it suddenly appears to me that this is a Gaussian Summary Problem -
As the total fuel ist the sum of increasing fuel per step (so for a distance of 5, you need 1+2+3+4+5 = 15 fuel),
we can simply apply Gauss: (n^2 + n) / 2 for n = distance.

Runtime:

* Solution 1: 1.8ms
* Solution 2: 2.2ms