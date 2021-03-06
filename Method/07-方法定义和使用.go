package main

import "fmt"

func main0701() {
	var stu = Student{
		Person: Person{
			Name: "张三",
			Age:  19,
			Sex:  "女",
		},
		Id:    101,
		Score: 100,
	}
	//对象.方法
	stu.PrintInfo()
	//结构体变量可以直接使用结构体指针对应的方法
	stu.EditInfo("李四", 25, "男")
	fmt.Println("==========================")
	stu.PrintInfo()
}
func main() {
	p1 := NewPerson("张三", "男", 18)
	//(*p1).SayHi()
	//go语言的语法糖自动把指针类型转换成变量
	p1.SayHi()

	//调用修改年龄的方法
	fmt.Println(p1.Age)
	p1.SetAge(20)
	fmt.Println(p1.Age)
	//不会修改值 值传递只是拷贝操作
	p1.SetAge2(30)
	fmt.Println(p1.Age)
}
