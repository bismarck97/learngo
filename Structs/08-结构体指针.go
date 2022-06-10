package main

import "fmt"

//结构体指针
type per struct {
	name, city string
	age        int
}

func main08() {

	var p2 = new(per)
	fmt.Printf("%T\n", p2)
	//go语言的语法糖会自动把指针类型转换成对应的变量
	//(*p2).name = "张三"
	p2.name = "张三"
	//(*p2).city = "北京"
	p2.city = "北京"
	//(*p2).age = 18
	p2.age = 18
	//fmt.Println(*p2)
	fmt.Printf("%#v\n", p2)

	//取结构体的地址进行实例化
	p3 := &per{"李四", "上海", 20}
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)
}
