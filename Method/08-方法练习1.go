package main

import "fmt"

//定义一个学生类,有六个属性,分别为姓名、性别、年龄、语文、数学、英语成绩。

type Stu struct {
	Name    string
	Sex     string
	Age     int
	Chinese float64
	Math    float64
	English float64
}

//一个打招呼的方法：介绍自己叫XX，今年几岁了。是男同学还是女同学。
func (s *Stu) sayHi(name, sex string, age int) {
	s.Name = name
	s.Age = age
	s.Sex = sex
	//对年龄和性别进行判断
	if s.Age < 0 && s.Age > 100 {
		s.Age = 0
	}
	if s.Sex != "男" && s.Sex != "女" {
		s.Sex = "男"
	}
	fmt.Printf("大家好，我叫%s,我今年%d岁了，我是一个%s生\n", s.Name, s.Age, s.Sex)
}

//两个计算自己总分数和平均分的方法。{显示:我叫XX,这次考试总成绩为X分,平均成绩为X分}
func (s *Stu) showScore(chinese, math, english float64) {
	s.Chinese = chinese
	s.Math = math
	s.English = english
	var sum float64 = s.Chinese + s.Math + s.English
	var ave float64 = sum / 3
	fmt.Printf("我叫%s,这次考试总成绩为%.1f分,平均成绩为%.1f分\n", s.Name, sum, ave)
}
func main08() {
	var stu Stu
	stu.sayHi("张三", "男", 19)
	stu.showScore(100, 80, 99)
}
