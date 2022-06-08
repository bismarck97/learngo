package main

import (
	"fmt"
)

func Cal(n1, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("输入运算符错误!")

	}
	return res
}
func main() {
	res := Cal(1.5, 2.6, '-')
	fmt.Println(res)
}
