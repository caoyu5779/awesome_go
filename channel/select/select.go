package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
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
	var c1, c2 = generator(), generator()
	var worker = createWorker(0) //nil channel

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	// 解决的问题是 同时收 c1,c2 的数据，谁来的快，收谁
	// c1 and c2 = nil
	for {
		var activeWorker chan<- int
		var activeValue int

		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n) //收到的值 缓存起来
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick: //反应系统状态
			fmt.Println("queue length = ", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
