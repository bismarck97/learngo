package main

//在定义结构体成员时，不要加var
type Student struct {
	Id    int
	Name  string
	Sex   string
	Age   int
	Addr  string
	Score float64
}

func Students() []Student {
	student := []Student{
		{
			Id:    101,
			Name:  "张三",
			Sex:   "男",
			Age:   18,
			Addr:  "成都",
			Score: 90,
		},
		{
			Id:    102,
			Name:  "李四",
			Sex:   "女",
			Age:   20,
			Addr:  "北京",
			Score: 80,
		},
		{
			Id:    103,
			Name:  "王五",
			Sex:   "男",
			Age:   18,
			Addr:  "上海",
			Score: 100,
		},
	}
	return student
}
