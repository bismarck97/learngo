package main

import "fmt"

//我们可以使用type关键字来定义一个函数类型，具体格式如下：
//定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
type calculation func(int, int) int

func add1(x, y int) int {
	return x + y
}
func sub1(x, y int) int {
	return x - y
}
func main05() {
	//声明一个calculation类型的变量，并且赋值为add函数 函数作为变量赋值没有()
	var cal calculation
	cal = add1
	fmt.Printf("type of c:%T\n", cal)
	//调用cal函数，并且传入两个int类型的参数
	fmt.Println(cal(1, 2))

	//将函数add负责给变量f1
	f := add1
	fmt.Printf("type of f:%T\n", f)
	fmt.Println(f(1, 2)) //像调用和add一样调用f
}
