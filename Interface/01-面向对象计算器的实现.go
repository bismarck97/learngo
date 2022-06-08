package main

import "fmt"

type Opt struct {
	num1 int
	num2 int
}
type AddOpt struct {
	Opt
}
type SubOpt struct {
	Opt
}

func (add *AddOpt) Operate() int {
	return add.num1 + add.num2
}
func (sub *SubOpt) Operate() int {
	return sub.num1 - sub.num2
}

//定义接口

type Opter interface {
	Operate() int
}

func main01() {
	//var add AddOpt
	//add.num1 = 10
	//add.num2 = 20
	//num1 := add.Operate()
	//fmt.Println(num1)
	//
	//var sub SubOpt
	//sub.num1 = 20
	//sub.num2 = 10
	//num2 := sub.Operate()
	//fmt.Println(num2)
	var o Opter
	o = &AddOpt{Opt{10, 20}}
	fmt.Println(o.Operate())
	o = &SubOpt{Opt{20, 10}}
	fmt.Println(o.Operate())
}
