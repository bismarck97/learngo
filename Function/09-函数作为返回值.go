package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

//定义一个函数,它的返回值是一个函数和error信息
//返回值函数的返回值是int
func do(s string) (func(int, int) int, error) {
	switch s {
	//如果传入的参数是+,返回add函数和空的错误信息
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		//如果传入参数有误返回空函数，以及错误信息
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main09() {
	//定义函数，传入要操作的值，接收值所对应的函数以及错误信息
	cal, err := do("+")
	//判断是否有错误信息
	if err == nil {
		res := cal(100, 200)
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

	cal, err = do("-")
	if err == nil {
		res := cal(300, 100)
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
	cal, err = do("*")
	if err == nil {
		res := cal(300, 100)
		fmt.Println(res)
	} else {
		//打印错误信息
		fmt.Println(err)
	}
}
