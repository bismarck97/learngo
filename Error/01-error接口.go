package main

import (
	"errors"
	"fmt"
)

//编辑时异常
//编译时异常
//运行时异常
func test(a, b int) (value int, err error) {
	//0不能作为除数
	if b == 0 {
		err = errors.New("0不能作为除数")
		return
	} else {
		value = a / b
		return
	}
}
func main01() {
	value, err := test(10, 0)
	//err如果不等于空表示有错误信息
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

}
