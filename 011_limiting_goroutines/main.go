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
	g.ch <- struct{}{}
	go func() {
		f()
		<-g.ch
	}()
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	go func() {
		for {
			<-ticker.C
			fmt.Printf("当前协程数: %d\n", runtime.NumGoroutine())
		}
	}()

	limiter := NewGoroutineLimiter(100)
	work := func() {
		time.Sleep(5 * time.Second)
	}
	for i := 0; i < 10000; i++ {
		//go work()
		limiter.Run(work)
	}
	time.Sleep(11 * time.Second)
}
