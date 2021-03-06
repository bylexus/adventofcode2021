package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 18 - Snailfish
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2021/lib"
)

type Day18Node interface {
	isValue() bool
	parent() *Day18Pair
	root() *Day18Pair
	setParentNode(p *Day18Pair)
	String() string
	reduceOne() bool
	magnitude() uint64
}

// ---------------------------- Pair -------------------------------------
// a pair contains 2 values (left, right), which can be either
// other pairs, or a final value node
type Day18Pair struct {
	parentNode *Day18Pair
	leftNode   Day18Node
	rightNode  Day18Node
}

func (n *Day18Pair) isValue() bool { return false }
func (n *Day18Pair) parent() *Day18Pair {
	return n.parentNode
}
func (n *Day18Pair) setParentNode(p *Day18Pair) {
	n.parentNode = p
}
func (n *Day18Pair) String() string {
	return fmt.Sprintf("[%v,%v]", n.leftNode, n.rightNode)
}

func (n *Day18Pair) nrOfParents() int {
	if n.parentNode != nil {
		return n.parentNode.nrOfParents() + 1
	} else {
		return 0
	}
}
func (n *Day18Pair) root() *Day18Pair {
	if n.parentNode == nil {
		return n
	} else {
		return n.parent().root()
	}
}

// explodeNode expects a PairNode that exactly contains
// 2 value nodes. Other sub-nodes will lead to an error.
// it executes the "explosion" by adding the left value to
// the nearest left value, adding the right value to the
// nearest right value, and replaces the actual node with a new
// 0-value node.
// ATTENTION: the replacement in the parent needs to be done by the caller!
func (n *Day18Pair) explodeNode(node *Day18Pair) *Day18Value {
	//
	leftVal := node.leftNode.(*Day18Value)
	leftAddTo := node.findLeftOf(leftVal)

	rightVal := node.rightNode.(*Day18Value)
	rightAddTo := node.findRightOf(rightVal)

	// add left side
	if leftAddTo != nil {
		leftAddTo.value += leftVal.value
	}
	// add right side
	if rightAddTo != nil {
		rightAddTo.value += rightVal.value
	}
	// return new node:
	// attention: replacement must be done by the caller:
	newNode := Day18Value{value: 0, parentNode: node.parentNode}
	return &newNode
}

// splits the actual value node into a pair node,
// and returns the new node. The Caller then must replace it
// in the tree:
func (p *Day18Pair) splitValueNode(valueNode *Day18Value) *Day18Pair {
	newNode := Day18Pair{parentNode: valueNode.parentNode}
	leftNode := Day18Value{
		parentNode: &newNode,
		value:      valueNode.value / 2,
	}
	rightNode := Day18Value{
		parentNode: &newNode,
		value:      valueNode.value - leftNode.value,
	}
	newNode.leftNode = &leftNode
	newNode.rightNode = &rightNode
	return &newNode
}

// Walk the tree from the root to find the first
// item that needs to be exploded.
// execute it, then return true (else false)
func (n *Day18Pair) reduceExplode(root Day18Node) bool {
	switch root.(type) {
	case *Day18Pair:
		e := root.(*Day18Pair)
		if e.nrOfParents() >= 4 {
			newNode := n.explodeNode(e)
			if e.parentNode.leftNode == root {
				e.parentNode.leftNode = newNode
			}
			if e.parentNode.rightNode == root {
				e.parentNode.rightNode = newNode
			}
			return true
		} else {
			ret := n.reduceExplode(e.leftNode)
			if ret == true {
				return true
			} else {
				return n.reduceExplode(e.rightNode)
			}
		}
	}
	return false
}

func (n *Day18Pair) reduceSplit(root Day18Node) bool {
	switch root.(type) {
	case *Day18Pair:
		e := root.(*Day18Pair)
		ret := n.reduceSplit(e.leftNode)
		if ret == true {
			return true
		} else {
			return n.reduceSplit(e.rightNode)
		}

	case *Day18Value:
		c := root.(*Day18Value)
		if c.value >= 10 {
			newNode := n.splitValueNode(c)
			if c.parentNode.leftNode == root {
				c.parentNode.leftNode = newNode
			}
			if c.parentNode.rightNode == root {
				c.parentNode.rightNode = newNode
			}
			return true
		}
	}
	return false
}

// execute ONE reduce operation, and return, indicating if
// any action was taken.
// repeat from outside if needed, or use reduce() method on a pair
// to reduce until done.
func (n *Day18Pair) reduceOne() bool {
	// process explodes:
	switch n.reduceExplode(n.root()) {
	case true:
		return true
	case false:
		return n.reduceSplit(n.root())
	}

	//process splits, if no explodes happened:
	return n.reduceSplit(n.root())
}

func (n *Day18Pair) walkLeftToRightToStopnode(node *Day18Pair, stopValue *Day18Value, lastValue *Day18Value) (bool, *Day18Value) {
	var found bool = false
	switch node.leftNode.(type) {
	case *Day18Pair:
		n := node.leftNode.(*Day18Pair)
		found, lastValue = n.walkLeftToRightToStopnode(n, stopValue, lastValue)
		if found == true {
			return found, lastValue
		}
	case *Day18Value:
		n := node.leftNode.(*Day18Value)
		if n == stopValue {
			return true, lastValue
		} else {
			lastValue = n
		}
	}

	switch node.rightNode.(type) {
	case *Day18Pair:
		n := node.rightNode.(*Day18Pair)
		found, lastValue = n.walkLeftToRightToStopnode(n, stopValue, lastValue)
		return found, lastValue
	case *Day18Value:
		n := node.rightNode.(*Day18Value)
		if n == stopValue {
			return true, lastValue
		} else {
			return false, n
		}
	}

	return false, lastValue
}

