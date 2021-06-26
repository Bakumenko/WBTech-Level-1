package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done() //Done выполняется для копии wg из анонимной функции и никак не влияет на WaitGroup в main
		}(wg, i) //передаем копию wg
	}

	//для корректной работы нужно передавать указатель на wg:
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(wg *sync.WaitGroup, i int) {
	//		fmt.Println(i)
	//		wg.Done() //Done выполняется для копии wg из анонимной функции
	//	}(&wg, i) //передаем копию wg
	//}
	wg.Wait()
	fmt.Println("exit")
}
