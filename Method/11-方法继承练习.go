package main

import "fmt"

//练习1：根据以下信息，实现对应的继承关系
//记者：我是记者  我的爱好是偷拍 我的年龄是34 我是一个男狗仔
//程序员：我叫孙权 我的年龄是23 我是男生 我的工作年限是 3年
type person struct {
	name string
	age  int
	sex  string
}
type journalist struct {
	person
	hobby string
}
type Programmer struct {
	person
	worktime string
}

func (p person) sayHi() {
	fmt.Printf("大家好，我是%s,我是%s生，我今年%d岁了,", p.name, p.sex, p.age)
}
func (j journalist) sayHi() {
	j.person.sayHi()
	fmt.Println("我的爱好是", j.hobby)
}
func (p Programmer) saiHi() {
	p.person.sayHi()
	fmt.Println("我的工作年限是", p.worktime, "年")
}
func main11() {
	var jour journalist = journalist{
		person: person{
			name: "男狗仔",
			age:  18,
			sex:  "男",
		},
		hobby: "偷拍",
	}
	jour.sayHi()

	var pro Programmer = Programmer{
		person: person{
			name: "孙权",
			age:  25,
			sex:  "女",
		},
		worktime: "3",
	}
	pro.saiHi()
}
