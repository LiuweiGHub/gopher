package main

/**
线程池解决什么问题：
解决线程频繁创建和销毁带来的开销、


线程池有什么问题：
如果线程池中的线程执行任务时，产生系统调用，则会被阻塞。如果大部分任务都产生系统调用，则线程池的线程会大量阻塞，消费能力会急剧下降。
引发任务的堆积。

虽然增加线程数量 可以在一定程度上缓解这种情况，但是过量的线程争抢CPU， 可能不仅消费能力不会上升，反而会下降、

*/

/**
Go如何解决上述问题？

GMP模型。 go实现了在线程中调度goroutine， 从而实现了线程少 并发高

*/

/**
CMP

p的个数在程序启动时确认， 默认是CPU个数。
由于m必须持有一个p，才能运行go代码， 所以同时运行的 M个数，也即线程数一般等同于CPU的个数，以达到尽可能的使用CPU而又不至于产生过多的线程切换开销。



队列轮转
系统调用
工作量窃取
''一般来讲，程序运行时就将GOMAXPROCS大小设置为CPU核数，可让Go程序充分利用CPU。
*/
