package main

import (
	"fmt"
	"strings"
)

var (
	str1 = "snow dog sun bird"
	str2 = "снег собака солнце"
	str3 = ""
	str4 = "世 界"
)

func reverseWords2(s string) string {
	slice := strings.Split(s, " ")
	if len(slice) <= 1 {
		return s
	} else {
		for i := 0; i < len(slice)/2; i++ {
			slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
		}
	}
	return strings.Join(slice[:], " ")
}

func main() {
	fmt.Println("Started string1 = ", str1)

	rev1 := reverseWords2(str1)

	fmt.Println("Finished string1 = ", rev1)

	fmt.Println("Started string2 = ", str2)

	rev2 := reverseWords2(str2)

	fmt.Println("Finished string1 = ", rev2)

	fmt.Println("Started string2 = ", str3)

	rev3 := reverseWords2(str3)

	fmt.Println("Finished string1 = ", rev3)

	fmt.Println("Started string2 = ", str4)

	rev4 := reverseWords2(str4)

	fmt.Println("Finished string1 = ", rev4)
}
