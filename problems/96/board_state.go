package sudoku

import (
	"bufio"
	. "fmt"
	"log"
	"strings"

	"github.com/courajs/go-euler/util"
)

type BoardState struct {
	title string
	cells [9][9]int
}

func (b BoardState) String() string {
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

func readBoards() chan BoardState {
	out := make(chan BoardState)
	go func() {
		defer close(out)

		f, err := util.DataFile("sudoku.txt")
		if err != nil {
			Println("Couldn't read data file sudoku.txt!")
			return
		}

		lines := bufio.NewScanner(f)
		for lines.Scan() {
			b := BoardState{title: lines.Text()}
			for row := 0; row < 9; row++ {
				lines.Scan()
				line := lines.Text()
				for col, char := range line {
					b.cells[row][col] = int(char - '0')
				}
			}
			out <- b
		}
		if lines.Err() != nil {
			log.Fatal("Error reading sudoku boards:", lines.Err())
		}
	}()
	return out
}
