package main

import (
	"fmt"
	"sync"
	"time"
)

var qps = make(chan struct{}, 10)

func handler(i int) {
	qps <- struct{}{}
	defer func() {
		<-qps
	}()
	time.Sleep(5 * time.Second)
	fmt.Printf("[%d] handler执行了一次\n", i)
}

// 这里演示如何用channel限制接口的并发数量
// 但实际上就是限制同时执行handler函数的协程数量
// 简单的利用channel特性完成需求
func main() {
	const P = 1000
	wg := sync.WaitGroup{}
	wg.Add(P)
	for i := 0; i < P; i++ {
		go func(i int) {
			defer wg.Done()
			handler(i)
		}(i)
	}
	wg.Wait()
}
