package main

func main012() {
	s := Student{
		Person: Person{
			Name: "张三",
			Age:  18,
			Sex:  "男",
		},
		Id:    101,
		Score: 100,
	}
	//默认使用子类的方法 采用就近原则
	//调用子类方法
	s.SayHi()
	//调用父类方法
	s.Person.SayHi()
}
