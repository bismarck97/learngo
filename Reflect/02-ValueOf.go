package main

import (
	"fmt"
	"reflect"
)

//通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)

	fmt.Printf("%v,%T", v, v)
	k := v.Kind() //拿到值对应的类型种类
	fmt.Println(k)
	//如何得到一个传入时候类型的变量
	switch k {
	case reflect.Int32:
		//把反射取到的值转换成一个int32类型的变量
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
	}
}

//通过反射设置变量的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	//Elem()用来根据指针取对应的值
	k := v.Elem().Kind() //拿到指针对应的值
	switch k {           //再做类型种类判断
	case reflect.Int32:
		v.Elem().SetInt(20)
	case reflect.Float32:
		v.Elem().SetFloat(3.14)
	}

}
func main02() {
	//valueOf
	var a int32 = 100
	//reflectValue(a)
	reflectSetValue(&a)
	fmt.Println(a)
	var b float32 = 1.234
	//reflectValue(b)
	reflectSetValue(&b)
	fmt.Println(b)
}
