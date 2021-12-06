package main

import (
	"fmt"
	"time"
)

/**
并发控制手段：
Channel ： 使用channel控制子协程
WaitGroup : 使用信号量控制
Context： 使用上下文控制

*/

func main() {
	channels := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go Process(channels[i])
	}

	for i, ch := range channels {
		<-ch
		fmt.Println("Routine ", i, "quit!")
	}
}

func Process(ch chan int) {
	time.Sleep(time.Second * 2)

	ch <- 1
}
