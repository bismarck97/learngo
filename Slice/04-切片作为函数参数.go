package main

import "fmt"

func InitData(num []int) {
	//切片作为函数参数是地址传递，形参可以改变实参的值
	for i := 0; i < len(num); i++ {
		num[i] = i
	}
}
func main() {
	//创建一个切片
	num := make([]int, 10)
	InitData(num)
	//打印切片
	for _, i2 := range num {
		fmt.Println(i2)
	}
}
