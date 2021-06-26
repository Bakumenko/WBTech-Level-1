package main

import (
	"sort"
)

var (
	array12 = []int{2, 5, 3, 1, 4}
	array22 = []int{5, 7, 6, 4, 8, 3}
)

func main() {
	sort.Ints(array12)
	sort.Ints(array22)

}
