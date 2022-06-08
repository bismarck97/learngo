package main

import "fmt"

type programmer struct {
	*Person
	id   int
	post string
}

func main03() {
	var pro programmer = programmer{
		id:   101,
		post: "程序员",
	}
	var per = Person{"张三", 18, "男"}
	//pro.post = "程序员"
	//将结构体变量 赋值给结构体指针类型
	pro.Person = &per
	//pro.Person = new(Person)
	//error: invalid memory address or nil pointer dereference

	//pro.Name = "张三" //会被覆盖掉
	//pro.Person.Name = "张先生"
	//pro.Age = 18
	//pro.id = 1

	fmt.Println(pro)
	fmt.Println(pro.Name)
	//fmt.Println(pro.Person.Name)
	fmt.Println(pro.Age)
	fmt.Println(pro.Sex)
	fmt.Println(pro.id)
	fmt.Println(pro.post)
}
