package main

import "fmt"

func main03() {
	a := 10
	b := 20
	//匿名内部函数
	//func(a, b int) {
	//	fmt.Println(a + b)
	//}(a, b)

	//f是函数类型对应的变量
	f := func(a, b int) {
		fmt.Println(a + b)
	}
	f(a, b)
	f(20, 30)
	fmt.Printf("%T\n", f)
	fmt.Println("===============================")
	//v是整型
	v := func(a, b int) int {
		return a + b
	}(a, b)
	fmt.Println(v)
	fmt.Printf("%T\n", v)
}
