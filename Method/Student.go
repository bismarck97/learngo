package main

import "fmt"

//结构体嵌套结构体

type Student struct {
	//通过匿名字段实现继承操作
	Person //结构体名称作为结构体成员
	Id     int
	Score  int
}

//方法重写

func (s *Student) SayHi() {
	fmt.Printf("大家好，我是%s,我是%s生,我今年%d岁了,我的分数是%d分\n", s.Name, s.Sex, s.Age, s.Score)
}

//在一个对象中不能出现相同的方法名 方法的接收者 带* 和不带* 表示一个相同的对象
//func (s *Student) SayHi(name string) {
//	fmt.Printf("大家好，我是%s,我是%s生,我今年%d岁了,我的分数是%d分\n", s.Name, s.Sex, s.Age, s.Score)
//}

func (stu *Student) PrintInfo() {
	fmt.Println("Id", stu.Id)
	fmt.Println("Name:", stu.Name)
	fmt.Println("Age:", stu.Age)
	fmt.Println("Sex:", stu.Sex)
	fmt.Println("Score:", stu.Score)
}

//传递结构体指针才能改变原对象的值

func (stu *Student) EditInfo(name string, age int, sex string) {
	//结构体指针可以直接调用结构体成员
	stu.Name = name
	stu.Age = age
	stu.Sex = sex
}
