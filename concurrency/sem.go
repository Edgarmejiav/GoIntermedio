package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSometing(i, &wg, c)

	}

	wg.Wait()
}

func doSometing(i int, wg *sync.WaitGroup, c chan int) {

	defer wg.Done()
	fmt.Printf("id %d stared\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("id %d Finished\n", i)
	<-c
}
