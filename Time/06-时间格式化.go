package main

import (
	"fmt"
	"time"
)

func main06() {
	now := time.Now()
	//Go语言中用2006 01 02 15 04 05 表示其他语言的YY mm dd HH MM SS
	//2006 1 2 3 4 5
	ret1 := now.Format("2006.01.02")
	fmt.Println(ret1)
	ret2 := now.Format("2006.01.02 15:04:05")
	fmt.Println(ret2)
	ret3 := now.Format("2006.01.02 03:04:05 PM")
	fmt.Println(ret3)
	ret4 := now.Format("2006.01.02 03:04:05.000 PM")
	fmt.Println(ret4)
	//解析字符串类型的时间
	timeStr := "2022/06/10 22:00:00"
	//1.拿到时区
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(loc)
		return
	}
	//2.根据时区去解析一个字符串格式的时间
	//timeObj, err := time.Parse("2006/01/02 15:04:05", timeStr)
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Printf("parse timeStr failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
}
