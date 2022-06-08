package main

import "fmt"

func main09() {
	//结构体切片指针操作
	var stu []Student = make([]Student, 3)
	p := &stu //*[]Student //结构体切片指针
	//(*p) = append((*p), Student{
	//	id:   101,
	//	name: "张三",
	//	age:  20,
	//	sex:  "男",
	//},
	//)
	(*p)[0] = Student{
		id:   101,
		name: "张三",
		age:  19,
		sex:  "男",
	}
	(*p)[1] = Student{
		id:   102,
		name: "李四",
		age:  18,
		sex:  "女",
	}
	(*p)[2] = Student{
		id:   103,
		name: "王五",
		age:  18,
		sex:  "男",
	}
	//fmt.Printf("%T\n", p)
	//fmt.Println(stu)
	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}
}
