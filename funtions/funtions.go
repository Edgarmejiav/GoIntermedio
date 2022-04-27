package main

import "time"

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()
	println(y)

	c := make(chan int)
	go func() {
		println("Start Funtion")
		time.Sleep(5 * time.Second)
		println("end Funtion")
		c <- 1

	}()
	<-c
}
