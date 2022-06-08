package main

import (
	"fmt"
)

func main01() {
	//结构体初始化:指定成员初始化
	var s1 Student = Student{
		Id:   101,
		Name: "张三",
		Sex:  "男",
		Age:  18,
		Addr: "成都",
	}
	fmt.Println(s1)
	//结构体定义完成后，结构体成员的使用
	var s2 Student
	s2.Id = 102
	s2.Age = 19
	s2.Sex = "男"
	s2.Addr = "北京"

}
