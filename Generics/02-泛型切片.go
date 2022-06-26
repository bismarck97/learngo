package main

type vector[T any] []T

func main() {
	v1 := vector[int]{50, 60}
	PrintSlice(v1)
	v2 := vector[string]{"张三", "李四", "王五"}
	PrintSlice(v2)
}
