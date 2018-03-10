package sudoku

import (
	"fmt"
	"sort"
)

type intSet map[int]bool

func (s *intSet) Delete(val int) {
	delete(*s, val)
}

func (s *intSet) Has(val int) bool {
	_, ok := (*s)[val]
	return ok
}

func (s *intSet) Count() int {
	return len(*s)
}

func (set *intSet) Values() (result []int) {
	for k := range *set {
		result = append(result, k)
	}
	return
}

func (s intSet) String() string {
	result := make([]int, 0, len(s))
	for k := range s {
		result = append(result, k)
	}
	sort.Ints(result)
	return fmt.Sprint(result)
}

func emptySet() intSet {
	return make(intSet)
}

func fullSet() intSet {
	result := make(intSet, 9)
	for i := 1; i <= 9; i++ {
		result[i] = true
	}
	return result
}
