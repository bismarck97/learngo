package main

import "fmt"

func hello() {
	fmt.Println("hello world")
}
func main() { //开启一个主goroutine去指向main函数
	go hello() //开启了一个goroutine去指向hello这个函数
	fmt.Println("hello main")
}
