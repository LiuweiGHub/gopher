package main

import (
	"log"
	"sync"
	"time"
)

/**
互斥锁：并发编程对共享资源控制的主要手段

mutex是一个结构体\\

mutex 重复解锁 panic的原因是什么？
 因为如果阻塞等待的协程很多，那么每次unlock都会释放信号量唤醒多个协程 ， 这会导致多个协程竞争锁，提升复杂度。而且会造成很多不必要的协程切换


*/

type Mutex struct {
	state int32 //表示互斥锁状态， 实际分成四个部分
	// 1.Locked 表是否锁定
	// 2.Woken 表示是否有协程已经被唤醒
	// 3.Starving 表示该Mutex是否受理饥饿状态（什么是饥饿状态： 有协程阻塞超过1ms ）
	// 4. Waiter 表示阻塞等待锁的协程个数，释放锁时根据此值判断是否需要释放信号量
	sema uint32 // 信号量，协程阻塞等待该信号量； 解锁的协程释放该信号量，从而唤醒等待该信号量的协程。
}

/**
自旋：
加锁时，如果当前Locked位为1，说明该锁当前由其他协程持有，尝试加锁的协程并不是马上转入阻塞，而是会持续 的探测Locked位是否变为0，这个过程即为自旋过程。
自旋时间很短，但如果在自旋过程中发现锁已被释放，那么协程可以立即获取锁。此时即便有协程被唤醒也无法获取 锁，只能再次阻塞。

自旋的好处是：当加锁失败时不必立即转入阻塞，有一定机会获取到锁，这样可以避免协程的切换。
自旋的问题： 如果加锁的协程特别多， 每次都通过自旋获得锁， 那么之前被阻塞的协程无法获取到锁，从而进入饥饿状态

怎么解决自旋产生的问题：
每个mutex都有两个模式， normal 和 starving，
normal：协程获取不到锁， 如果满足自旋就先进入自旋
starving：如果有很多协程阻塞等待， 释放锁的协程肯定会释放信号量唤醒一个协程，该协程唤醒抢锁时 发现已经被自旋的协程抢走了，自己只能再次阻塞，这时它会判断自身是否已经阻塞超过1ms。
如果超过1ms 则会将该mutex标记为饥饿模式，此时不会启动自旋过程，那一定会有协程被唤醒 （有协程切换）并成功获取锁。
*/

type foo struct {
	n int
	sync.Mutex
}

func main() {

	f := foo{n: 17}
	go func(f foo) {
		for {
			log.Println("g2: try to lock foo ...")
			f.Lock()
			log.Println("g2: lock foo ok")
			f.Unlock()
			log.Println("g2: Unlock foo ok")
		}
	}(f)

	f.Lock()
	log.Println("g1: lock foo ok")

	go func(f foo) {
		for {
			log.Println("g3：try to lock foo...")
			f.Lock()
			log.Println("g3: lock foo ok")
			time.Sleep(5 * time.Second)
			f.Unlock()
			log.Println("g3: unlock foo ok")
		}
	}(f) // 复制了一个已经locked的锁

	time.Sleep(1000 * time.Second)
	f.Unlock()
	log.Println("g1: unlock foo ok")

}
