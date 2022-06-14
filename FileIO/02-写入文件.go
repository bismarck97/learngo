package main

import (
	"fmt"
	"io"
	"os"
)

func main0201() {
	//os.Create 创建文件时 文件不存在会创建一个新文件 如果文件存在会覆盖源内容
	fp, err := os.Create("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	//写入文件
	//\n在文件中是换行  windows文本文件换行是\r\n
	fp.WriteString("HelloWorld\r\n")
	fp.WriteString("你好世界")
}
func main0202() {
	//创建文件
	fp, err := os.Create("./b.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件
	defer fp.Close()
	//1.将字符切片写入到文件中
	//b := []byte{'h', 'e', 'l', 'l', 'o'}
	////文件指针.write(字符切片)
	//n, _ := fp.Write(b)
	//fmt.Println(n)

	//2.将字符串转换成字符切片写入到文件中
	//str := "HelloWorld"
	str := "锄禾日当午"
	//字符串和字符切片允许相互转换
	b := []byte(str)
	n, _ := fp.Write(b)
	fmt.Println(n)
}
func main0203() {
	//fp, err := os.Create("./ab.txt")
	//打开文件
	//os.Open(文件名)只读方式打开文件
	//os.Open("./a.txt")
	//Openfile不能创建新文件
	//os.Openfile(文件名,打开模式,打开权限)
	fp, err := os.OpenFile("./a.txt", os.O_RDWR, 6)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	//通过字符串获取字符切片
	//b := []byte("锄禾日当午")
	////当使用writeAt进行指定位置插入数据时会依次覆盖源内容
	//fp.WriteAt(b, 0)

	//获取文件的字符格式
	//n, _ := fp.Seek(3, io.SeekStart)
	//负数是向左偏移，正数是向右偏移
	n, _ := fp.Seek(3, io.SeekEnd)
	fmt.Println(n)
	b := []byte("abcd")
	fp.WriteAt(b, n)
}
