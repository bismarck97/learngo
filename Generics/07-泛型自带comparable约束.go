package main

import (
	"fmt"
)

func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
	fmt.Println(findFunc([]string{"烤鸡", "烤鸭", "烤鱼", "烤面筋"}, "烤面筋"))
}
