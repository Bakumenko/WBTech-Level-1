package main

import "fmt"

var (
	slice3          = []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	searchedNumber3 = 8
)

func binarySearch2(slice []int, target int, start int, end int) int {
	mid := (start + end) / 2
	if start <= end {
		if target < slice[mid] {
			return binarySearch2(slice, target, start, mid-1)
		} else if slice[mid] < target {
			return binarySearch2(slice, target, mid+1, end)
		} else if slice[mid] == target {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println("Slice = ", slice3)
	index := binarySearch2(slice3, searchedNumber3, 0, len(slice3)-1)
	fmt.Printf("%v is index of %v", index, searchedNumber3)
}
