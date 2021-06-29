package main

import "fmt"

var s = "string with name Oleg, symbol 世界, ⌘ and русское слово"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Started string = ", s)

	rev := reverse(s)

	fmt.Println("Finished string = ", rev)
}
