package main

import (
  "fmt"
  "os"
  p96 "github.com/courajs/go-euler/problems/96"
)

var solvers = map[string]func(){
  "96": p96.Solve,
}

func printIds() {
  fmt.Printf("%d: %s\n", p96.ID, p96.Title)
}

func main() {
  if len(os.Args) > 1 {
    arg := os.Args[1]
    if f, ok := solvers[arg]; ok {
      f()
    } else {
      printIds()
    }
  } else {
    printIds()
  }
}
