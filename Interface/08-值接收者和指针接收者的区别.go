package main

import "fmt"

type mover interface {
	move()
}
type sayer interface {
	say()
}

//接口继承
type animal interface {
	mover
	sayer
}

type per struct {
	name string
	age  int
}

//使用值接收者实现接口：类型的值和类型的指针都能保存倒接口变量中
//func (p per) move() {
//	fmt.Printf("%s在跑...\n", p.name)
//}
//使用指针接收者实现接口：只有类型指针能够保存倒接口变量中
func (p *per) move() {
	fmt.Printf("%s在跑...\n", p.name)
}

func (p *per) say() {
	fmt.Printf("%s在叫...\n", p.name)
}
func main08() {
	var m mover
	//p1 := per{ //p1是per类型的值
	//	name: "张三",
	//	age:  18,
	//}
	p2 := &per{ //p2是per类型的指针
		name: "李四",
		age:  20,
	}
	//m = p1 //无法赋值给指针接收者，因为p1是per类型的值，没有实现接口
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer //定义一个sayer类型的变量
	s = p2
	s.say()
	fmt.Println(s)

	var a animal //定义了一个animal类型的变量
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}
