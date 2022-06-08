package main

import "fmt"

func main() {
	//var arr []interface{}
	arr := make([]interface{}, 3)
	arr[0] = 123
	arr[1] = 3.1456
	arr[2] = "hello"

	for _, v := range arr {
		//对数据v进行类型断言
		//如果类型断言成功，则v的类型为int data是数据,ok为true或者false

		data, ok := v.(int)
		if ok {
			fmt.Println(data)
		}
		if value, ok := v.(int); ok {
			fmt.Println("整型数据:", value)
		} else if value, ok := v.(float64); ok {
			fmt.Println("浮点型数据:", value)
		} else if value, ok := v.(string); ok {
			fmt.Println("字符串数据:", value)
		}

	}
}
