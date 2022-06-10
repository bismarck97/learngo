package main

import "fmt"

func main01() {
	var p *int
	fmt.Println(p)
	//为指针变量创建一块内存空间

	p = new(int) //new函数返回的是一个指针  创建好的空间值为数据类型的默认值 0
	//打印p的值
	fmt.Println(p)
	//打印p指向的内存地址的值
	fmt.Println(*p)
	fmt.Println("===========================")
	q := new(int)
	*q = 1000
	fmt.Println(*q)
}
