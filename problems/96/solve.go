package sudoku

import (
	. "fmt"
	"strings"
)

const ID = 96
const Title = "Sudoku"

// TODO func seq(to int) []byte {}

// main
func Solve() {
	puzzles := readBoards()

	for b := range puzzles {
		Println(solveBoard(b))
	}
}

// solving algorithm
func solveBoard(in BoardState) BoardState {
	solver := MakeSolver(&in)
	// prune possibilities with all the info we have from already filled cells
	solver.each(solver.pruneNeighborPossibilities)

	progress := true
	for progress {
		progress = false
		// Commit any solved squares - ones where only one possibility remains
		solver.each(func(cell *Cell) {
			if cell.possibilities.Count() == 1 {
				progress = true
				cell.value = cell.possibilities.Values()[0]
				cell.possibilities = emptySet()
				solver.pruneNeighborPossibilities(cell)
			}
		})
		if !progress && !solver.solved() {
			// Find any squares which are the last their row, column, or section
			// which can hold some required number
		}
	}

	if !solver.solved() {
		Println("board too hard:")
		Println(solver)
		panic("ahh")
	}

	return solver.ToBoard()
}

// individual Cells
type Cell struct {
	row, col, value int
	possibilities   intSet
}

func (c *Cell) solved() bool {
	return c.value != 0
}

func (c *Cell) String() string {
	return Sprintf("(%d,%d:%d:%v)", c.row, c.col, c.value, c.possibilities)
}

// Overall Solver struct
type Solver struct {
	title string
	cells [9][9]Cell
}

func (s *Solver) String() string {
	b := strings.Builder{}
	for i := 0; i < 9; i++ {
		Fprintln(&b, s.Row(i))
	}
	return b.String()
}

func (s *Solver) each(f func(*Cell)) {
	for row := range s.cells {
		for col := range s.cells[row] {
			f(&s.cells[row][col])
		}
	}
}

// Pull in values from the board state, and
// initialize the possibility sets
func MakeSolver(board *BoardState) *Solver {
	result := &Solver{title: board.title}
	c := Cell{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			c.row = row
			c.col = col
			c.value = board.cells[row][col]
			if c.value == 0 {
				c.possibilities = fullSet()
			} else {
				c.possibilities = emptySet()
			}
			result.cells[row][col] = c
		}
	}

	return result
}

// accessors for various neighbor sets of a cell
func (s *Solver) Row(row int) (result [9]*Cell) {
	for i := range result {
		result[i] = &s.cells[row][i]
	}
	return result
}
func (s *Solver) Col(col int) (result [9]*Cell) {
	for i := range result {
		result[i] = &s.cells[i][col]
	}
	return result
}
func (s *Solver) Square(sectionNum int) (result [9]*Cell) {
	big_row := sectionNum / 3
	big_col := sectionNum % 3
	low_row := big_row * 3
	low_col := big_col * 3
	high_row := low_row + 3
	high_col := low_col + 3

	idx := 0
	for i := low_row; i < high_row; i++ {
		for j := low_col; j < high_col; j++ {
			result[idx] = &s.cells[i][j]
			idx++
		}
	}
	return result
}

func (s *Solver) ToBoard() BoardState {
	result := BoardState{title: s.title}

	s.each(func(cell *Cell) {
		result.cells[cell.row][cell.col] = cell.value
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

func (s *Solver) pruneNeighborPossibilities(c *Cell) {
	if c.value == 0 {
		return
	}
	for _, neighbor := range s.Row(c.row) {
		neighbor.possibilities.Delete(c.value)
	}
	for _, neighbor := range s.Col(c.col) {
		neighbor.possibilities.Delete(c.value)
	}
	section := c.row/3*3 + c.col/3
	for _, neighbor := range s.Square(section) {
		neighbor.possibilities.Delete(c.value)
	}
}
