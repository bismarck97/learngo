package main

import "fmt"

//定义结构体 存储5名学生 三门成绩

type Stu struct {
	id    int
	name  string
	score [3]int //结构体成员为数组
}
type Stus struct {
	id    int
	name  string
	score int
}

//定义求学生平均值和总成绩
func avgScore(stu [3]Stu) ([3]Stus, [3]int) {
	var stuScore [3]Stus
	var scoreSum [3]int
	for i := 0; i < len(stu); i++ {
		stuScore[i].id = stu[i].id
		stuScore[i].name = stu[i].name
		//先算出每名学生总成绩
		for j := 0; j < len(stu[i].score); j++ {
			scoreSum[i] += stu[i].score[j]
		}
		//平均成绩
		stuScore[i].score = scoreSum[i] / 3
	}
	return stuScore, scoreSum
}

//练习：3个学生，每个学生3门课程，求其平均成绩和总成绩
func main05() {
	arr := [3]Stu{
		{
			id:    101,
			name:  "张三",
			score: [3]int{80, 90, 60},
		}, {
			id:    102,
			name:  "李四",
			score: [3]int{100, 99, 98},
		}, {
			id:    103,
			name:  "王五",
			score: [3]int{100, 80, 95},
		},
	}
	//平均成绩

	s, nums := avgScore(arr)
	for i := 0; i < len(s); i++ {
		fmt.Printf("编号:%d 姓名:%s  平均成绩为:%d 总成绩为%d\n", s[i].id, s[i].name, s[i].score, nums[i])
	}
}
