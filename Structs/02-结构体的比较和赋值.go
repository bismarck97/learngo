package main

import "fmt"

func main02() {
	//结构体的比较和赋值
	//两个结构体可以使用 == 或 != 运算符进行比较，但不支持 > 或 <
	//在 Go 语言中，Go 结构体有时候并不能直接比较
	//当其基本类型包含：slice、map、function 时，是不能比较的。
	//若强行比较，就会导致直接报错的情况。

	s1 := Student{
		Id:   101,
		Name: "张三",
		Sex:  "男",
		Age:  18,
		Addr: "成都",
	}
	s2 := Student{
		Id:   101,
		Name: "张三",
		Sex:  "男",
		Age:  18,
		Addr: "成都",
	}
	s3 := Student{
		Id:   101,
		Name: "李四",
		Sex:  "女",
		Age:  20,
		Addr: "上海",
	}

	fmt.Println("s1==s2", s1 == s2)
	fmt.Println("s1==s3", s1 == s3)

	//同类型的两个结构体变量可以互相赋值
	var tmp Student
	tmp = s3
	fmt.Println("tmp = ", tmp)
}
