package main

import "fmt"

func main03() {
	//输出学生名字
	//s := Students()
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i].Name)
	//}

	student := Students()
	//练习1：计算以上学生成绩的总分
	var totalScore float64
	for i := 0; i < len(student); i++ {
		totalScore += student[i].Score
	}
	fmt.Println("学生总成绩为：", totalScore)

	//练习2：输出以上学生成绩中最高分
	maxScore := student[0].Score //先赋值第一个元素赋值
	//用循环进行比较
	for i := 0; i < len(student); i++ {
		if student[i].Score > maxScore {
			maxScore = student[i].Score
		}
	}
	fmt.Println("最高成绩为：", maxScore)
}
