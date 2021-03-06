package main

import (
	"context"
	"fmt"
	"time"
)

/**
context 比 waitGroup 的优势在于， 它可以控制 协程 派生的goroutine， 即既管你，又管你孩子
控制多级goroutine
*/

/**
context 接口 ， 提供四个方法

1.Deadline() (deadline time.Time, ok bool)
2.Done() <-chan struct{}
3.Err() error
4.value(key interface{}) interface{}
*/

/**
官方四种实现：
emptyCtx
cancelCtx    cancel时会把所有child都cancel
timerCtx
valueCtx


关系
             emptyCtx
context ---> cancelCtx ----->timerCtx
             valueCtx
*/

/**
WithcCancel() 做三件事；
1. 初始化cancelCtx
2. 将1创建的实例加入父节点的children中（父节点可以被cancel，如果不支持就继续往上查， 如果都不支持，就启动一个协程等待父节点结束，然后再关闭该context）
3. 返回实例和cancel（）方法
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutine.")

	cancel()

	time.Sleep(5 * time.Second)
}

func HandelRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest Running.")
			time.Sleep(2 * time.Second)
		}

	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis Running.")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase Running.")
			time.Sleep(2 * time.Second)
		}
	}
}
