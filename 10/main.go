package main

import "fmt"

var array = [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

func main() {

	//group := make(map[int][]int)

	for temp := range array {

		fmt.Println(temp)
	}
}
