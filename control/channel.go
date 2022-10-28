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

var active = make(chan struct{}, 2)
var jobs = make(chan int, 10)

func main() {

	var c1 chan int
	<-c1

	c := make(chan int, 3)
	go producer(c)
	go consumer(c)
	select {} //临时阻塞，用于演示

	// go func() {
	// 	for i := 0; i < 8; i++ {
	// 		jobs <- (i + 1)
	// 	}
	// 	close(jobs)
	// }()

	// var wg sync.WaitGroup

	// for j := range jobs {
	// 	wg.Add(1)
	// 	go func(j int) {
	// 		active <- struct{}{}
	// 		fmt.Printf("handle jod: %d\n", j)
	// 		time.Sleep(2 * time.Second)
	// 		<-active
	// 		wg.Done()
	// 	}(j)
	// }
	// wg.Wait()

	// 声明错误 使用 errors.New()
	// e := errors.New("test")

	// fmt.Println(e)
	// channels := make([]chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	channels[i] = make(chan int)
	// 	go Process(channels[i])
	// }

	// for i, ch := range channels {
	// 	<-ch
	// 	fmt.Println("Routine ", i, "quit!")
	// }
}

// func Process(ch chan int) {
// 	time.Sleep(time.Second * 2)

// 	ch <- 1
// }

func producer(c chan<- int) {
	var i int = 1
	for {
		time.Sleep(2 * time.Second)
		ok := trySend(c, i)
		if ok {
			fmt.Printf("[Producer]: send [%d] to channel\n", i)
			i++
			continue
		}
		fmt.Printf("[Producer]: try send [%d], but channel is full.\n", i)
	}
}

func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

func consumer(c <-chan int) {
	for {
		i, ok := tryRecv(c)
		if !ok {
			fmt.Printf("[Consumer]: try recv from channel, but the channel is empty.\n")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("[Consumer]: recv [%d] from channel.\n", i)
		if i >= 3 {
			fmt.Println("[consumer]: exit")
			return
		}
	}
}

func tryRecv(c <-chan int) (int, bool) {
	select {
	case i := <-c:
		return i, true
	default:
		return 0, false
	}
}
