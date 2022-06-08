package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}
type student struct {
	person
	score int
}
type teacher struct {
	person
	subject string
}

func (s *student) SayHello() {
	fmt.Printf("大家好，我是%s,我是%s生，我今年%d岁了,我的成绩是%d分\n", s.name, s.sex, s.age, s.score)
}
func (t *teacher) SayHello() {
	fmt.Printf("大家好，我是%s,我是%s生，我今年%d岁了,我是教%s的老师\n", t.name, t.sex, t.age, t.subject)
}

//1
//接口的定义和使用
//接口定义了规则  方法实现了规则
//接口是虚的 方法是实的
//格式 type 接口名 interface {方法列表}
//方法名(参数列表)(返回值列表)
type humaner interface {

	// SayHello  方法 函数的声明没有具体的实现 具体的实现要根据对象不同 实现方式也不同 2
	//接口中的方法必须有具体的实现
	SayHello()
}

func main02() {
	var stu student = student{
		person: person{
			name: "张三",
			age:  18,
			sex:  "男",
		},
		score: 100,
	}
	var tea teacher = teacher{
		person: person{
			name: "李四",
			age:  22,
			sex:  "男",
		},
		subject: "语文",
	}
	//3
	//定义接口类型
	//接口做了统一的处理，先实现接口 再根据接口实现对应的方法
	//在需求改变时 不需要改变接口 只需要修改方法
	var h humaner
	h = &stu
	h.SayHello()

	h = &tea
	h.SayHello()

	//stu.SayHello()
	//tea.SayHello()
}
