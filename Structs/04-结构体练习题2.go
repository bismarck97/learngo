package main

import "fmt"

func test(stu Student) {
	stu.Id = 666
	fmt.Println("stu", stu)
}

func Avg(student []Student) float64 {
	var avg float64

	for i := 0; i < len(student); i++ {
		avg += student[i].Score
	}
	return avg / float64(len(student))
}
func main04() {
	//s := Student{Id: 101}
	//fmt.Println(s)
	//平均成绩
	fmt.Println(Avg(Students()))
}
