package main

import "fmt"

func main0601() {
	//空接口，可以接收任意类型的数据
	//任意类型都实现了空接口-->空接口变量可以存储任意值
	var i interface{} //相当于万能指针
	i = 1
	fmt.Println(i)
	fmt.Printf("%p\n", &i)
	fmt.Println("==================")
	i = "hello"
	fmt.Println(i)
	fmt.Printf("%p\n", &i)
	fmt.Println("==================")
	var arr = [3]int{1, 2, 3}
	i = arr
	fmt.Println(i)
	fmt.Printf("%p\n", &i)
	fmt.Println("==================")
	a := 10
	i = &a
	fmt.Println(i)
	fmt.Printf("%p\n", &i)
	var m = make(map[string]interface{}, 100)
	m["name"] = "张三"
	m["age"] = 18
	m["hobby"] = []string{"蓝球", "game", "吃饭"}
}
func main0602() {
	//空接口切片
	var i []interface{}
	//i := make([]interface{}, 3)
	i = append(i, 1, "hello", [3]int{1, 2, 3})
	//for j := 0; j < len(i); j++ {
	//	fmt.Println(i[j])
	//}
	for index, value := range i {
		fmt.Println(index, ":", value)
	}
}
