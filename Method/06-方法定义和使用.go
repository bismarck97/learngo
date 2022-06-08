package main

import "fmt"

//func add(a, b int) int {
//	return a + b
//}
//方法要绑定在对象上面

//type的意义：1.定义函数类型
//2.为已存在的数据类型起别名

type Integer int //取别名
//func(方法接收者)方法名(参数列表)返回值类型

func (i Integer) add(a Integer) Integer {
	return i + a
}

func main06() {
	//var a int = 100
	//var b int = 200
	//fmt.Println(add(a, b))
	//根据数据类型绑定方法
	var i Integer = 100
	value := i.add(200)
	fmt.Println(value)
}
