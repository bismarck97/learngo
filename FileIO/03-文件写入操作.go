package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//os.O_WRONLY	只写
//os.O_CREATE	创建文件
//os.O_RDONLY	只读
//os.O_RDWR	读写
//os.O_TRUNC	清空
//perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。

//write
func write() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	str := "hello world"
	fileObj.Write([]byte(str))         //[]byte
	fileObj.WriteString("Hello World") //string
}
func writeByBufio() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("Go Java .NET") //将内容写入缓冲区
	writer.Flush()                     //将缓存中的内容写入文件
}
func writeByIoUtil() {
	str := "张三李四王五HelloWorldGoJavaC#C++"
	//直接把一个文本信息写入文件，返回值是错误信息
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
}
func main() {
	//write()
	//writeByBufio()
	writeByIoUtil()
}
