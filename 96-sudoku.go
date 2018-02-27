package main

import (
  . "fmt"
  "os"
  "bufio"
  "strings"
  "sort"
)

type IntSet map[int]bool

func FullSet() IntSet {
  result := make(IntSet)
  for i:=1;i<=9;i++ {
    result[i]=true
  }
  return result
}
func (s IntSet) String() string {
  result := make([]int, 0, len(s))
  for k := range s {
    result = append(result, k)
  }
  sort.Ints(result)
  return Sprint(result)
}

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
  possibilities IntSet
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
  // apply info we have from already filled cells
  solver.each(func(cell *Cell) {
    for _,neighbor:=range cell.Row() { delete(neighbor.possibilities,cell.value) }
    for _,neighbor:=range cell.Col() { delete(neighbor.possibilities,cell.value) }
    for _,neighbor:=range cell.Square() { delete(neighbor.possibilities,cell.value) }
  })
  for i:=0;i<9;i++ {
    Println(solver.cells[i][0].Row())
  }
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

  idx := 0
  for i:=low_row; i < high_row; i++ {
    for j:=low_col; j < high_col; j++ {
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
  s.eachPos(func(_,_ int, cell *Cell) {
    f(cell)
  })
}

func MakeSolver(board *StaticBoard) Solver {
  result := Solver{title: board.title}
  result.eachPos(func(row, col int, cell *Cell) {
    cell.board = &result
    cell.row = row
    cell.col = col
    cell.value = board.cells[row][col]
    if cell.value == 0 {
      cell.possibilities = FullSet()
    }
  })

  return result
}

func (s *Solver) ToBoard() StaticBoard {
  result := StaticBoard{title: s.title}

  s.eachPos(func(row, col int, cell *Cell) {
    result.cells[row][col] = cell.value
  })
  return result
}



func (_ Euler) P96This() {
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
