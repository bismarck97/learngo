package main

import (
	"fmt"
)

//其他包调用函数要大写，类似于public

//1.循环遍历
func test(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

//2.计算总和
func Plus(arr ...int) {
	sum := 0
	//for i := 0; i < len(arr); i++ {
	//	sum += arr[i]
	//}


	for _, arg := range arr {
		sum += arg
	}

	fmt.Println(sum)
}

func main02() {
	//循环遍历
	//function.Test(1, 2, 3, 4, 5)
	//计算总和
	Plus(1, 2, 3, 4, 5)

}
