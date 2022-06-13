package main

import "fmt"

func main0301() {
	//defer调用的函数并没有直接使用 而是先加载到栈区内存中，再函数结束时进行调用，从后向前调用
	defer fmt.Println("hello")
	defer fmt.Println("world")
	//defer 函数调用
	defer fmt.Println("你好")
	fmt.Println("世界")
}
func test4() {
	fmt.Println("hello world")
}

var value int

func test5(a, b int) {
	value = a + b
}
func test6(value int) {
	fmt.Println(value)
}
func main0302() {
	test4()
	//函数中有返回值不能使用defer调用
	defer test5(10, 20)
	test6(value)
}
func main() {
	a := 10
	b := 20
	//定义函数类型变量
	func(a, b int) { //10，20  变量先传入函数
		fmt.Println(a)
		fmt.Println(b)
	}(a, b)
	//defer func() {//100，200  函数最后去找a和b
	//	fmt.Println(a)
	//	fmt.Println(b)
	//}()
	a = 100
	b = 200
	//先打印此结果
	fmt.Println(a)
	fmt.Println(b)
}
