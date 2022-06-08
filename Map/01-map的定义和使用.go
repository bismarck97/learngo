package main

import "fmt"

func main01() {
	//map的定义
	//key必须是支持==和!=操作符的类型
	//定义一个key是整数value是字符串的map
	//两种方式：定义
	var m1 map[int]string
	fmt.Println(m1)

	//指定map容量，容量不够可以自动扩容
	m2 := make(map[int]string, 3)
	fmt.Println(m2)
	//直接使用key赋值，key在字典中是唯一的，不能重复。map输出也是无序的
	m2[1] = "张三"
	m2[2] = "李四"
	m2[3] = "王五"
	m2[4] = "赵六"
	m2[5] = "孙七"
	fmt.Println(m2)

	//也可以直接定义时赋值
	m3 := map[int]string{1: "hello", 2: "world", 3: "Go"}
	fmt.Println(m3)

}
