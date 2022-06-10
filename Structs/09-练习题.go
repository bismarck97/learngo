package main

import (
	"fmt"
)

//使用“面向对象”的思维方式编写一个学生信息管理系统。
//学生有id、姓名、年龄、分数等信息
//程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
func main() {
	stus := newStuMgr()
	//初始化的信息
	stus.initStu()
	fmt.Println("欢迎来到学生管理系统")
	for {
		fmt.Println("请输入相关命令操作 1:展示所有学生信息  2:添加学生  3:编辑学生信息  4:删除学生  0:退出学生管理系统")
		var input int
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			//1:展示所有学生信息
			stus.ShowStudent()
		case 2:
			// 2:添加学生
			stus.addStudent(addStu())
		case 3:
			//3:编辑学生信息
			s := findStu(stus, stus.allStu)
			stus.modStudent(s)
		case 4:
			//4:删除学生
			s := stus.SearchStu()
			stus.delStudent(s)
		case 0:
			//退出系统
			//os.Exit(0)
			goto breakSwitch
		default:
			fmt.Println("输入操作有误，请重新输入:")
			break
		}
	}
breakSwitch:
	fmt.Println("欢迎再次使用学生管理系统")
}
