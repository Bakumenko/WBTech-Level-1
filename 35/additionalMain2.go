package main

import "fmt"

var (
	slice2 = []interface{}{
		12321,
		12322,
		"123321",
		124321,
		"asa",
		"oleg",
	}
)

func isPalindromeUniversal(targer interface{}) bool {
	str := fmt.Sprintf("%v", targer)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	for _, val := range slice2 {
		if isPalindromeUniversal(val) {
			fmt.Printf("%v is palindrome\n", val)
		} else {
			fmt.Printf("%v is not palindrome\n", val)
		}
	}
}
