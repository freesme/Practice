package main

import (
	"fmt"
	"time"
)

var shouldQuit = make(chan struct{})

func main() {

	var c1, c2, c3 chan int
	var i1, i2 int

	// 用于处理异步IO操作
	select {
	case i1 = <-c1:
		fmt.Println("received", i1, "from c1")
	case i2 = <-c2:
		fmt.Println("received", i2, "from c2")
	case c3, ok := (<-c3):
		if ok {
			fmt.Println(c3, "from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}

	var resChan = make(chan int)
	// 超时处理
	select {
	case data := <-resChan:
		doData(data)
	case <-time.After(time.Second * 2):
		fmt.Println("2 second timeout")
	}
	// 在某些情况下是存在不希望channel缓存满了的需求的，可以用如下方法判断
	ch := make(chan int, 5)
	data := 0
	select {
	case ch <- data:
	default:
		// 做相应操作，比如丢弃data。视需求而定
	}
}

func doData(data int) {
	fmt.Println("do data ", data)
}
