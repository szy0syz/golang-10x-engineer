package main

import (
	"fmt"
	"time"
)

// 演示遍历Channel

var ch = make(chan int, 3)

func add2Ch() {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func traveseChannel() {
	//for ele := range ch { // 遍历取走管道中的元素
	//	fmt.Println(ele)
	//}

	// 另一种写法
	for {
		ele, ok := <-ch
		if !ok { // ch已空并关闭
			break
		}
		fmt.Println(ele)
	}
	fmt.Println("bye")
}

func main7() {
	go add2Ch()
	go traveseChannel()

	time.Sleep(3 * time.Second)
}
