package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	sex  string
}

func main08() {
	//定义结构体变量
	var stu Student = Student{
		id:   101,
		name: "张三",
		age:  20,
		sex:  "男",
	}
	//定义结构体指针指向变量的地址
	var p *Student
	//结构体指针指向结构体变量地址
	p = &stu
	(*p).name = "李四"
	fmt.Println(stu)
	//通过指针可以直接操作结构体成员
	p.name = "王五"
	fmt.Println(stu)
	//p1 := &stu
	//p2 := &stu.id

	//fmt.Printf("%T\n", p1) //*main.Student
	//fmt.Printf("%T\n", p2) //*int
	//fmt.Printf("%T\n", p)
	//结构体变量的地址是结构体中第一个元素的地址
	//fmt.Printf("%p\n", &stu)
	//fmt.Printf("%p\n", &stu.id)
}
