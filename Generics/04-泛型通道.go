package main

import "fmt"

// C 类型别名
type C[T any] chan T

func main04() {
	c1 := make(C[int], 10)
	c1 <- 1
	c1 <- 2

	c2 := make(C[string], 10)
	c2 <- "hello"
	c2 <- "world"
	fmt.Println(<-c1, <-c2)
}
