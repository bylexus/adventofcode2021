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

## Day 08 -	Seven Segment Search

This was the first day that took me quite a while to get a proper solution.

Part 1 was straight forward, just count the appeareances of some randomly arranged strings.

Part 2 was hard - first I tried to incrementally find out which character belongs to what segment of the
7-digit number - but that was too hard and error-prone: I didn't get a solution after a kilometer of code and debug statements...

Then I tried another approach:

I simply relied on statistics:

Each digit is made of 7 segments: 
For easier identification, I number each digit's segment:
```
     0
    dddd
1  e    a 2
   e    a
3   ffff
4  g    b 5
   g    b
    cccc
	  6

```

But all segments are used a different number of times over all 10 digits:
e.g. the segment 0 is used in 8 digits, while the segment 1 is only used in 6 digits.

Each digit from 0-9 uses a certain number of single segments.
For all 10 digits, we create a segment statistic: How many times
is a single segment used for all digits?

I then create a sum for each digit:
- get the occurence number of for all segments for a digit
- sum up those usage number

--> it appears that this usage number is UNIQUE for all of the 10 digits!
so we can simply identify the single output digit string by calculating
the statistics as described above, and done!

Example:

Digit 1 uses 2 segments (#2 and #5, see above). Those 2 segments appear as follows
in all 10 digits:

* #2: 8 times for all 10 digits
* #5: 9 times for all 10 digits
* Adds up to 8 + 9 = 17 for the digit "1"

This sum is unique for all 10 digits, and can be calculated beforehand.

Then I do the same for the 4 output digits, find the corresponding sum in the
pre-calculated stats, and got the real digit number for each digit.

As an example, the output digit string "ac" is interpreted as follows:

"a" and "c" are different segments of a 2-segment digit, but we don't know which.
But we DO know that for all input digit strings, "a" appears 9 times, while "c"
appears 8 times. The sum of these appearance per whole digit is 17,
so "ac" must be the digit "1" (sum matches).


--> then it was a simple lookup for the correct digit, and sum up, done.

* Solution 1 took: 0.04ms
* Solution 2 took: 0.32ms

## Day 09 - Smoke Basin

The first part - finding the lowest locations, was super-simple: just process each location, and check if
the surrounding locations are higher, done.

For the 2nd part, it was obvious to me to take some kind of recursive fill algorithm - and I ended by
"flood-filling" the pools from below: Starting with each basin bottom (low points), I flooded the basin
"upwards" in all direction, until the top was reached.
This is done with a simple recursive depth-first walk approach.

Finally, the solutions took:

* Solution 1: 0.13ms
* Solution 2: 2.1ms