func (n *Day18Pair) findFirstNodeRighOfStartNode(node *Day18Pair, startValue *Day18Value, foundStart bool) (found bool, result *Day18Value) {
	switch node.leftNode.(type) {
	case *Day18Pair:
		n := node.leftNode.(*Day18Pair)
		foundStart, result = n.findFirstNodeRighOfStartNode(n, startValue, foundStart)
		if foundStart == true && result != nil {
			return foundStart, result
		}
	case *Day18Value:
		n := node.leftNode.(*Day18Value)
		if n == startValue {
			return true, nil
		} else if foundStart == true {
			return foundStart, n
		}
	}

	switch node.rightNode.(type) {
	case *Day18Pair:
		n := node.rightNode.(*Day18Pair)
		foundStart, result = n.findFirstNodeRighOfStartNode(n, startValue, foundStart)
		if foundStart == true && result != nil {
			return foundStart, result
		}
	case *Day18Value:
		n := node.rightNode.(*Day18Value)
		if n == startValue {
			return true, nil
		} else if foundStart {
			return foundStart, n
		}
	}

	return foundStart, nil
}

// finds nearest value "left of" the given element
func (n *Day18Pair) findLeftOf(val *Day18Value) *Day18Value {
	_, node := n.walkLeftToRightToStopnode(n.root(), val, nil)
	return node
}

// finds nearest value "right of" the given element
func (n *Day18Pair) findRightOf(val *Day18Value) *Day18Value {
	_, node := n.findFirstNodeRighOfStartNode(n.root(), val, false)
	return node
}

// runs reduceAll until all is done,
// so reduce completely
func (n *Day18Pair) reduce() {
	for {
		if n.reduceOne() == false {
			break
		}
	}
}

func (n *Day18Pair) magnitude() uint64 {
	return 3*n.leftNode.magnitude() + 2*n.rightNode.magnitude()
}

// ---------------------------- Pair -------------------------------------

// ---------------------------- Value -------------------------------------
type Day18Value struct {
	parentNode *Day18Pair
	value      int
}

func (n *Day18Value) isValue() bool { return true }
func (n *Day18Value) parent() *Day18Pair {
	return n.parentNode
}
func (n *Day18Value) setParentNode(p *Day18Pair) {
	n.parentNode = p
}

func (n *Day18Value) String() string {
	return fmt.Sprint(n.value)
}
func (n *Day18Value) reduceOne() bool {
	if n.value > 9 {
		return true
	}
	return false
}
func (n *Day18Value) root() *Day18Pair {
	return n.parent().root()
}

func (n *Day18Value) magnitude() uint64 {
	return uint64(n.value)
}

// ---------------------------- Value -------------------------------------

type Day18 struct {
	solution1 uint64
	solution2 uint64
	input     []string
}

func (p *Day18) GetName() string {
	return "Snailfish"
}

// pair string is always [...,...].
func (p *Day18) parsePair(input string) Day18Node {
	if len(input) == 1 {
		p := Day18Value{value: lib.ToInt(input)}
		return &p
	}

	pair := Day18Pair{}
	// remove []
	content := input[1 : len(input)-1]

	splitIndex := 0
	brackets := 0
	// left side:
	for i, c := range content {
		if c == '[' {
			brackets++
			continue
		}
		if c == ']' {
			brackets--
		}
		if brackets == 0 {
			pair.leftNode = p.parsePair(content[:i+1])
			pair.leftNode.setParentNode(&pair)
			splitIndex = i + 1
			break
		}
	}
	// now pos i must be a ',', then the rest:
	if content[splitIndex] != ',' {
		panic("Oops - ',' expected at " + string(splitIndex))
	}
	pair.rightNode = p.parsePair(content[splitIndex+1:])
	pair.rightNode.setParentNode(&pair)

	return &pair
}

func (p *Day18) Init() {
	// Read input
	// lines := lib.ReadInputLines("input/day18-test.txt")
	// lines := lib.ReadInputLines("input/day18-test-full.txt")
	p.input = lib.ReadInputLines("input/day18-input.txt")
}

func sumNodes(n1 Day18Node, n2 Day18Node) Day18Node {
	nr := Day18Pair{
		leftNode:  n1,
		rightNode: n2,
	}
	n1.setParentNode(&nr)
	n2.setParentNode(&nr)
	nr.reduce()
	return &nr
}

func (p *Day18) Run1() {
	numbers := make([]Day18Node, 0)

	// create Pair from input:
	for _, line := range p.input {
		root := p.parsePair(line).(*Day18Pair)
		numbers = append(numbers, root)
	}
	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		n1 := sum
		n2 := numbers[i]
		sum = sumNodes(n1, n2)
	}
	p.solution1 = sum.magnitude()
}

func (p *Day18) Run2() {
	lines := p.input
	var maxMagnitute uint64 = 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			n1 := p.parsePair(lines[i]).(*Day18Pair)
			n2 := p.parsePair(lines[j]).(*Day18Pair)
			sum := sumNodes(n1, n2)
			mag := sum.magnitude()
			if mag > maxMagnitute {
				maxMagnitute = mag
			}
		}
	}
	p.solution2 = maxMagnitute
}

func (p *Day18) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day18) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
