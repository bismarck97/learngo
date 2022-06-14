package main

import (
	"bufio"
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

func main0301() {
	//只读方式打开文件
	fp, err := os.Open("./a.txt")
	//1.文件不存在
	//2.文件权限
	//3.文件打开上限
	if err != nil {
		fmt.Println(err)
		return
	}
	b := make([]byte, 1024)
	//换行也会作为字符的一部分进行读取
	fp.Read(b)

	//fmt.Println(b)
	//fmt.Printf("%s", b)
	fmt.Println(string(b))
}

func main0302() {
	readAll()
}
func main0303() {
	fp, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建切片缓冲区
	r := bufio.NewReader(fp)
	//读取一行内容
	b, _ := r.ReadBytes('\n')

	//打印切片中字符的ASCII值
	//fmt.Println(b)

	//将切片转成string打印 汉字
	fmt.Print(string(b))
	//如果在文件截取中 没有标志位(分隔符)到文件末尾自动停止 EOF -1 文件结束标准
	b, _ = r.ReadBytes('\n')

	fmt.Println(string(b))
	defer fp.Close()
}
