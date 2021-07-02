package main

import (
	"fmt"
	"sync"
)

var (
	m     = make(map[rune]struct{})
	slice = []string{
		"asd",
		"фыв",
		"aasd",
		"фывв",
		"世界",
		"世界界",
		" ",
		"  ",
	}
)

func checkUnique(s string, mu *sync.Mutex) bool {
	runes := []rune(s)
	mu.Lock()
	for _, rune := range runes {
		if _, ok := m[rune]; ok {
			mu.Unlock()
			return false
		} else {
			m[rune] = struct{}{}
		}
	}
	mu.Unlock()
	return true
}

func main() {
	var mu sync.Mutex
	for _, s := range slice {
		fmt.Printf("All symbols of (%v) are uniq: %v\n", s, checkUnique(s, &mu))
	}
}
