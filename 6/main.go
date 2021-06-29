package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(c chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case <-c:
			fmt.Println("Worker is finished")
			wg.Done()
			return
		default:
			fmt.Println("Wait for finish")
		}
	}
}

func closeWorker(c chan bool) {
	time.Sleep(2 * time.Second)
	c <- true
}

func main() {
	fmt.Println("Start")
	c := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(c, &wg)

	go closeWorker(c)

	wg.Wait()
	fmt.Println("End")
}
