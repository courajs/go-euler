package sudoku

import (
	. "fmt"
)

const ID = 96
const Title = "Sudoku"

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
	solver.each((*Cell).pruneNeighborPossibilities)

	progress := true
	for progress {
		progress = false
		solver.each(func(cell *Cell) {
			if len(cell.possibilities) == 1 {
				progress = true
				cell.value = cell.possibilities.Keys()[0]
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

// individual Cells
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

// accessors for various neighbor sets of a cell
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

// Overall Solver struct
type Solver struct {
	title string
	cells [9][9]Cell
}

func (s *Solver) each(f func(*Cell)) {
	for row := range s.cells {
		for col := range s.cells[row] {
			f(&s.cells[row][col])
		}
	}
}

// Cells of a board have a pointer back out to the board,
// so that a cell can find its own neighbors
func MakeSolver(board *BoardState) *Solver {
	result := &Solver{title: board.title}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			val := board.cells[row][col]
			result.cells[row][col] = Cell{
				row:   row,
				col:   col,
				value: val,
				board: result,
			}
			if val == 0 {
				result.cells[row][col].possibilities = fullSet()
			} else {
				result.cells[row][col].possibilities = emptySet()
			}
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

func (cell *Cell) pruneNeighborPossibilities() {
	for _, neighbor := range cell.Row() {
		neighbor.possibilities.Delete(cell.value)
	}
	for _, neighbor := range cell.Col() {
		neighbor.possibilities.Delete(cell.value)
	}
	for _, neighbor := range cell.Square() {
		neighbor.possibilities.Delete(cell.value)
	}
}
