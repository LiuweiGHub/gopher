package main

import (
	"os"
	"runtime/pprof"
	"time"
)

// pprof Go自带的性能分析库

/*
1.CPU Profile
2.Memory Profile(Heap Profile)
3.Block Profile
4.Goroutine Profile
*/

/*
1.针对工具类：使用 runtime/pprof 库
2.针对服务：使用 net/http/pprof 库

*/

func main() {
	f, _ := os.Create("/tmp/cpuprofille")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	time.Sleep(time.Second * 3)

	f1, _ := os.Create("/tmp/memprofile")
	pprof.WriteHeapProfile(f1)
	f1.Close()
}
