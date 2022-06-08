package main

func main10() {
	var stu Student = Student{
		Person: Person{
			Name: "张三",
			Age:  18,
			Sex:  "男",
		},
		Id:    101,
		Score: 100,
	}
	//子类继承父类结构体，允许使用父类结构体成员 允许使用父类的方法
	stu.SayHi()
}
