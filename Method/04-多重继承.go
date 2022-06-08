package main

import "fmt"

type TestA struct {
	Name string
	Id   int
}
type TestB struct {
	TestA
	sex string
	age int
}

//结构体不能嵌套本结构体
//结构体可以嵌套本结构体指针类型类型 链表

type TestC struct {
	//*TestC
	TestB
	score int
}

func main04() {
	var stu TestC
	stu.Name = "张三"
	stu.Id = 101
	stu.age = 19
	stu.sex = "男"
	stu.score = 100
	fmt.Println(stu)
}
