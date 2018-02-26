package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

type Cell struct {
  value int
}
type Board struct {
  title string
  cells [9][9]int
}

func (_ Euler) P96Sudoku() {
  unsolved := make(chan Board, 50)
  solved := make(chan Board, 50)
  go readBoards(unsolved)
  go solveBoards(unsolved, solved)
  b := <-solved
  fmt.Println(b)
}

func solveBoards(in, out chan Board) {
  defer close(out)
  for b := range in {
    for i, row := range b.cells {
      for j := range row {
        b.cells[i][j] = 9
      }
    }
    out <- b
  }
}

func readBoards(out chan Board) {
  defer close(out)

  data_path := data_path("96-sudoku.txt")
  f, err := os.Open(data_path)
  if err != nil {
    fmt.Println("Couldn't read data file!")
    return
  }

  lines := bufio.NewScanner(f)
  for lines.Scan() {
    b := Board{title: lines.Text()}
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


func (b Board) String() string {
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
