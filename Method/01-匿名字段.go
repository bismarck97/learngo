package main

import "fmt"

func main01() {
	var stu Student = Student{Person{"张三", 20, "女"}, 101, 100}

	//stu.Id = 101
	//stu.Name = "张三"
	//stu.Age = 18
	//stu.Score = 100
	//stu.Sex = "女"

	fmt.Println(stu)
}
