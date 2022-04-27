package main

import (
	"fmt"
	"time"
)

// Multiplexacion

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)
	for i := 0; i < 2; i++ {
		select {
		case chanelMgs1 := <-c1:
			fmt.Println(chanelMgs1)
		case chanelMgs2 := <-c2:
			fmt.Println(chanelMgs2)
		}
	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
