package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//read by bufio
func readByBufio() {
	//打开文件
	fileObj, err := os.Open("./xx.txt") //相对路径，相对执行的程序同目录下的xx.txt
	if err != nil {
		fmt.Printf("open file failed,error%v\n", err)
		return
	}
	//延迟执行，到最后关闭文件
	defer fileObj.Close() //关闭文件
	reader := bufio.NewReader(fileObj)
	//用死循环来读取文件
	for {
		line, err := reader.ReadString('\n')
		//读取到最后直接return结束循环
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read filr by bufio failed,err%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

//ioutil
func readByIoutil() {
	content, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		fmt.Printf("read file by ioutil failed,err%v\n", err)
		return
	}
	fmt.Println(string(content))
}
func main02() {
	readByIoutil()
}
