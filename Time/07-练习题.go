package main

import (
	"fmt"
	"time"
)

func test1() {
	//获取当前时间，格式化输出为2017/06/19 20:30:05格式。
	now := time.Now()
	t := now.Format("2006/01/02 15:04:05")
	fmt.Println(t)
}
func test2(f func()) {
	//编写程序统计一段代码的执行耗时时间，单位精确到微秒。
	//获取开始时长
	start := time.Now()
	//闭包调用要执行的函数
	f()
	//调用Sub求起始时间到执行后的时间
	sub := time.Now().Sub(start)
	//fmt.Println(time.Since(start))//time.Since(start) 等价于 time.Now().Sub(start)
	fmt.Println(sub)
	fmt.Printf("执行总时间%v微秒\n", int(sub)/1000)
}

func main07() {
	test2(test1)
}
