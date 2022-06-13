package main

import "fmt"

// Opt 父类
type Opt struct {
	num1 int
	num2 int
}

// AddOpt 加法子类
type AddOpt struct {
	Opt
}

// SubOpt 减法子类
type SubOpt struct {
	Opt
}

// Operate 加法的方法实现
func (add *AddOpt) Operate() int {
	return add.num1 + add.num2
}

// Operate 减法的方法实现
func (sub *SubOpt) Operate() int {
	return sub.num1 - sub.num2
}

//设计模式 对于面向对象基于M(模型)V(视图)C(控制器)的设计模式
//工厂设计模式

// OptFractory 对象
type OptFractory struct {
}

// 设计模式 对于面向对象基于M(模型)V(视图)C(控制器)的设计模式

//OptCalc 工厂设计模式
func (o *OptFractory) OptCalc(num1, num2 int, opt string) (value int) {
	//通过运算符 进行分类计算
	//通过接口进行统一处理
	var opter Opter
	//通过不同运算符进行判断
	switch opt {
	case "+":
		//创建加法对象
		var add AddOpt = AddOpt{Opt{num1, num2}}
		//value = add.Operate()
		opter = &add
	case "-":
		//创建减法对象
		var sub SubOpt = SubOpt{Opt{num1, num2}}
		//value = sub.Operate()
		opter = &sub
	}
	//value = opter.Operate()
	//通过多态来实现接口计算操作
	value = Fractory(opter)
	return
}

// Opter 定义接口
type Opter interface {
	Operate() int
}

//多态的实现
func Fractory(o Opter) (value int) {
	value = o.Operate()
	return
}
func main01() {
	//1.基于继承和方法
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

	//2.基于继承方法和接口的操作
	//var o Opter
	//o = &AddOpt{Opt{10, 20}}
	//fmt.Println(o.Operate())
	//o = &SubOpt{Opt{20, 10}}
	//fmt.Println(o.Operate())

	//3.基于继承 方法 接口 多态和设计模式
	var optf OptFractory
	num := optf.OptCalc(100, 200, "-")
	fmt.Println(num)
}
