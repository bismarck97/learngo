package main

import (
	"fmt"
	"time"
)

func main01() {
	now := time.Now() //获取当前时间的时间对象
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println(year, month, day, hour, minute, second)
}
