package main

import (
	"fmt"
	"sync"
)

var (
	slice4 = []interface{}{
		12321,
		12322,
		"123321",
		124321,
		"asa",
		"oleg",
	}

	resultSyncMap = sync.Map{}
)

func isPalindromeUniversalSyncMapConcurrency(targer interface{}, wg *sync.WaitGroup) {
	str := fmt.Sprintf("%v", targer)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			resultSyncMap.Store(targer, false)
			wg.Done()
			return
		}
	}
	resultSyncMap.Store(targer, true)
	wg.Done()
}

func syncMapToMap(resultSyncMap sync.Map) map[interface{}]bool {
	m := map[interface{}]bool{}
	resultSyncMap.Range(func(key, value interface{}) bool {
		m[fmt.Sprint(key)] = value.(bool)
		return true
	})
	return m
}

func main() {
	var wg sync.WaitGroup
	for _, val := range slice4 {
		wg.Add(1)
		go isPalindromeUniversalSyncMapConcurrency(val, &wg)
	}

	wg.Wait()

	m := syncMapToMap(resultSyncMap)
	for key, val := range m {
		if val {
			fmt.Printf("%v is palindrome\n", key)
		} else {
			fmt.Printf("%v is not palindrome\n", key)
		}
	}
}
