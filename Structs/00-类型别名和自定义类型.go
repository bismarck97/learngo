package main

import "fmt"

//基于int类型的自定义类型

type Int int

//int的类型别名

type Integer = int

//一样的是值传递
func swapInteger(a, b Integer) {
	a, b = b, a
	fmt.Println(a, b)
}

//一样的是值传递
func swapInt(a, b Int) {
	a, b = b, a
	fmt.Println(a, b)
}
func main00() {
	var a Integer = 1
	var b Integer = 2
	swapInteger(a, b)
	fmt.Println(a, b)
	fmt.Println("==========================")
	var c Int = 3
	var d Int = 4
	swapInt(c, d)
	fmt.Println(c, d)
	fmt.Printf("类型别名：%T\n", a)

	fmt.Printf("自定义类型：%T\n", c)
}
