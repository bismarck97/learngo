package main

import "fmt"

func main07() {
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
