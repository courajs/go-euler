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
  data_path := data_path("96-sudoku.txt")
  fmt.Println(data_path)
  f, err := os.Open(data_path)
  if err != nil {
    fmt.Println("Couldn't read data file!")
    return
  }

  lines := bufio.NewScanner(f)
  boards := make([]Board, 0, 50)
  for lines.Scan() {
    b := Board{title: lines.Text()}
    for row := 0; row < 9; row++ {
      lines.Scan()
      line := lines.Text()
      for col, char := range line {
        b.cells[row][col] = int(char - '0')
      }
    }
    boards = append(boards, b)
  }

  fmt.Println(boards[0])
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
