package main

import (
	"fmt"
	"runtime"
	"time"
)

type GoroutineLimiter struct {
	limit int
	ch    chan struct{}
}

func NewGoroutineLimiter(n int) *GoroutineLimiter {
	return &GoroutineLimiter{
		limit: n,
		ch:    make(chan struct{}, n),
	}
}

func (g *GoroutineLimiter) Run(f func()) {
	g.ch <- struct{}{} // 把元素放入管道环

	go func() {
		f()
		<-g.ch // 把元素从管道里腾出位置来
	}()
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		// 每隔一秒打印一次协程数量
		for {
			<-ticker.C
			fmt.Printf("当前协程数量：%d\n", runtime.NumGoroutine())
		}
	}()

	limiter := NewGoroutineLimiter(100)

	work := func() {
		time.Sleep(10 * time.Second)
	}
	for i := 0; i < 10000; i++ {
		// 我们不能直接调用work，而是统一一个入口，让入口来限制
		//go work()
		limiter.Run(work)
	}
	time.Sleep(10 * time.Second)
}
