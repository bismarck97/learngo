package main

import "fmt"

func demo1(i int) {
	//recover()使用一堆要在错误出现之前 进行拦截
	//通过匿名函数和recover()进行错误拦截
	defer func() {
		//recover可以从panic异常中重新获取控制权
		//返回值为接口类型
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(1)
	var p *int
	*p = 123 //err //recover作用范围
	fmt.Println(2)
	var arr [10]int
	//如果传递超出数组下班值 导致数组下标越界
	arr[i] = 100 //系统处理 导致程序崩溃
	fmt.Println(arr)
}
func demo2() {
	fmt.Println("hello world")
}
func main() {
	demo1(10)
	//可以正常执行下一个函数
	demo2()
}
