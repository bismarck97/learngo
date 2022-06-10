package main

import (
	"fmt"
)

type stu struct {
	id          int
	name, class string
}

//newStu是stu类型的构造函数
func newStu(id int, name, class string) *stu {
	return &stu{id: id, name: name, class: class}
}

//学员管理的类型
type stuMgr struct {
	allStu []*stu
}

//stuMgr的构造函数，创建一个学员切片
func newStuMgr() *stuMgr {
	return &stuMgr{allStu: make([]*stu, 0, 100)}
}

//初始化学生信息
func (s *stuMgr) initStu() {
	s1 := newStu(101, "张三", "1班")
	s2 := newStu(102, "李四", "2班")
	s.addStudent(s1)
	s.addStudent(s2)
}

// ShowStudent 1.展示所有学生信息
func (s *stuMgr) ShowStudent() {
	for _, v := range s.allStu {
		fmt.Printf("学号:%d 名字:%s 班级:%s\n", v.id, v.name, v.class)
	}
}

//2.添加学生
func (s *stuMgr) addStudent(newStu *stu) {
	s.allStu = append(s.allStu, newStu)
}

//3.编辑学生信息
func (s *stuMgr) modStudent(newStu *stu) {
	for i := 0; i < len(s.allStu); i++ {
		if newStu.id == s.allStu[i].id {
			s.allStu[i] = newStu //根据切片的索引直接把新学生的信息赋值
			fmt.Println("学生信息修改成功")
			return
		}
	}
}

//4.删除学生信息
func (s *stuMgr) delStudent(newStu *stu) {
	for i := 0; i < len(s.allStu); i++ {
		if newStu.id == s.allStu[i].id {
			//删除对应ID索引的元素
			s.allStu = append(s.allStu[:i], s.allStu[(i+1):]...)
			fmt.Println("删除成功")
			return
		}
	}
	fmt.Println("没有找到对应的信息")
}

//获取用户输入的学员信息
func addStu() *stu {
	var (
		id          int
		name, class string
	)
	fmt.Println("请输入学生学号")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入学生姓名")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学生班级")
	fmt.Scanf("%s\n", &class)
	return newStu(id, name, class)
}

// SearchStu 根据id查找学生信息
func (s *stuMgr) SearchStu() *stu {
	fmt.Println("请输入学员ID:")
	var id int
	fmt.Scanf("%d\n", &id)
	for i := 0; i < len(s.allStu); i++ {
		if id == s.allStu[i].id {
			return s.allStu[i]
		}
	}
	//如果没找到返回空的学生信息
	return newStu(0, "", "")
}

//查找学生信息并返回新的学生信息
func findStu(stus *stuMgr, allStu []*stu) *stu {
	for {
		s := stus.SearchStu()
		if s.id == 0 {
			fmt.Println("输入ID信息错误，请重新输入")
		} else {
			goto returnStu
		}
	}
returnStu:
	return addStu()

}
