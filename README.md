# Advent of Code 2021

This repo contains my solutions and a diary for [Advent of Code 2021](https://adventofcode.com/2021/). This year I chose GO as my language: I don't know GO, so
I take this as opportunity to learn it (a bit at least).

## About GO

I use this year's event to learn a small bit of GO. Here are my conclusions so far:

* Good: Go is super-simple: it has a very small set of keywords, and a relatively small, but powerful stdlib.
* Good: type inference: Go does infer almost any type, which is a little like writing in an untyped scripting language (which it is not!)
* Good: all-in-one compiler and package management.
* Meh: Go does not know generics. This makes it hard(er) to write generic collection functions like map(), filter() etc.
  This can be circumvented by creating collection functions for your type, but hey, I don't want to re-invent the wheel
  for each new type...
* Meh: Sometimes it is not directly clear if a function parameter is taken by value or by reference: This can lead to 
  unwanted behaviour: e.g. if you pass a Struct by value, it is copied into the function. This MAY be an advantage, but
  for me, it is annoying most of the time (as the struct will need to be modified within the function).

Around Day 15 I noticed that I got more proficient with GO. I know how to best build
data structures, I know how call by value / by reference works etc.
It became clear to me that the key to GO lies in the "implicit" types:
So if it walks like a duck, and talks like a duck, it must be a duck. That makes
things easier.

What I find hard to unserstand or to handle is to find out IF it is a duck: because
there is no explicit interface declaration for a certain type (it is not neccessary
to declare "Type A implements Interface B", just implement the methods), it's hard
to find out if a certain type can be used as "Duck"... I like explicity in programming languages...

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

## Day 10 - Syntax Scoring

This one was a stack push/pop excercise: 

For the first part, I parsed the input lines character by character:
If an opener appears, push it on to the stack. If a closer appears, pop the stack:
If the popped element is NOT the opener of the actual closer, then we have an error.

If the line could be parsed without stack errors, we get a (possibly) incomplete line.
The remaining stack is needed for the 2nd part, so keep it.

the 2nd part was very simple, as I got the correct data structure already from the 1st part:
just pop the stack for each incomplete line from behind, which contains only openers, and calc the corresponding
closer sum, done.

It was more tricky to calc the total sum correctly, in the end :-)

Run time:

* Solution 1: 1.5ms
* Solution 2: 0.01ms

 (Note to self: This problem would be a great excercise for the BZT module M411 - Algos and Datastructures)


## Day 11 - Dumbo Octopus

This one remembers me of Game of Life: a mutation problem over time.

Part 1 was straight forward: just implement the modification needed per step, and repeat the flashing until all
flashes are processed. As a data structure I chose a 2d fixed array.

I was in fear before I opened the 2nd part if the chosen simple data structure was not enough: I could imagine
the following possible variations for part 2:

- count flashes after an uinmaginally large numer of loops
- there are not only 10x10 squids, but infinite in all directions
- check when the same pattern appears again

But nothing of that sort was asked, it was much simpler. So I could go with my 2d array,
but I decided to add some "sugar" to my library, and introduded a Point structure with X/Y values, and methods to
create a hash key from coords, and coords from a hash key.

Then I used a coord map as data structure, just for the sake of it :-)

Run times:

* Solution 1: 1.8ms
* Solution 2: 6ms

## Day 12 - Passage Pathing

This one was a bit harder. At first sight, this seemed like a simple graph walk thing.
For the first part, this was correct. The only twist was the rule that large caves can be re-visited.
This made the bookkeeping of visited nodes a bit harder. I had to reset sub-parts of the graph for the next
run, which took me some time.

Then the 2nd part made it harder again, as I didn't expect re-visiting a single node 2 times...
So I introduced a little trick: Instead of adapting the algorithm, I simply added a "virtual" cave (a copy of a small cave)
to the graph, and ran it (separately for each small cave). This way, the algorithm could stay the same,
while I just had to fiddle a bit with the virtual copy.

Also, because now I had the paths through the graph for multiple runs, I had to remove double paths, which I did by create a unique set
of path hashes.

I'm not completely satisfied with the run times this time:

* Solution 1: 17ms
* Solution 2: 1.8s --> almost too much for my goal...


## Day 13 - Transparent Origami

I have the feeling that this year is definitively easier than the last years...
Today's problem reminded me of the start map problem some years (?) ago,
where you have to move stars until they form a password. But this year's problem was MUCH easier
to solve.

So almost a no-brainer, using a map for coordinates has again proven to be a much more efficient solution that an x/y array.

In a refactoring attempt, I tried if I can use the point struct itself as map key,
and yes, this works, as long as the struct only contains comparable types!

This is really helpful, as I don't have to create a key hash!

Unfortunately this took me until to day 13 to realize...

Run times:

* Solution 1: 0.06ms
* Solution 2: 0.38ms

## Day 14 - Extended Polymerization

That one was a bit tricky - The naïve solution, which worked fine for part one, just exploded into my face
in part 2 :-)

I first tried to form the whole polymer chain completely - absolutely fine for part 1. In part 2, this exploded,
as this is an exponentially growing problem. So a better solution was needed.

I had a moment until it dawns me: I can keep a "pair count", and just count the chars inserted into each pair:

Example: 
start: NNCB --> forms pairs NN, NC, CB, each 1 time. Also, we have a char count of N = 2, C = 1, B = 1.

now we process the rules, e.g.:
CH -> B: Nothing happens: CH does not exist as a pair yet, so nothing to insert.
NN -> C :
     - For each existing pair NN, C is inserted (increase the C count NN times)
     - Additionally, All NN pairs are broken up into 2 NEW pairs NC, CN, also NN times.
After all rules are processed in one run, add the newly added pairs to the pair map.
... loop

in the end I have a map of single chars => counts, voila.

Runs pretty fast :-)

* Solution 1: 0.36ms
* Solution 2: 1.45ms

## Day 15 - Chiton

This one screamed "Dijkstra!", and indeed that was the solution here.
I just implemented a straight-forward dijkstra for both solutions.

The most time I spend was to enlarge the grid in the 2nd part :-) until I noticed
that "wrapping around" the risk levels was not just a modulo, but started with 1 again.... Grrr...

Run times:

* Solution 1: 35ms
* Solution 2: 3166ms, including grid enlargement process

## Day 16 - Packet Decoder

OK, this is the parser / AST day :-) So we had to read a stream of bits, which form
single "packages" with a header and data (which can be sub-packages).

The first part was to parse the stream into packets (aka "Tokens"), and form a package tree (aka Abstract Syntax Tree).

Then the 2nd part was simply walking the AST and evaluate. Fin.

I loose much time in correctly interpreting the package structure: types and versions can be 0, too, I figured out too late :-)

All in all a nice little parsing and evaluating problem.

Run times:

* Solution 1: 0.14ms
* Solution 2: 0.0037ms
