package main

import "fmt"

func main04() {
	//指针数组
	var p [2]*int
	var i int = 10
	var j int = 20
	p[0] = &i
	p[1] = &j
	fmt.Println(p[0], *p[0])
	fmt.Println(p[1], *p[1])
}
