package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p *Person) SayHi() {
	fmt.Printf("大家好，我是%s,我是%s生,我今年%d岁了\n", p.Name, p.Sex, p.Age)
}
