package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	//从0开始截取切片
	//slice := s[0:]
	//从2之前截取切片，包前不包后
	//slice := s[:2]
	//从2到5之间截取切片
	//slice := s[2:5]
	//从0到2之间截取切片，容量是5
	slice := s[0:2:5]
	fmt.Println(cap(slice))
	fmt.Println(slice)
	//s[n]	切片s中索引位置为n的项
	//s[:]	从切片s的索引位置0到len(s)-1处所获得的切片
	//s[low:]	从切片s的索引位置low到len(s)-1处所获得的切片
	//s[:high]	从切片s的索引位置0到high处所获得的切片，len=high
	//s[low:high]	从切片s的索引位置low到high处所获得的切片，len=high-low
	//s[low:high:max]	从切片s的索引位置low到high处所获得的切片，len=high-low，cap=max-low
	//len(s)	切片s的长度，总是<=cap(s)
	//cap(s)	切片s的容量，总是>=len(s)

}
