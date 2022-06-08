package function

import "fmt"

//为了让其他包的文件使用Cal函数，需要将C大写 类似其他语言的public

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
