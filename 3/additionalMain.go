package main

import (
	"fmt"
	"sync"
)

var (
	slice = []int{2, 4, 6, 8, 10}
	sum   int
)

func squaring(index int, value int, wg *sync.WaitGroup, m *sync.Mutex) {
	sqv := value * value
	m.Lock()
	sum += sqv
	m.Unlock()
	wg.Done()
}

func main() {
	fmt.Printf("Started array = %v\n", slice)
	var wg sync.WaitGroup
	var m sync.Mutex
	for index, val := range slice {
		wg.Add(1)
		go squaring(index, val, &wg, &m)
	}
	wg.Wait()
	fmt.Printf("Sum = %v\n", sum)
}
