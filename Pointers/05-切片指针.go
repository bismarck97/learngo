package main

import "fmt"

func main05() {
	//切片指针
	var slice []int = []int{1, 2, 3, 4, 5}
	var p *[]int = &slice
	*p = append(*p, 6, 7, 8, 9, 10)

	fmt.Println(slice)
	fmt.Println(*p)
}
