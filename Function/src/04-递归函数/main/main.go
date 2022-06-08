package main

import (
	"day01/src/04-递归函数/function"
	"fmt"
)

func main() {
	function.Test(6)
	num := function.Factorial(5)
	fmt.Println(num)
}
