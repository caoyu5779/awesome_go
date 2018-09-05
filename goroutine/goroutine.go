package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutine 协程 Coroutine
// 轻量级 线程
// 非抢占式多任务处理， 由协程主动交出控制权
// 编译器，解释器，虚拟机层面的多任务
// 多个协程 可以在一个或多个线程上运行

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { //race condition! 数据访问冲突
			for {
				a[i]++
				//手动交出控制权 让其他协程 运行
				runtime.Gosched()
				//fmt.Printf("hello from goroutine %d \n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
