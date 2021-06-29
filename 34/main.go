package main

import "fmt"

var sls = []string{
	"asd",
	"фыв",
	"aasd",
	"фывв",
	"世界",
	"世界界",
	" ",
	"  ",
}

func checkUniq(s string) bool {
	runes := []rune(s)
	m := make(map[rune]struct{})
	for _, rune := range runes {
		if _, ok := m[rune]; ok {
			return false
		} else {
			m[rune] = struct{}{}
		}
	}
	return true
}

func main() {
	for _, s := range sls {
		fmt.Printf("All symbols of (%v) are uniq: %v\n", s, checkUniq(s))
	}
}
