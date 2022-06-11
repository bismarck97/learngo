package main

import (
	"fmt"
	"reflect"
)

//通过反射获得类型
func reflectType(x interface{}) {
	//不知道别人调用我这个函数的时候会传进来什么类型的变量
	//1.通过类型断言  问题：只能一个个的猜
	//2.借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj) //*reflect.rtype

}

type Cat struct {
}
type Dog struct {
}

func main01() {
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 10
	reflectType(b)
	var c string = "张三"
	reflectType(c)

	//结构体类型
	var d Cat
	reflectType(d)
	var e Dog
	reflectType(e)
	//slice
	var f []int
	reflectType(f)
	var g []string
	reflectType(g)
}
