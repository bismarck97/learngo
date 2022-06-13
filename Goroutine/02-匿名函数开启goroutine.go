package main

import (
	"fmt"
)

func main() {
	wg.Add(10000) //计数牌+1
	for i := 0; i < 10000; i++ {
		go func(i int) { //不传参数的话就是个闭包
			fmt.Println("hello", i)
			wg.Done() //计数牌-1
		}(i) //开启了一个goroutine去指向hello这个函数
	}
	fmt.Println("hello main")
	wg.Wait() //阻塞，等所有线程都执行完成之后才结束
}
