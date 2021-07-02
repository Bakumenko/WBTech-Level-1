package main

import (
	"fmt"
	"sync"
)

var (
	oleg  = map[string]bool{}
	slice = []string{
		"aasd",
		"фыв",
		"aasd",
		"фывв",
		"世界",
		"世界界",
		" ",
		"  ",
	}
)

func checkUnique(s string, mu *sync.Mutex) {
	runes := []rune(s)
	m := make(map[rune]struct{})
	for _, rune := range runes {
		if _, ok := m[rune]; ok {
			mu.Lock()
			oleg[s] = false
			mu.Unlock()
			return
		}
	}
	mu.Lock()
	oleg[s] = true
	mu.Unlock()
}

func main() {
	var mu sync.Mutex
	for _, s := range slice {
		go checkUnique(s, &mu)
	}

	for key, value := range oleg {
		fmt.Printf("All symbols of (%v) are uniq: %v\n", key, value)
	}
}
