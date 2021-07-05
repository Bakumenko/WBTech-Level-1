package main

import (
	"fmt"
	"sync"
)

var (
	slice3 = []interface{}{
		12321,
		12322,
		"123321",
		124321,
		"asa",
		"oleg",
	}

	resultMap = map[interface{}]bool{}
)

func isPalindromeUniversalConcurrency(targer interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {
	str := fmt.Sprintf("%v", targer)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			mu.Lock()
			resultMap[targer] = false
			mu.Unlock()
			wg.Done()
			return
		}
	}
	mu.Lock()
	resultMap[targer] = true
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, val := range slice3 {
		wg.Add(1)
		go isPalindromeUniversalConcurrency(val, &wg, &mu)
	}

	wg.Wait()
	for key, val := range resultMap {
		if val {
			fmt.Printf("%v is palindrome\n", key)
		} else {
			fmt.Printf("%v is not palindrome\n", key)
		}
	}
}
