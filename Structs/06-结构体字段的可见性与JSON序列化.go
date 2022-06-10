package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性与JSON序列化
//如果一个Go语言包中定义的标识符是首字母大写的，那么就是对外可见的
//如果一个结构体中的字段名首字母是大写的，那么该字段就是对外可见的
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
	Title    string    `json:"title"` //结构体标签（Tag）
	Students []student `json:"students_list" db:"student" xml:"ss"`
}

func main06() {
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
	//fmt.Printf("%#v\n", c1)
	//JSON序列化：Go语言中的数据->JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)
	//JSON反序列化：JSON格式的字符串->Go语言中的数据
	//jsonStr := `{"Title":"班级一","Students":[{"Id":0,"Name":"学生00"},{"Id":1,"Name":"学生01"},{"Id":2,"Name":"学生02"},{"Id":3,"Name":"学生03"},{"Id":4,"Name":"学生04"},{"Id":5,"Name":"学生05"},{"Id":6,"Name":"学生06"},{"Id":7,"Name":"学生07"},{"Id":8,"Name":"学生08"},{"Id":9,"Name":"学生09"}]}`
	jsonStr := `{"title":"班级一","Students":[{"Id":0,"Name":"学生00"},{"Id":1,"Name":"学生01"},{"Id":2,"Name":"学生02"},{"Id":3,"Name":"学生03"}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	fmt.Printf("%#v\n", c2)
}
