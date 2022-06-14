package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//文件拷贝 从一个文件中拷贝数据到新的文件中
	fp1, err1 := os.Open("./aaa.jpg")
	fp2, err2 := os.Create("./bbb.jpg")

	if err2 != nil || err1 != nil {
		fmt.Println("拷贝失败")
		return
	}
	defer fp1.Close()
	defer fp2.Close()

	b := make([]byte, 1024*1024)
	for {
		n, err := fp1.Read(b)
		//读到最后一行结束拷贝
		if err == io.EOF {
			break
		}
		fp2.Write(b[:n])
	}
	fmt.Println("拷贝完成")
}
