package main

import (
	"fmt"
	"time"
)

func mySleepInSecond1(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
	fmt.Printf("%v seconds left\n", seconds)
}

func mySleepInSecond2(seconds int) {
	start := time.Now()
	sec := time.Duration(seconds)
	for {
		now := time.Now()
		elapsed := now.Sub(start) / time.Second
		if sec <= elapsed {
			fmt.Printf("%v seconds left\n", seconds)
			return
		}
	}
}

func mySleepInSecond3(seconds int) {
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	<-timer.C
	fmt.Printf("%v seconds left\n", seconds)
}

func main() {
	fmt.Println("Start")
	fmt.Println("Start first sleep")
	mySleepInSecond1(2)
	fmt.Println("Start second sleep")
	mySleepInSecond2(3)
	fmt.Println("Start third sleep")
	mySleepInSecond3(4)
	fmt.Println("End")
}
