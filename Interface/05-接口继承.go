package main

import "fmt"

type Humaner interface { //子集
	SayHi()
}
type personer interface { //超集 一组子集的集合
	Humaner
	Sing(name string)
}
type stu struct {
	name string
	age  int
	sex  string
}

func (s *stu) SayHi() {
	fmt.Printf("大家好，我是%s,我今年%d岁了,我是%s生\n", s.name, s.age, s.sex)
}
func (s *stu) Sing(name string) {
	fmt.Printf("大家好,我叫%s,我给大家唱一首:%s\n", s.name, name)
}
func main05() {
	//创建对象
	var s stu = stu{
		name: "张三",
		age:  19,
		sex:  "女",
	}
	var h Humaner //子集
	h = &s
	h.SayHi()

	var p personer //超集
	p = &s
	p.Sing("天路")

	//将超集转成子集 反过来不允许
	h = p
	h.SayHi()
}
