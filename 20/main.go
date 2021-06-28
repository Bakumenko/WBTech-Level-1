package main

import "fmt"

func main() {
	slice := []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a") //создается новый массив на который ссылается копия слайса
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice)
	}(slice)
	fmt.Println(slice)
}
