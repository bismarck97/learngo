package main

import (
	"fmt"
)

//在函数定义时调用函数本身 递归函数
//死递归
func Test(a int) {
	//在递归函数中一定要有出口 return
	//先在此函数中执行递归完再返回
	if a == 0 {
		return
	}
	a--
	Test(a)
	fmt.Println(a)
}

//举个例子，我们来计算阶乘n=1*2*3*....*n

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	Test(6)
	num := Factorial(5)
	fmt.Println(num)
}
