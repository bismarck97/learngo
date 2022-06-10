package main

import (
	"fmt"
	"time"
)

func main04() {
	now := time.Now()
	//Add 原时间上加一个小时
	t1 := now.Add(time.Hour)
	fmt.Println(now)
	fmt.Println(t1)
	//Sub 求两个时间之间的差值
	fmt.Println(t1.Sub(now))
	//Equal 判断两个时间是否相同
	fmt.Println(t1.Equal(now))
	//Before 判断一个时间点是否在另一个时间点之前 返回真，否则返回假
	fmt.Println(t1.Before(now))
	//After 判断一个时间点是否在另一个时间点之后 返回真，否则返回假
	fmt.Println(t1.After(now))
}
