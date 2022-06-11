package main

import (
	"fmt"
	"reflect"
)

//Field(i int) StructField	根据索引，返回索引对应的结构体字段的信息。
//NumField() int	返回结构体成员字段数量。
//FieldByName(name string) (StructField, bool)	根据给定字符串返回字符串对应的结构体字段的信息。
//FieldByIndex(index []int) StructField	多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
//FieldByNameFunc(match func(string) bool) (StructField,bool)	根据传入的匹配函数匹配需要的字段。
//NumMethod() int	返回该类型的方法集中方法的数目
//Method(int) Method	返回该类型方法集中的第i个方法
//MethodByName(string)(Method, bool)	根据方法名返回该类型方法集中的方法

type Student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

// Study 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s Student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s Student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

//根据反射去获取结构体中方法的函数
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.NumMethod()) //取到变量的方法数量

	v := reflect.ValueOf(x)
	//因为下面需要拿到具体的方法去调用，所以使用的是值信息
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args) //调用方法
	}
	//通过方法名获取结构体的方法
	t.MethodByName("Sleep")              //(method,bool)
	methodObj := v.MethodByName("Sleep") //value
	fmt.Println(methodObj.Type())
}

//结构体反射
func main0401() {
	stu1 := Student{
		Name:  "张三",
		Score: 100,
	}
	//通过反射去获取结构体中的所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())

	//t.NumField() //返回结构体字段数量
	//遍历结构体遍历的所有字段
	for i := 0; i < t.NumField(); i++ {
		//根据结构体字段的索引去取字段
		filedObj1 := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v\n", filedObj1.Name, filedObj1.Type, filedObj1.Tag)
		fmt.Println(filedObj1.Tag.Get("json"), filedObj1.Tag.Get("ini"))
	}
	//根据名字取取结构体中的字段
	filedObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v\n", filedObj2.Name, filedObj2.Type, filedObj2.Tag)
	}
}
func main() {
	stu1 := Student{
		Name:  "张三",
		Score: 90,
	}
	printMethod(stu1)
}
