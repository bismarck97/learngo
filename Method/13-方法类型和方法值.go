package main

import (
	"fmt"
)

type funcdemo func()

func main() {
	//对象
	s := Student{
		Person: Person{
			Name: "张三",
			Age:  18,
			Sex:  "男",
		},
		Id:    101,
		Score: 100,
	}
	//代码区
	//方法类型（函数类型） 变量
	f := s.SayHi
	f()
	fmt.Printf("%T\n", f)
	fmt.Println(f)
	fmt.Println("===================")

	f = s.Person.SayHi
	fmt.Printf("%T\n", f)
	fmt.Println(f)
	fmt.Println("===================")

	var f1 funcdemo
	f1 = s.SayHi
	fmt.Printf("%T\n", f1)
	fmt.Println(f1)
}
