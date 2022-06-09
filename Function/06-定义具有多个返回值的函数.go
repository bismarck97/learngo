package main

import "fmt"

//定义具有多个返回值的函数,多返回值也支持类型简写

func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
func main06() {
	//函数调用
	x, y := calc(100, 200)
	fmt.Println(x, y)
}
