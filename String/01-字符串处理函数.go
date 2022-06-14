package main

import (
	"fmt"
	"strings"
)

func main0101() {
	str := "HelloWorld"
	//在一个字符串中查找是否有另一个字符串出现，模糊查找，返回值为bool类型
	value := strings.Contains(str, "llo")
	if value {
		fmt.Println("找到")
	} else {
		fmt.Println("未找到")
	}
}
func main0102() {
	//字符串拼接
	s := []string{"1234", "2345", "3456", "5678"}
	//将一个字符串切片 拼接成一个字符串
	str := strings.Join(s, "-")
	fmt.Println(str)
}
func main0103() {
	str := "你好世界"
	//从一个字符串中查找相对应的另一个字符串，返回下标
	i := strings.Index(str, "世界")
	fmt.Println(i) //6,汉字一个占3个
}
func main0104() {
	str := "你瞅啥"
	//把一个字符串重复n次 n取值范围要大于等于0
	s := strings.Repeat(str, 3)
	fmt.Println(s)
}
func main0105() {
	str := "HelloWorld"
	//字符串替换元素，参数 字符串 被替换内容 替换内容 替换次数  替换次数小于0表示全部替换
	s := strings.Replace(str, "l", "w", -1) //屏蔽关键字
	fmt.Println(s)
}
func main0106() {
	str := "1234-2345-3456-5678"
	//字符串切割 返回值是字符串切片
	s := strings.Split(str, "-")
	fmt.Println(s)
}
func main0107() {
	str := "===hello   world==="
	//去掉前后指定的内容
	s := strings.Trim(str, "=")
	fmt.Println(s)
}
func main0108() {
	str := "   are u  ok  ?   "
	//去掉字符串的空格 并返回有效数据切片
	s := strings.Fields(str)
	fmt.Println(s)
}
