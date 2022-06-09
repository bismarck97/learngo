package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

//函数作为参数
func cal(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func Sub(x, y int) int {
	return x - y
}
func main08() {
	r1 := cal(1, 2, Add)
	fmt.Println(r1)

	r2 := cal(1, 2, Sub)
	fmt.Println(r2)
}
