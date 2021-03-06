package main

/**
GO 内存管理
*/
/**
Go抛弃了C/C++中开发中管理内存的方式， 增加了逃逸分析和GC
主要是指管理 堆内存。 GO栈内存，开发不用关心。


计算机的存储体系：
- CPU寄存器
- Cache  (cache是内存数据的缓存，可以降低cpu访问内存时间) 1级cache --> 3级cache
   - L1  // L1速率最快, 与CPU速率最接近
   - L2
   - L3
- 内存
- 硬盘等辅助存储设备
- 鼠标等


虚拟内存：
屏蔽了底层RAM和磁盘
某进程访问数据时，当cache没有命中，就去访问虚拟内存获取数据
访问内存，实际访问的是虚拟内存，虚拟内存通过页表查看，当前要访问的虚拟内存地址，是否已经加载到物理内存，如果已
加载，就获取；如果没有，就从磁盘加载到物理内存，并把物理内存地址和虚拟内存地址更新到页表。
引入虚拟内存后，每个进程都要有各自的虚拟内存，这样内存的并发访问问题的粒度就从多进程级别，降低到多线程级别。



堆和栈：
原来 从虚拟内存再进一层，看虚拟内存中的堆和栈， 也就是进程对内存的管理。
代码中使用的内存地址都是虚拟内存地址，而不是真实的物理内存地址。
堆和栈只是虚拟内存上 两块不同功能的内存区域：
- 栈在高地址，从高地址向低地址增长；
- 堆在低地址，从低地址向高地址增长；
栈内存有更好的局部性， 堆内存访问的两块数据，如果不在同一页上，CPU访问数据的时间可能就上去了

栈内存管理：
- 栈的内存不需要回收
- 分配比堆上快

堆（heap）内存管理：
- 堆内存无论是主动free  还是被动垃圾回收，都需要消耗CPU


内存对齐
释放内存
内存碎片
*/

/*
TCMalloc:

引入虚拟内存后，让内存的并发访问问题的粒度从多进程级别，降低到多线程级别。
这是更快分配内存的第一个层次。

TCMalloc是Thread Cache Malloc的简称，是Go内存管理的起源，Go的内存管理是借鉴了TCMalloc，随着Go的迭代，Go的内存管理与TCMalloc不一致地方在不断扩大，但其主要思想、原理和概念都是和TCMalloc一致的

同一进程的所有线程共享相同的内存空间，他们申请内存时需要加锁，如果不加锁就存在同一块内存被2个线程同时访问的问题。

TCMalloc的做法是什么呢？为每个线程预分配一块缓存，线程申请小内存时，可以从缓存分配内存，这样有2个好处：

1、为线程预分配缓存需要进行1次系统调用，后续线程申请小内存时，从缓存分配，都是在用户态执行，没有系统调用，缩短了内存总体的分配和释放时间，这是快速分配内存的第二个层次。
2、多个线程同时申请小内存时，从各自的缓存分配，访问的是不同的地址空间，无需加锁，把内存并发访问的粒度进一步降低了，这是快速分配内存的第三个层次。
*/

/**
每个go程都有自己的栈，栈的初始化大小是 2k， 100万的go程需要2G内存， 会自动扩容。
*/

/**
空间换时间的思想
空间换时间是一种常用的性能优化思想，这种思想其实非常普遍，比如Hash、Map、二叉排序树等数据结构的本质就是空间换时间，在数据库中也很常见，比如数据库索引、索引视图和数据缓存等，再如Redis等缓存数据库也是空间换时间的思想。
*/
