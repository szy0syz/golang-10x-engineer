package main

import (
	"fmt"
	"sync"
	"time"
)

// 演示遍历Channel

var twg sync.WaitGroup

func add2Ch1() {
	defer twg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//close(ch)
}

func traverseChannel1() {
	defer twg.Done()
	for ele := range ch {
		fmt.Println(ele)
	}
	fmt.Println("bye")
}

func main() {
	twg.Add(2)
	go add2Ch1()
	go traverseChannel1()
	go func() {
		time.Sleep(time.Hour)
		// 暂时性防止死锁，但runtime没有上帝模式，不知道这个睡眠以后是否会解锁，所以只能憨憨的等待
	}()
	twg.Wait()
}
