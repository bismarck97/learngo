package main

import "fmt"

//结构体字段的可见性与JSON序列化
//如果一个Go语言包中定义的标识符是首字母大写的，那么就是对外可见的
type student struct {
	Id   int
	Name string
}

//构造函数
func newStudent(id int, name string) student {
	return student{
		Id:   id,
		Name: name,
	}
}

type class struct {
	Title    string
	Students []student
}

func main() {
	//创建一个班级变量c1

	c1 := class{
		Title:    "班级一",
		Students: make([]student, 0, 20),
	}
	//往班级c1中添加学生
	fmt.Println(c1)
	for i := 0; i < 10; i++ {
		//造10个学生
		tmpStu := newStudent(i, fmt.Sprintf("学生%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)
}
