package main

import (
	"fmt"
	"sync"
	"time"
)

/**
三个接口：
Add       add设置的值必须与实际等待的goroutine个数一致
Done
Wait
*/

func main() {
	var wg sync.WaitGroup

	wg.Add(2) //设置计数器，数值即为goroutine的个数

	go func() {
		time.Sleep(time.Second)

		fmt.Println("Gorputine 1 finished!")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second)
		fmt.Println("Goroutine 2 finished!")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All goroutine finished")
}

/**
补充知识点：

信号量是什么？
信号量是unix系统提供的一种保护共享资源的机制，用于防止多个线程同时访问某个资源

- 当信号量 > 0 时，表示资源可用，获取信号量时，系统将信号量自动减1
- 当信号量 = 0 时，表示资源暂不可用，获取信号量时，当前线程进入睡眠，当信号量为正时被唤醒


CAS算法： 比较交换算法  compareAndSwap

*/
//Add伪代码
//func  (wg *waitGroup) Add(delta int) {

//}
