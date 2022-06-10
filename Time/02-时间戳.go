package main

import (
	"fmt"
	"time"
)

func main02() {
	now := time.Now() //获取当前时间的时间对象
	//时间戳 1970.1.1  ->现在的秒数
	timeStamp1 := now.Unix()     //秒数
	timeStamp2 := now.UnixNano() //纳秒数
	fmt.Println(timeStamp1, timeStamp2)
	//将时间戳转换为具体的时间格式
	//1654858067 1654858067118054200
	t := time.Unix(1654858067, 0)
	fmt.Println(t)
}
