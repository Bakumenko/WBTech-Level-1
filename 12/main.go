package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
	//чтобы изменить значение p нужно раизменовать p (получить ячейку на которую указывает p) и присвоить новое значение
	//*p = b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
