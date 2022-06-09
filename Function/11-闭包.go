package main

import (
	"fmt"
	"strings"
)

//闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
func adder1() func(int) int { //2.往上找函数

	var x int                //3.声明一个变量x   // 7.x=10  //11.30
	return func(y int) int { //5.赋值操作 y=10  //9.向其赋值  //12.继续向下执行......

		x += y
		return x //6.x=10 并且返回  //10.此时x中有原来的值，再加上传入的值 返回30
	}
}
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//两个函数返回值
func Calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main011() {
	//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效
	var f1 = adder1()   //1.声明一个adder变量
	fmt.Println(f1(10)) //4：从变量中找值
	fmt.Println(f1(20)) //8:继续找
	fmt.Println(f1(30))

	fmt.Println("===========================")

	var f2 = adder2(10)
	fmt.Println(f2(10))
	fmt.Println(f2(20))
	fmt.Println(f2(30))
	fmt.Println("================")

	jpgFunc := makeSuffixFunc(".jpg")
	//txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	//fmt.Println(txtFunc("test")) //test.txt
	f01, f02 := Calc(10)
	fmt.Println(f01(1), f02(2)) //11 9
	fmt.Println(f01(3), f02(4)) //12 8
	fmt.Println(f01(5), f02(6)) //13 7
}
