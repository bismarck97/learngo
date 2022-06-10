package main

import (
	"fmt"
	"time"
)

func main03() {
	fmt.Println("start...")
	//时间间隔
	n := 5 //type int
	//让程序暂停5秒
	time.Sleep(time.Duration(n) * time.Second) //纳秒  使用int类型要显式转换成Duration类型
	fmt.Println("over...")
}
