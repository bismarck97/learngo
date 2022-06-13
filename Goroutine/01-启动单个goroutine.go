package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello world", i)
	//通知主线程结束
	wg.Done() //计数牌-1
}
func main01() { //开启一个主goroutine去指向main函数

	wg.Add(10000) //计数牌+1
	for i := 0; i < 10000; i++ {
		go hello(i) //开启了一个goroutine去指向hello这个函数
	}
	fmt.Println("hello main")
	//time.Sleep(time.Second)
	wg.Wait() //阻塞，等所有线程都执行完成之后才结束
}
