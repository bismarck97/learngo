package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

// NewPerson 构造函数
func NewPerson(name, sex string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
		Sex:  sex,
	}
}

//什么时候应该使用指针类型接收者
//需要修改接收者中的值
//接收者是拷贝代价比较大的大对象
//保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

// SetAge 是一个修改年龄的方法
//指针接收者指的就是接收者的类型是指针
func (p *Person) SetAge(newAge int) {
	p.Age = newAge
}

func (p Person) SetAge2(newAge int) {
	p.Age = newAge
}

func (p *Person) SayHi() {
	fmt.Printf("大家好，我是%s,我是%s生,我今年%d岁了\n", p.Name, p.Sex, p.Age)
}
