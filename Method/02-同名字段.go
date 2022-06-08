package main

import "fmt"

//采用就近原则 使用子类信息
func main02() {
	var teacher Teacher = Teacher{

		Person: Person{"张三", 18, "男"},
		Id:     101,
		Name:   "张老师",
		Course: "Go语言",
	}

	//teacher.Id = 100
	//teacher.Name = "张老师"
	////要采用父类的参数
	//teacher.Person.Name = "张三"
	//teacher.Age = 18
	//teacher.Course = "Go语言"

	fmt.Println(teacher)
}
