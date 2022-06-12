package main

import (
	"fmt"
)

func main() {
	//var arr []interface{}

	arr := make([]interface{}, 3)
	arr[0] = 123
	arr[1] = 3.1456
	arr[2] = "hello"

	for _, v := range arr {
		//对数据v进行类型断言,返回值有两个，原数据和bool类型
		//如果类型断言成功，则v的类型为int data是数据,ok为true或者false

		//data, ok := v.(int)
		//if ok {
		//	fmt.Println(data)
		//}

		//switch t := v.(type) {
		//switch v.(type) {
		//case int:
		//	fmt.Println("整型数据", v)
		//case float64:
		//	fmt.Println("浮点型数据", v)
		//case string:
		//	fmt.Println("字符串数据", v)
		//}
		if value, ok := v.(int); ok {
			fmt.Println("整型数据:", value)
		} else if value, ok := v.(float64); ok {
			fmt.Println("浮点型数据:", value)
		} else if value, ok := v.(string); ok {
			fmt.Println("字符串数据:", value)
		} else if value, ok := v.([3]int); ok {
			fmt.Println("数组数据:", value)
		}
	}
}
