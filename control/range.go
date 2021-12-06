package main

import "fmt"

/**
range 是 go提供的一种迭代遍历手段， 可操作的数据结构有： 数组、切片、channel、map

如过 range channel， channel中如果有数组 则读取出来，然后结束。（range chan 只能读取一个值） ， 如果chan中无值 ，则阻塞等待。如果chan已经关闭，则range也解除阻塞并退出


协程泄漏： 协程永远阻塞，无法退出，称为协程泄漏、

chan是CSP的具体实现。用多多个Go程间通信。内部确保并发安全。 chan是FIFO原则，多个Go程同时访问，不需要加锁
*/

func main() {

	//c := make(chan int)
	//
	//go func() {
	//	for i := 0; i < 10; i++  {
	//		c <- i
	//	}
	//	close(c)
	//}()
	//
	//for i := range c {
	//	fmt.Println(i)
	//}

	data := make(chan string)

	exit := make(chan bool)

	go func() {
		for d := range data {
			fmt.Println(d)
		}

		fmt.Println("received over!")
		exit <- true
	}()

	data <- "aaaaa"
	data <- "Linux"

	data <- "aaaaa"
	data <- "Linux"
	data <- "aaaaa"
	data <- "Linux"

	data <- "aaaaa"
	data <- "Linux"
	data <- "aaaaa"
	data <- "Linux"
	close(data)

	<-exit

}
