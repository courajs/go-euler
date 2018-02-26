package main

import (
  "fmt"
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

func MakeSolver(board *StaticBoard) Solver {
  return Solver{title: board.title}
}

func (s *Solver) ToBoard() StaticBoard {
  return StaticBoard{title: s.title}
}



func (_ Euler) P96Sudoku() {
  unsolved := make(chan StaticBoard, 50)
  solved := make(chan StaticBoard, 50)
  go readBoards(unsolved)
  go solveBoards(unsolved, solved)
  b := <-solved
  fmt.Println(b)
}

func solveBoards(in, out chan StaticBoard) {
  defer close(out)
  for b := range in {
    solver := MakeSolver(&b)
    b = solver.ToBoard()
    out <- b
  }
}

func readBoards(out chan StaticBoard) {
  defer close(out)

  data_path := data_path("96-sudoku.txt")
  f, err := os.Open(data_path)
  if err != nil {
    fmt.Println("Couldn't read data file!")
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
