package main

import "github.com/szy0syz/golang-10x-engineer/gift/database"

func main() {
	ch := make(chan database.Gift, 100)

	go func() {
		for i := 0; i < 10000; i++ {
			ch <- database.Gift{}
		}
		close(ch)
	}()

	for {
		// 如果读写速率不一致，则会导致这行代码阻塞
		gift, ok := <-ch
		if !ok { // ok == false的条件必须满足如下两个条件：1.Channel已空 2.Channel已关闭
			break // Channel 数据空且已关闭
		}
		// 写入redis缓存
		_ = gift
	}
}
