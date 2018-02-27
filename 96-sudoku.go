package main

import (
  . "fmt"
  "os"
  "bufio"
  "strings"
)

type StaticBoard struct {
  title string
  cells [9][9]int
}
type Solver struct {
  title string
  cells [9][9]Cell
}
type Cell struct {
  row, col, value int
  possibilities []int
  board *Solver
}

func (c *Cell) solved() bool {
  return c.value != 0
}

func (c *Cell) String() string {
  if c.solved() {
    return Sprintf("(%d,%d:%d)", c.row, c.col, c.value)
  } else {
    return Sprintf("(%d,%d:%v)", c.row, c.col, c.possibilities)
  }
}

func solveBoard(in StaticBoard) StaticBoard {
  solver:= MakeSolver(&in)
  Println(solver.cells[0][0].Square())
  return solver.ToBoard()
}

func (c *Cell) Row() [9]*Cell {
  result := [9]*Cell{}
  for i := range result {
    result[i] = &c.board.cells[c.row][i]
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
func (c *Cell) Square() [9]*Cell {
  result := [9]*Cell{}

  low_row := c.row / 3
  low_col := c.col / 3
  high_row := low_row + 3
  high_col := low_col + 3

  for i:=low_row; i < high_row; i++ {
    for j:=low_col; j < high_col; j++ {
      idx := i*3 + j
      result[idx] = &c.board.cells[i][j]
    }
  }
  return result
}

type cellHandler func(row, col int, cell *Cell)

func (s *Solver) eachCell(f cellHandler) {
  for row := range s.cells {
    for col := range s.cells[row] {
      f(row, col, &s.cells[row][col])
    }
  }
}

func MakeSolver(board *StaticBoard) Solver {
  result := Solver{title: board.title}
  result.eachCell(func(row, col int, cell *Cell) {
    cell.board = &result
    cell.row = row
    cell.col = col
    cell.value = board.cells[row][col]
    if cell.value == 0 {
      cell.possibilities = []int{1,2,3,4,5,6,7,8,9}
    }
  })

  return result
}

func (s *Solver) ToBoard() StaticBoard {
  result := StaticBoard{title: s.title}

  s.eachCell(func(row, col int, cell *Cell) {
    result.cells[row][col] = cell.value
  })
  return result
}



func (_ Euler) P96Sudoku() {
  unsolved := make(chan StaticBoard) //, 50)
  // solved := make(chan StaticBoard, 50)
  go readBoards(unsolved)
  // go solveBoards(unsolved, solved)
  b := solveBoard(<-unsolved)
  Println(b)
}

func solveBoards(in, out chan StaticBoard) {
  defer close(out)
  for b := range in {
    out <- solveBoard(b)
  }
}


func readBoards(out chan StaticBoard) {
  defer close(out)

  data_path := data_path("96-sudoku.txt")
  f, err := os.Open(data_path)
  if err != nil {
    Println("Couldn't read data file!")
    return
  }

  lines := bufio.NewScanner(f)
  for lines.Scan() {
    b := StaticBoard{title: lines.Text()}
    for row := 0; row < 9; row++ {
      lines.Scan()
      line := lines.Text()
      for col, char := range line {
        b.cells[row][col] = int(char - '0')
      }
    }
    out <- b
  }
}


func (b StaticBoard) String() string {
  var result strings.Builder
  result.WriteString(b.title)
  result.WriteRune('\n')
  for _, row := range b.cells {
    for _, char := range row {
      result.WriteRune(rune(char + '0'))
    }
    result.WriteRune('\n')
  }
  return result.String()
}
