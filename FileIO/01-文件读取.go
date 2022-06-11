package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	fileObj, err := os.Open("./xx.txt") //相对路径，相对执行的程序同目录下的xx.txt
	if err != nil {
		fmt.Printf("open file failed,error%v\n", err)
		return
	}
	//读取文件的内容
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("open file failed,error%v\n", err)
	}
	fmt.Printf("read %d bytes form file.\n", n)
	fmt.Println(string(tmp))
	//延迟执行，到最后关闭文件
	defer fileObj.Close() //关闭文件
}
