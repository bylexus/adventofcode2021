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
* Ugly: Slices.... Oh boy, if you don't be very careful, you will mess up...
  Slices are somewhat.... Scary... As soon as you re-assign sub-slices, bad things happen.
  An example:
```go
a := []int{1,2,3,4,5,6}
b := []int{7,8,9,10}
c := a[0:3]
c = assing(c, b[2:])
a[0] = 42 // boom! now c[0] is ALSO 42... 
```
  You should copy a slice first, to avoid such behaviour. So you have to especially careful if working with slices. And it is NOT AT ALL 
  clear that things like the one shown above happen...

* Ugh: No ternary operator (e.g. `a := foo == true ? "yes" : "no"`). So simple one-liners like the one shown are just not possible,
  they can only be solved like this:
```go
var a string
if foo == true {
  a = "yes"
} else {
  a = "no"
}
```

I mean, really??? sorry, that's just insane...


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

## Run times of all solutions

```plain
Nr.   |Title                              |    Time Part 1|    Time Part 2|     Total time
------|-----------------------------------|---------------|---------------|---------------
day01 |Sonar Sweep                        |        2.481??s|        3.006??s|        5.487??s
day02 |Dive!                              |        7.182??s|         7.32??s|       14.502??s
day03 |Binary Diagnostic                  |       81.136??s|     2.821329ms|     2.902465ms
day04 |Giant Squid                        |      632.188??s|     1.713833ms|     2.346021ms
day05 |Hydrothermal Venture               |     18.43038ms|     26.95379ms|     45.38417ms
day06 |Lanternfish                        |       13.971??s|        6.776??s|       20.747??s
day07 |The Treachery of Whales            |     4.172039ms|     6.309854ms|    10.481893ms
day08 |Seven Segment Search               |       48.125??s|      446.089??s|      494.214??s
day09 |Smoke Basin                        |      246.672??s|     4.419703ms|     4.666375ms
day10 |Syntax Scoring                     |     5.066944ms|       18.782??s|     5.085726ms
day11 |Dumbo Octopus                      |     6.596268ms|    13.213615ms|    19.809883ms
day12 |Passage Pathing                    |    29.483121ms|    1.85672923s|   1.886212351s
day13 |Transparent Origami                |       127.63??s|      758.774??s|      886.404??s
day14 |Extended Polymerization            |     3.862528ms|     4.864432ms|      8.72696ms
day15 |Chiton                             |    44.324491ms|   3.469638937s|   3.513963428s
day16 |Packet Decoder                     |      136.618??s|        3.655??s|      140.273??s
day17 |Trick Shot                         |     13.98944ms|          200ns|     13.98964ms
day18 |Snailfish                          |    15.420243ms|   134.653367ms|    150.07361ms
------------------------------------------|---------------|---------------|---------------
Total time                                |   142.641457ms|   5.522562692s|   5.665204149s

Total run time (parallel runs): 3.53026127s
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

That one was a bit tricky - The na??ve solution, which worked fine for part one, just exploded into my face
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

## Day 17 - Trick Shot

First, I tried to calculate things mathematically - which I'm sure is possible, if you are smart enough - I am not!

So I did it the brute-force way: just tried different start velocities for x/y direction, and simulated the 
shot's steps.

For the start y velocity, I just implied that this will not need to be larger than 500.... And it worked.

Run times:

* Solution 1: 5.7ms (no optimizations)
* Solution 2: 0 (could be done in run 1 already)

## Day 18 - Snailfish

Phew.... That was the hardest one so far.... It seemed like a relatively simple
tree traversal problem, but boy, was I wrong :-))

The biggest problem here was: The structure itself is a tree, but you have to 
treat it as flat list for adapting the reductions.

So I lost a great amount of time to implement the reduce() method.

In the end I settled to a relatively simple solution:

1. for reducing the tree, I flattened it to a list
2. then I startet looking for rules that apply from the beginning
3. as soon as a rule matched, I executed it and start over again
4. until nothing was to be done in a loop.

My solution is somewhat ineffective, as I have to create / copy a lot of arrays...

So my run times exploded:

* Solution 1: 140ms
* Solution 1: 2051ms :-(

**Refactoring:**

It nagged me that I had to copy the tree to a flat array to find left/right neighbours... This MUST be possible
by just walking the tree!

So I invested some minutes to refactor it, and voil??, that's it: Instead of flattening the tree to an array,
I simply walk it to:
1) find the element left/right of the actual one, to execute the explosion
2) find the next element to reduce

This is MUCH better for memory AND runtime:

* Solution 1: 6.7ms
* Solution 1: 104ms

## Day 19 - Beacon Scanner

Sorry - this is my limit - This one is way off-limit for my brain or at least I am not
ready to invest that many time. I skip that day.

## Day 20 - Trench Map

At first that one seemed relatively complex - because all 9x9 empty patches would toggle on/off for each round,
as my algo input suggested. This is unfortunate, as "empty" space don't keep empty - it toggles.

After a bit of thinking I found a very simple solution: I can just remember what "unknown" empty 9x9 patches are at the moment:
on or off. Then I just had to calculate all pixels that are NOT fully empty / full. This includes one more pixel row/col for
each round, so each round the pixels to calculate increases (a bit).

So in the end, simple if you see the clue behind it :-)

Run times:

* Solution 1: 17ms
* Solution 2: 931ms

**Refactoring:**

I felt the urge to make it a bit more efficient - so instead of operating with strings / runes, I use a byte as pixel - 0 or 1.
This makes the algorithm slightly more efficient:

* Solution 1: 9ms
* Solution 2: 528ms

## Day 21 - Dirac Dice

OK, that's a tough one - I solved part 1 naively. For the 2nd part I have to figure out some kind of memoization.

My ideas so far:

- keep / cache the number of future wins (future = all possible permutations for future rolls) for each player for each actual point
  (e.g. if I reach 7 points, and I already reached 7 points with another combination before, I know already all the future wins).

## Day 22 - Reactor Reboot

1st part solved with single pixels - which defintifely will not work for the 2nd part. I *think* I let that one go - 
this is far too complex for me to implement.

## Day 23 - Amphipod

Some kind of "Tower of hanoi" problem? No idea how to solve that - I guess this is my mental limit :-)

## Day 24 - Arithmetic Logic Unit

A virtual CPU, allright - but I guess just try-and-error will not work here. So I guess I have to analyze the instructions
and see what happens.

Or maybe I just try the brute-force method for part one? Let's see...


## Day25 - Sea Cucumbers

The last one was relatively simple - but unfortunately, I missed some stars in the last days...
So no final sequence for me :-)

Maybe there is a way to make today's solution more performant - I used a double-buffered map
(so instantiated a new map TWICE every turn). I think about that a bit, to make it more
performant...