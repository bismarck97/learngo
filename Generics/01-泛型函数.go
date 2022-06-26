package main

import "fmt"

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}
func main01() {
	PrintSlice[int]([]int{10, 20, 30, 40, 50})
	PrintSlice[string]([]string{"张三", "李四", "王五"})
	PrintSlice[float64]([]float64{1.1, 1.2, 1.3})

	//省略显示类型
	PrintSlice([]int64{11, 22, 33, 44, 55})
}
