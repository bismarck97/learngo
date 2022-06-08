package main

import "fmt"

func main0101() {
	//切片定义 var 切片名 []数据类型
	//var s []int
	//fmt.Println(s)
	//自动推导类型创建切片 make([]数据类型，长度)
	s := make([]int, 5)
	//通过下标为切片赋值
	s[0] = 123
	s[1] = 234
	s[2] = 345
	s[3] = 456
	s[4] = 567
	//s[6] = 789//err.
	//通过append添加元素
	s = append(s, 678, 789, 890)
	fmt.Println(s)
	//切片的长度
	fmt.Println(len(s))
}
func main0102() {
	s := make([]int, 5)
	s[0] = 123
	s[1] = 234
	s[2] = 345
	s[3] = 456
	s[4] = 567

	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i])
	//}
	for i, v := range s {
		fmt.Println(i, v)
	}
}
func main() {
	//不写元素个数叫切片，必须写元素个数的叫数组
	var s []int = []int{1, 2, 3, 4, 5}
	for i, i2 := range s {
		fmt.Println(i, i2)
	}
}
