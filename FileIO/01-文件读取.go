package main

import (
	"fmt"
	"io"
	"os"
)

func readFromFile() {
	//打开文件
	fileObj, err := os.Open("./xx.txt") //相对路径，相对执行的程序同目录下的xx.txt
	if err != nil {
		fmt.Printf("open file failed,error%v\n", err)
		return
	}
	//定义字节切片保存文件信息
	var tmp = make([]byte, 128)
	//读取文件的内容
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("open file failed,error%v\n", err)
	}
	fmt.Printf("read %d bytes form file.\n", n)
	fmt.Println(string(tmp))
	//延迟执行，到最后关闭文件
	defer fileObj.Close() //关闭文件
}
func readAll() {
	//打开文件
	fileObj, err := os.Open("./xx.txt") //相对路径，相对执行的程序同目录下的xx.txt
	if err != nil {
		fmt.Printf("open file failed,error%v\n", err)
		return
	}
	//读取文件的内容
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		//EOF表示End Of File
		if err == io.EOF {
			//把当前读了多少个字节的数据打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("open file failed,error%v\n", err)
			return
		}
		fmt.Printf("read %d bytes form file.\n", n)
		fmt.Println(string(tmp[:n]))

		//延迟执行，到最后关闭文件
		defer fileObj.Close() //关闭文件
		fmt.Println(string(tmp))
	}
}

func main01() {
	readAll()
}
