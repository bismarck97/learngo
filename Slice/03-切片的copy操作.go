package main

import "fmt"

func main() {
	//copy(切片1，切片2)
	srcSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6, 6, 6}
	//直接将srcSlice切片的值赋给dstSlice元素中相同的位置，而dstSlice的元素被替换掉
	//copy(dstSlice, srcSlice)
	//fmt.Println("dstSlice:", dstSlice)输出：//dstSlice: [1 2 6 6 6 6 6]

	copy(srcSlice, dstSlice)
	fmt.Println("srcSlice:", srcSlice) //srcSlice: [6 6]

	//在进行拷贝时,拷贝的长度为两个slice中长度较小的长度值

}
