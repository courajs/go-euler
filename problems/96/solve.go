package sudoku

import (
	. "fmt"
	"sort"
)

const ID = 96
const Title = "Sudoku"

func Solve() {
	unsolved := readBoards()
	solved := solveBoards(unsolved)
	for b := range solved {
		Println(b)
	}
}

type intSet map[int]bool

func (set *intSet) keys() (result []int) {
	for k := range *set {
		result = append(result, k)
	}
	return
}

func emptySet() intSet {
	return make(intSet)
}
func fullSet() intSet {
	result := make(intSet)
	for i := 1; i <= 9; i++ {
		result[i] = true
	}
	return result
}
func (s intSet) String() string {
	result := make([]int, 0, len(s))
	for k := range s {
		result = append(result, k)
	}
	sort.Ints(result)
	return Sprint(result)
}

type Solver struct {
	title string
	cells [9][9]Cell
}
type Cell struct {
	row, col, value int
	possibilities   intSet
	board           *Solver
}

func (c *Cell) solved() bool {
	return c.value != 0
}

func (c *Cell) String() string {
	return Sprintf("(%d,%d:%d:%v)", c.row, c.col, c.value, c.possibilities)
}

func (c *Cell) Row() [9]*Cell {
	result := [9]*Cell{}
	for i := range result {
		result[i] = &(c.board.cells[c.row][i])
	}
	return result
}
func (c *Cell) Col() [9]*Cell {
	result := [9]*Cell{}
	for i := range result {
		result[i] = &c.board.cells[i][c.col]
	}
	return result
}
func (c *Cell) Square() (result [9]*Cell) {
	big_row := c.row / 3
	big_col := c.col / 3
	low_row := big_row * 3
	low_col := big_col * 3
	high_row := low_row + 3
	high_col := low_col + 3

	idx := 0
	for i := low_row; i < high_row; i++ {
		for j := low_col; j < high_col; j++ {
			result[idx] = &c.board.cells[i][j]
			idx++
		}
	}
	return result
}

type posHandler func(row, col int, cell *Cell)

func (s *Solver) eachPos(f posHandler) {
	for row := range s.cells {
		for col := range s.cells[row] {
			f(row, col, &s.cells[row][col])
		}
	}
}

func (s *Solver) each(f func(*Cell)) {
	s.eachPos(func(_, _ int, cell *Cell) {
		f(cell)
	})
}

func MakeSolver(board *BoardState) *Solver {
	result := &Solver{title: board.title}
	result.eachPos(func(row, col int, cell *Cell) {
		cell.board = result
		cell.row = row
		cell.col = col
		cell.value = board.cells[row][col]
		if cell.value == 0 {
			cell.possibilities = fullSet()
		}
	})

	return result
}

func (s *Solver) ToBoard() BoardState {
	result := BoardState{title: s.title}

	s.eachPos(func(row, col int, cell *Cell) {
		result.cells[row][col] = cell.value
	})
	return result
}

func (s *Solver) solved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !s.cells[i][j].solved() {
				return false
			}
		}
	}
	return true
}

func solveBoards(in chan BoardState) chan BoardState {
	out := make(chan BoardState, 50)
	go func() {
		defer close(out)
		for b := range in {
			out <- solveBoard(b)
		}
	}()
	return out
}

func (cell *Cell) pruneNeighborPossibilities() {
	for _, neighbor := range cell.Row() {
		delete(neighbor.possibilities, cell.value)
	}
	for _, neighbor := range cell.Col() {
		delete(neighbor.possibilities, cell.value)
	}
	for _, neighbor := range cell.Square() {
		delete(neighbor.possibilities, cell.value)
	}
}

func solveBoard(in BoardState) BoardState {
	solver := MakeSolver(&in)
	// prune possibilities with all the info we have from already filled cells
	solver.each((*Cell).pruneNeighborPossibilities)

	progress := true
	for progress {
		progress = false
		solver.each(func(cell *Cell) {
			if len(cell.possibilities) == 1 {
				progress = true
				cell.value = cell.possibilities.keys()[0]
				cell.possibilities = emptySet()
				cell.pruneNeighborPossibilities()
			}
		})
	}

	if !solver.solved() {
		Println("board too hard:")
		for i := 0; i < 9; i++ {
			Println(solver.cells[i][0].Row())
		}
		panic("ahh")
	}

	return solver.ToBoard()
}
