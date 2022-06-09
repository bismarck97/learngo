package main

import "fmt"

//pannic和recover
func A() {
	fmt.Println("func A")
}
func B() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func B error")
		}
	}()
	panic("panic in B")
}
func C() {
	fmt.Println("func C")
}

func main10() {
	A()
	B()
	C()
}
