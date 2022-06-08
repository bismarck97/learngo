package main

import "fmt"

func main06() {
	//new创建切片指针
	var p *[]int
	fmt.Printf("%p\n", p)
	p = new([]int)
	fmt.Printf("%p\n", p)
	*p = append(*p, 1, 2, 3, 4, 5)
	for i, v := range *p {
		fmt.Println(i, v)
	}
}
