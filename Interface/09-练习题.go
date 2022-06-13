package main

import "fmt"

//1.定义结构体
type Calc struct {
	num1, num2 int
}

//加法
type addCalc struct {
	Calc
}

//加法构造函数
func newAdd(num1, num2 int) *addCalc {
	return &addCalc{Calc{num1, num2}}
}

//减法
type subCalc struct {
	Calc
}

//减法构造函数
func newSub(num1, num2 int) *subCalc {
	return &subCalc{Calc{num1, num2}}
}

//乘法
type mulCalc struct {
	Calc
}

//乘法构造函数
func newMul(num1, num2 int) *mulCalc {
	return &mulCalc{Calc{num1, num2}}
}

//除法
type divCalc struct {
	Calc
}

//除法构造函数
func newDiv(num1, num2 int) *divCalc {
	return &divCalc{Calc{num1, num2}}
}

//取余
type modCalc struct {
	Calc
}

//取余构造函数
func newMod(num1, num2 int) *modCalc {
	return &modCalc{Calc{num1, num2}}
}

//2.定义接口
type operaer interface {
	operation() int
}

//3.多态实现

func opera(o operaer) int {
	return o.operation()
}

//4.方法绑定对象
func (add *addCalc) operation() int {
	return add.num1 + add.num2
}

func (sub *subCalc) operation() int {
	return sub.num1 - sub.num2
}

func (mul *mulCalc) operation() int {
	return mul.num1 * mul.num2
}
func (div *divCalc) operation() int {
	return div.num1 / div.num2
}

func (mod *modCalc) operation() int {
	return mod.num1 % mod.num2
}

//5.工厂模式
//5.1定义空结构体
type calcFractory struct {
}

//5.2实现工厂对象方法
func (calc *calcFractory) Calc(num1, num2 int, symbol string) int {
	//定义接口
	var opt operaer
	switch symbol {
	case "+":
		//add := addCalc{Calc{num1, num2}}
		//opt = &add
		//利用构造函数来创建加法对象
		opt = newAdd(num1, num2)
	case "-":
		//sub := subCalc{Calc{num1, num2}}
		//opt = &sub
		//利用构造函数来创建减法对象
		opt = newSub(num1, num2)
	case "*":
		//mul := mulCalc{Calc{num1, num2}}
		//opt = &mul
		//利用构造函数来创建乘法对象
		opt = newMul(num1, num2)
	case "/":
		//div := divCalc{Calc{num1, num2}}
		//opt = &div
		//利用构造函数来创建除法对象
		opt = newDiv(num1, num2)
	case "%":
		//mod := modCalc{Calc{num1, num2}}
		//opt = &mod
		//利用构造函数来创建取余对象
		opt = newMod(num1, num2)
	default:
		//如果输入其他的字符，返回0
		opt = newMul(0, 0)
	}
	//多态实现接口
	return opera(opt)
}

func main() {
	//定义工厂对象
	var calc calcFractory
	num := calc.Calc(200, 300, "-")
	fmt.Println(num)
}
