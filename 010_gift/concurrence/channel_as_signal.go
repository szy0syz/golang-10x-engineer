package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		fmt.Println("子协程结束")
		ch <- struct{}{}
	}()
	<-ch // read channel

	testEmptyStruct()
}

type A struct{}
type B struct{}

// 空结构体不占用内存，获取他们地址时runtime会返回统一的值
func testEmptyStruct() {
	a := A{}
	b := B{}
	fmt.Println("%p, %p", &a, &b)
	typeA := reflect.TypeOf(a)
	typeB := reflect.TypeOf(b)
	fmt.Printf("%d, %d", typeA.Size(), typeB.Size())
	// %p, %p &{} &{}
	// 0,  0
}
