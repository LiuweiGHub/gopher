第九章 定时器

1.Timer使用场景
    1> 设置超时时间
    2> 延迟执行

2.Timer动作
    1> 创建(自动启动)
    2> stop  返回值为bool， 有效期内返回true
    3> 重置 reset  返回值不可靠


 原理篇：
 每个go程序，都有一个单独的协程负责管理所有的Timer，该协程监控到某timer过期后，则发送当前时间到管道中，也就是timer.C

 Timer数据结构
 Timer.C即面向Timer用户的，Timer.r是面向底层的定时器实现
 每创建一个Timer， 意味着创建一个runtimeTimer变量，然后交给系统监控
 type Timer struct {
    C  <-chan Time   // 管道，上层应用跟据此管道接收事件
    r runtimeTimer   // runtime定时器，该定时器即系统管理的定时器，对上层应用不可见
 }



3.Ticker  周期型定时器， 数据结构与Timer完全一致
ticker使用完后，要进行资源释放，否则会产生资源泄漏

Ticker跟Timer的重要区就是提供了period这个参数，据此决定timer是一次性的，还是周期性的

Ticker不能重置

必须主动的 显示的停止， 否则不断消耗CPU


3.timer
一次性定时器Timer和周期性定时器Ticker都是依赖底层的timer记时

回顾一下定时器创建过程，创建Timer或Ticker实际上分为两步： 1. 创建一个管道 2. 创建一个timer并启动（注意此timer不是Timer，而是系统协程所管理的timer。）
