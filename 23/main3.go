package main

import (
	"fmt"
	"sort"
)

var (
	slice3          = []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	searchedNumber3 = 5
)

func main() {
	fmt.Println("Slice = ", slice3)
	i := sort.Search(len(slice3), func(i int) bool { return slice3[i] >= searchedNumber3 })
	if i < len(slice3) && slice3[i] == searchedNumber3 {
		fmt.Printf("%v is index of %v", searchedNumber3, i)
	} else {
		fmt.Printf("%d not found in %v\n", searchedNumber3, slice3)
	}
}
