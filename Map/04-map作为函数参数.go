package main

import "fmt"

//map作为函数参数是引用传递
func test(m map[int]string) {
	delete(m, 1)
}
func main() {
	m1 := map[int]string{1: "hello", 2: "world", 3: "go"}
	test(m1)
	//形参可以影响实参
	fmt.Println(m1)
}
