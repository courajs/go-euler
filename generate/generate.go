package main

import (
  "text/template"
  "os"
)

const format = `package main
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
}`


func main() {
  dir,_ := os.Open("problems");
  problems,_ := dir.Readdirnames(0)

  out,_ := os.Create("solve.go")

  tmpl := template.Must(template.New("solve").Parse(format))
  tmpl.Execute(out, problems)
}
