package main

import (
	"fmt"
	"time"
)

func parent() {
	go child()
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("parent: %d\n", i)
		//time.Sleep(500 * time.Millisecond)
	}
}

func child() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("child: %c\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main0() {
	go parent()
	//go child()
	fmt.Println("main")
	time.Sleep(5 * time.Second)
}

// 这里可以证明Go里没有父子协程的概念
//parent: 97
//parent: 98
//parent: 99
//parent: 100
//parent: 101
//child: a
//child: b
//child: c
//child: d
//child: e
