package problems

type AocProblem interface {
	Init()
	Run1()
	Run2()
	GetSolution1() string
	GetSolution2() string
	GetName() string
}
