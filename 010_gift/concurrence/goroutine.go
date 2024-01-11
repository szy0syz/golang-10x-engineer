package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	cpuCount := runtime.NumCPU()
	fmt.Println("逻辑核心数: ", cpuCount)
	runtime.GOMAXPROCS(cpuCount / 2)

	const P = 1000000
	for i := 0; i < P; i++ {
		go time.Sleep(3 * time.Second)
	}

	fmt.Println("进程中存活的协程数：", runtime.NumGoroutine())
	//逻辑核心数:  10
	//进程中存活的协程数： 1000001
}
