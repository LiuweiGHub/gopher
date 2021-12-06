package main

import (
	"fmt"
	"time"
)

/**
单向channel， 只能用于接收或发送数据
并没有真正单向channel， 只是一种使用限制
只读： chanName <-chan int
只写： chanName chan<- int
*/

func readChan(c <-chan int) {
	<-c
}

func writeChan(c chan<- int) {
	c <- 1
}

func main() {
	var myChan = make(chan int, 10)
	writeChan(myChan)
	readChan(myChan)

	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)
	go addNumToChan(chan1)
	go addNumToChan(chan2)

	for {
		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("get element fron chan2: %d\n", e)
		default:
			fmt.Printf("No element in chan1 or chan2 \n")
			time.Sleep(1 * time.Second)

		}
	}
}

/**
1.select可以监控多个channel， 当其中一个有数据时 就将数据读出来
2.select多个case执行顺序是随机的
3.select语句读channel不会阻塞，编译后明确传入不阻塞参数
*/

func addNumToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

/**
range可以持续从channel中读出数据，类似遍历，如果channel中没有数组会阻塞当前协程
*/

func chanRange(chanName chan int) {
	for e := range chanName {
		fmt.Printf("get element from chan %d \n", e)
	}
}
