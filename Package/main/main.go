package main

import (
	"fmt"
	ca "learngo/package/calc"
)

//Go语言不允许导入包而不使用
//Go语言中不允许循环导入包

//多用于来做一些初始化的操作，类似java的静态代码块
func init() {
	fmt.Println("main.init()")
}
func main() {
	ret := ca.Add(1, 2)
	fmt.Println(ret)
}
