package main

import (
	"fmt"
	"time"
)

func worker2(flag *bool) {
	for {
		if *flag {
			fmt.Println("Worker is finished")
			return
		}
	}
}

func main() {
	fmt.Println("Start")
	flag := false
	go worker2(&flag)

	time.Sleep(time.Second)
	flag = true
	time.Sleep(time.Second)
	fmt.Println("End")
}
