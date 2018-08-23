package main

import (
	"fmt"
	"sync"
)

func chanDemo() {
	//var c chan int // c == nil
	var wg sync.WaitGroup // 系统提供的 库
	var workers [10]worker
	for i := 0; i<10; i++{
		workers[i] = createWorker(i ,&wg)
	}
	wg.Add(20)

	for i, worker := range workers{
		worker.in <-'a'+i
	}

	for i, worker := range workers{
		worker.in <-'A'+i
	}
	wg.Wait()

	////wait for all of them then exit
	//for _,worker := range workers{
	//	<- worker.done
	//	<- worker.done
	//}
}

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d Received %c\n",
			id, n)
		//go func() {
		//	done <- true //这里因为 内部 在等待接收 done， 外部调用在往里传递 数据，导致死锁。
		//	// 这里使用一个goroutine并发向外返回 done，结果死锁问题。 但是个人认为这里还是有 概率死锁的
		//}()
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker{
	w := worker{
		in : make(chan int),
		done : func() {
			wg.Done()
		},
	}
	go doWorker(id, w)

	return w
}

func main() {
	chanDemo()
}
