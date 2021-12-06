package main

/**
所谓读写锁，其实就是读写互斥锁， 是mutex的改进版

读写锁解决什么问题？
程序中读多写少，如果使用mutex， 所有读写都会串行化，性能大大降低

读写锁 可以 让  多个读操作同时持有锁，并发能力大大 提升
*/

type RwMutex struct {
	m           Mutex  // 用于控制写锁,获得写锁首先要获取该锁，如果有一个写锁在进行，那么再到来的写锁将会阻塞于此
	writerSem   uint32 //写阻塞等待的信号量，最后一个读者释放锁时会释放信号量
	readerSem   uint32 //读阻塞的协程等待的信号量，持有写锁的协程释放锁后会释放信号量
	readerCount int32  // 记录读者个数
	readerWait  int32  // 记录写阻塞时，读者个数

}

/**

RwMutex 四个接口

RLock()
RUnlock()
Lock()
Unlock()
*/

/**
写操作 如何 阻止写操作 ？
依赖互斥锁，同mutex  竞争locked的赋值权

写操作 如何 阻止读操作？（最精华的技巧）
我们知道RWMutex.readerCount是个整型值，用于表示读者数量，不考虑写操作的情况下，每次读锁定将该值+1， 每次解除读锁定将该值-1，所以readerCount取值为[0, N]，
N为读者个数，实际上最大可支持2^30个并发读者。
当写锁定进行时，会先将readerCount减去2^30，从而readerCount变成了负值，
此时再有读锁定到来时检测到 readerCount为负值，便知道有写操作在进行，只好阻塞等待。
而真实的读操作个数并不会丢失，只需要将 readerCount加上2^30即可获得。

所以，写操作将readerCount变成负值来阻止读操作的/

读操作 如何 阻止写操作？
读锁定会先将RWMutext.readerCount加1，此时写操作到来时发现读者数量不为0，会阻塞等待所有读操作结束。 所以，读操作通过readerCount来将来阻止写操作的。

为什么写锁定不会被饿死？
写操作到来时，会把RWMutex.readerCount值拷贝到RWMutex.readerWait中，用于标记排在写操作前面的读者 个数。 前面的读操作结束后，除了会递减RWMutex.readerCount，还会递减RWMutex.readerWait值，当 RWMutex.readerWait值变为0时唤醒写操作。


*/
