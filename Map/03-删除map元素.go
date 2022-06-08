package main

import "fmt"

func main03() {
	//删除map元素
	m1 := map[int]string{1: "hello", 2: "world", 3: "Go"}
	//根据map中的key，删除map中的值
	delete(m1, 1) //删除key为1的内容
	fmt.Println(m1)
}
