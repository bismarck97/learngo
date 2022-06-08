package main

import "fmt"

func swap(p *[5]int) {
	(*p)[0] = 89
}

func main03() {
	//数组指针 数组作为函数参数时，是作为值传递的，用数组指针可以进行地址传递
	a := [5]int{1, 2, 3, 4, 5}

	var p *[5]int = &a
	swap(p)
	fmt.Println(a)
}
