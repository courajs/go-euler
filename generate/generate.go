package main

import (
  // "fmt"
  "text/template"
  "os"
)


type thing struct {
  Imports []string
}


func main() {
  dir, err := os.Open("problems")
  if err != nil { panic(err) }
  probs, err := dir.Readdir(0)
  if err != nil { panic(err) }
  nums := make([]string, len(probs))
  for i:=range nums {
    nums[i] = probs[i].Name()
  }
  // fmt.Println(nums)

  out, err := os.Create("solve.go")
  if err != nil { panic(err) }

  problems := nums
  tmpl, err := template.New("solve").Parse(
`package main
import (
  "fmt"
  "os"
{{range .}}
  p{{.}} "github.com/courajs/go-euler/problems/{{.}}"
  {{- end}}
)

var solvers = map[string]func(){
  {{- range .}}
  "{{.}}": p{{.}}.Solve,
  {{- end}}
}

func printIDs() {
  {{- range .}}
  fmt.Printf("%d: %s\n", p{{.}}.ID, p{{.}}.Title)
  {{- end}}
}

func main() {
  if len(os.Args) > 1 {
    arg := os.Args[1]
    if f, ok := solvers[arg]; ok {
      f()
    } else {
      printIDs()
    }
  } else {
    printIDs()
  }
}`)
  if err != nil { panic(err) }
  err = tmpl.Execute(out, problems)
  if err != nil { panic(err) }
}
