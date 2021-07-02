package main

import "fmt"

var (
	sliceA1 = []int{1, 2, 3}
	sliceA2 = []int{4, 5}
)

func concatenationSlices(s1 []int, s2 []int) []int {
	megre := append(s1, s2...)
	return megre
}

func main() {
	megre := concatenationSlices(sliceA1, sliceA2)
	fmt.Println("Slice1: ", sliceA1)
	fmt.Println("Slice2: ", sliceA2)
	fmt.Println("Merging: ", megre)
}
