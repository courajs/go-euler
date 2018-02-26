package main

import (
  "os"
  "fmt"
  "reflect"
  "regexp"
  "strings"
  "path/filepath"
)

func data_path(filename string) string {
  go_path := os.Getenv("GOPATH")
  return filepath.Join(go_path, "src", "github.com", "courajs", "go-euler", "data", filename)
}

type Euler struct{}

type ident struct {
  number, name string
}

var identParser *regexp.Regexp = regexp.MustCompile(`P(\d+)(\w+)`)
func identsFor(n string) ident {
  matches := identParser.FindStringSubmatch(n)
  return ident{matches[1], strings.ToLower(matches[2])}
}

func main() {
  ps := reflect.ValueOf(Euler{})
  PS := ps.Type()
  num_methods := PS.NumMethod()

  by_num  := make(map[string]reflect.Value)
  by_name := make(map[string]reflect.Value)
  ids := make([]ident, num_methods)

  for i := 0; i < num_methods; i++ {
    method_name := PS.Method(i).Name
    method_func := ps.Method(i)

    id := identsFor(method_name)

    ids[i] = id
    by_num[id.number] = method_func
    by_name[id.name] = method_func
  }
  if len(os.Args) > 1 {
    arg := os.Args[1]
    if f, ok := by_num[arg]; ok {
      f.Call(nil)
    } else if f, ok := by_name[arg]; ok {
      f.Call(nil)
    } else {
      printIds(ids)
    }
  } else {
    printIds(ids)
  }
}

func printIds(ids []ident) {
  for _, id := range ids {
    fmt.Printf("%s: %s\n", id.number, id.name)
  }
}


