package main

import (
	"fmt"
	"sync"
	"time"
)

// 为了避免刚刚那种低效的控制协程
// 我们使用sync.WaitGroup

var (
	wg = sync.WaitGroup{}
)

func init() {
	wg.Add(2)
}

func parent1() {
	go child1()
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("parent: %d\n", i)
		//time.Sleep(500 * time.Millisecond)
	}
	defer wg.Done()
}

func child1() {
	defer wg.Done()
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("child: %c\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main1() {
	go parent1()
	fmt.Println("main")
	wg.Wait()
}

//main
//child: a
//parent: 97
//parent: 98
//parent: 99
//parent: 100
//parent: 101
//child: b
//child: c
//child: d
//child: e
