package main

import (
	"fmt"
	"time"
)

func main() {
	TickerDemo()
}

func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Printf("%s生病请上好大夫%s\n", "ss", "aa")
	}
}

//定时聚合任务 公交车场景
func TickerLaunch() {

}
