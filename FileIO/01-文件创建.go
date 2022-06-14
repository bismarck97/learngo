package main

import (
	"fmt"
	"os"
)

func main01() {
	fp, err := os.Create("./a.txt")
	//文件创建失败
	if err != nil {
		fmt.Println(err)
		return //如果return出现在主函数中 表示程序的结束
	}
	//延迟调用，关闭文件
	//1.占有内存和缓冲区
	//2.文件打开上限 65535
	defer fp.Close()

	fmt.Println("文件创建成功")

}
