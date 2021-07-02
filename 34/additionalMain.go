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

func checkUnique(s string, wg *sync.WaitGroup, mu *sync.Mutex) {
	runes := []rune(s)
	m := make(map[rune]struct{})
	for _, rune := range runes {
		if _, ok := m[rune]; ok {
			mu.Lock()
			oleg[s] = false
			mu.Unlock()
			wg.Done()
			return
		} else {
			m[rune] = struct{}{}
		}
	}
	mu.Lock()
	oleg[s] = true
	mu.Unlock()
	wg.Done()
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, s := range slice {
		wg.Add(1)
		go checkUnique(s, &wg, &mu)
	}

	wg.Wait()
	for key, value := range oleg {
		fmt.Printf("All symbols of (%v) are uniq: %v\n", key, value)
	}
}
