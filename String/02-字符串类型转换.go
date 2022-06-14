package main

import (
	"fmt"
	"strconv"
)

func main0201() {
	//将其他类型转成字符串
	//1.将bool类型转成字符串
	//str := strconv.FormatBool(true)

	//2.将int类型转成字符串，第一个参数表示具体的数据，第二个数据表示具体的进制
	//str := strconv.FormatInt(120, 10)

	//3.将int类型直接转成10进制的字符串
	//str := strconv.Itoa(123)

	//4.将浮点型转换成字符串 参数为 数据 'f',保留小数个数，会四舍五入，位数(64,32)
	str := strconv.FormatFloat(1.234, 'f', 2, 64)
	fmt.Println(str)
}
func main0202() {
	//将字符串转换成其他内容
	//1.将字符串转成bool
	//str := "123false123"
	////只能将"flase"或"true"转成bool类型 有多余的数据 是无效的
	//b, err := strconv.ParseBool(str)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(b)
	//}

	//2.将字符串转成整型 数据要符合整数的规则
	//str := "123"
	//参数为：数据 进制 位数
	//value, _ := strconv.ParseInt(str, 10, 64)
	//fmt.Println(value)
	//直接按照10进制转换
	//value, _ := strconv.Atoi(str)
	//fmt.Println(value)

	//3.将字符串转成浮点型
	str := "1.2345"
	//参数为 数据 位数
	value, _ := strconv.ParseFloat(str, 64)
	fmt.Println(value)
}
func main() {
	b := make([]byte, 0, 1024)
	//将bool类型放在指定切片中
	b = strconv.AppendBool(b, false)
	//将整数类型放在指定切片中
	b = strconv.AppendInt(b, 123, 10)
	//将浮点类型放在指定切片中
	b = strconv.AppendFloat(b, 3.1415, 'f', 3, 64)
	//将字符串类型放在指定切片中
	b = strconv.AppendQuote(b, "张三")
	fmt.Println(string(b))
}
