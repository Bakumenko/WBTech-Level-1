package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100       //меняет значение в ячейке с индексом 0
	v = append(v, b) //добавляем новый элемент в слайс, тем самым создаем новый локальный массив, на который указывает слайс

}
func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
