package main

import "fmt"

var (
	slice = []int{
		12321,
		12322,
		123321,
		124321,
	}
)

func isPalindrome(number int) bool {
	copyNumber := number
	var reverse int

	for copyNumber != 0 {
		remainder := copyNumber % 10
		reverse = reverse*10 + remainder
		copyNumber /= 10
	}

	return number == reverse
}

func main() {
	for _, val := range slice {
		if isPalindrome(val) {
			fmt.Printf("%v is palindrome\n", val)
		} else {
			fmt.Printf("%v is not palindrome\n", val)
		}
	}
}
