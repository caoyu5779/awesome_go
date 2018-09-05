package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	//var c chan int // c == nil
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d Received %d\n",
			id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)

	go worker(id, c)

	return c
}

func main() {
	chanDemo()
}
