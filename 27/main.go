package main

import (
	"fmt"
	"strings"
)

var (
	s1 = "snow dog sun"
	s2 = "снег собака солнце"
	s3 = ""
	s4 = "世界"
)

func reverseWords1(s string) string {
	slice := strings.Split(s, " ")
	if len(slice) <= 1 {
		return s
	} else {
		var res string
		for i := len(slice) - 1; i > 0; i-- {
			res = res + slice[i] + " "
		}
		return res + slice[0]
	}
}

func main() {
	fmt.Println("Started string1 = ", s1)

	rev1 := reverseWords1(s1)

	fmt.Println("Finished string1 = ", rev1)

	fmt.Println("Started string2 = ", s2)

	rev2 := reverseWords1(s2)

	fmt.Println("Finished string1 = ", rev2)

	fmt.Println("Started string2 = ", s3)

	rev3 := reverseWords1(s3)

	fmt.Println("Finished string1 = ", rev3)

	fmt.Println("Started string2 = ", s4)

	rev4 := reverseWords1(s4)

	fmt.Println("Finished string1 = ", rev4)
}
