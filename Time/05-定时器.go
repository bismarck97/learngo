package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器
	//返回值是个通道
	for tmp := range time.Tick(time.Microsecond * 2) {
		fmt.Println(tmp)
	}
}
