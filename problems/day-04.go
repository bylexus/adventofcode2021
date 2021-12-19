package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 04 - Giant Squid
// ----------------------------------------------------------------------------

import (
	"fmt"
	"math"
	"strconv"

	"alexi.ch/aoc2021/lib"
)

// board is a y, x 5x5 grid
type board [][]int64

func newBoard(width int, height int) board {
	board := make(board, height)
	for i := 0; i < height; i++ {
		board[i] = make([]int64, width)
	}
	return board
}

type Day04 struct {
	input     []string
	solution1 int64
	solution2 int64

	numbers          []int64
	boards           []board
	boardsWon        []bool
	lastNrIndexDrawn int
}

func (p *Day04) GetName() string {
	return "Giant Squid"
}

func (p *Day04) Init() {
	// Read input
	// p.input = lib.ReadInputLines("input/day04-test.txt")
	p.input = lib.ReadInputLines("input/day04-input.txt")
	p.boards = make([]board, 0)

	// 1st line are the numbers
	numbers := lib.SplitStringByRegex(p.input[0], ",\\s*")
	p.numbers = make([]int64, len(numbers))
	for i, n := range numbers {
		res, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			panic(err)
		}
		p.numbers[i] = res
	}

	// now parse the boards
	lineNr := 1
	boardNr := 0
	for {
		if lineNr >= len(p.input) {
			break
		}
		// Boards are 5x5 grids:
		board := newBoard(5, 5)
		for bLine := 0; bLine < 5; bLine++ {
			numbers := lib.SplitStringByRegex(p.input[lineNr], "\\s+")
			for i, n := range numbers {
				res, err := strconv.ParseInt(n, 10, 64)
				if err != nil {
					panic(err)
				}
				board[bLine][i] = res
			}
			lineNr++
		}
		p.boards = append(p.boards, board)
		boardNr++
	}
	p.boardsWon = make([]bool, len(p.boards))
}

// we mark a number as drawn by negating it (< 0 --> marked)
// 0 is marked with MinInt64, as there is no -0
func (p *Day04) markNumber(number int64, board board) {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == number {
				board[y][x] *= -1
				if number == 0 {
					// we mark the 0 with MinInt, since there is no -0 :-)
					board[y][x] = math.MinInt64
				}
			}
		}
	}
}

// Check if a board has "won" (has a full line horizontally or vertically)
func (p *Day04) checkBoard(board board) bool {
	// check horizontally:
	for y := 0; y < len(board); y++ {
		allMarked := true
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] >= 0 {
				allMarked = false
			}
		}
		if allMarked == true {
			return true
		}
	}

	// check vertically:
	for x := 0; x < len(board[0]); x++ {
		allMarked := true
		for y := 0; y < len(board); y++ {
			if board[y][x] >= 0 {
				allMarked = false
			}
		}
		if allMarked == true {
			return true
		}
	}

	return false
}

// calcs the sum of NOT marked board tiles (tiles >= 0)
func (p *Day04) sumUnmarked(board board) int64 {
	var sum int64 = 0
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] > 0 {
				sum += board[y][x]
			}
		}
	}
	return sum
}

func (p *Day04) Run1() {
	for nrIndex, draw := range p.numbers {
		// save where we were for the 2nd solution:
		p.lastNrIndexDrawn = nrIndex
		for i, board := range p.boards {
			p.markNumber(draw, board)
			// as soon as the first board has won, we have the solution:
			if p.checkBoard(board) == true {
				// mark board as won for the 2nd part of the puzzle:
				p.boardsWon[i] = true
				sum := p.sumUnmarked(board)
				p.solution1 = sum * draw
				return
			}
		}
	}
	panic("Oops, no solution! Not good!")
}

func (p *Day04) Run2() {
	// we start at the index we stopped in solution 1: so we can save some
	// re-draws:
	for nrIndex := p.lastNrIndexDrawn + 1; nrIndex < len(p.numbers); nrIndex++ {
		draw := p.numbers[nrIndex]
		for i, board := range p.boards {
			p.markNumber(draw, board)
			if p.checkBoard(board) == true {
				// if a board has won, we need to check if it is the last:
				if p.boardsWon[i] != true {
					p.boardsWon[i] = true
					nrOfBoardsWon := 0
					for _, hasWon := range p.boardsWon {
						if hasWon {
							nrOfBoardsWon++
						}
					}
					// yes, last board that won:
					if nrOfBoardsWon == len(p.boards) {
						sum := p.sumUnmarked(board)
						p.solution2 = sum * draw
						return
					}
				}
			}
		}
	}
	panic("Oops, no solution! Not good!")
}

func (p *Day04) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day04) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
