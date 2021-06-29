package main

import "fmt"

var slice = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -15, -17, -9, 1}

func getTens(number float32) int {
	if number >= 0 {
		return int(number/10) * 10
	} //else {
	//return int(number/10 - 1) * 10
	//}
}

func main() {

	group := make(map[int][]float32)

	for _, temp := range slice {
		tens := getTens(temp)
		group[tens] = append(group[tens], temp)
	}
	fmt.Println("Result: ", group)
}
