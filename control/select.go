package main

import (
	"fmt"
	"time"
)

/**
select 是go提供调度语言 层面的 IO多路复用机制
select 中各个case执行顺序是随机的
如果所有的case都不可读，且没有default分支的情况下，程序会阻塞
已 关闭的channel也是可读的， close操作也是一种输入,close的channel也能返回，此时 ok = false
*/
func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 1
		time.Sleep(time.Second * 5)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready")
	case <-chan2:
		fmt.Println("chan2 ready")
	default:
		fmt.Println("default")

	}

	chan3 := make(chan int)
	chan4 := make(chan string)

	go func() {
		close(chan2)
	}()

	go func() {

		close(chan3)
	}()

	select {
	case i, ok := <-chan4:
		fmt.Println(i, ok, "chan4")
	case i, ok := <-chan3:
		fmt.Println(i, ok, "chan3")

	}
}

/**
总结：
select 语句中除default外， 每个case只能操作一个channel， 要么读要么写
select 语句中除default外， case的执行顺序是随机的

如果没有default语句，则 会阻塞等待任一case
select语句中读操作要判断是否读取成功，关闭的chan也可以读取
*/
