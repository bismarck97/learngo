package main

import "fmt"

func main02() {
	m1 := map[int]string{1: "hello", 2: "world", 3: "Go"}
	//遍历map
	for key, value := range m1 {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	fmt.Println("======================")

	//输出时可以进行判断
	v, ok := m1[1]
	if ok {
		fmt.Println("m[1] = ", v)
	} else {
		fmt.Println("key不存在")
	}

}
