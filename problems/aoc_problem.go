package problems

type AocProblem interface {
	Init()                // Called by main before Run1/2. Here you can read the input data.
	Run1()                // calcs the solution for problem 1
	Run2()                // calcs the solution for problem 2
	GetSolution1() string // returns a string that represents solution 1
	GetSolution2() string // returns a string that represents solution 2
	GetName() string      // a title of the problem
}
