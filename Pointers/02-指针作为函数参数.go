package main

import "fmt"

//指针作为函数参数 可以对原来的值进行修改

func Swap(x, y *int) {
	*x, *y = *y, *x
}

func main02() {
	x := 1
	y := 2
	Swap(&x, &y)
	fmt.Println(x, y)
}
