package main

import "fmt"

type DemoA struct {
	name string
	id   int
}
type DemoB struct {
	age int
	sex string
}
type DemoC struct {
	DemoA
	DemoB
	score int
}

func main05() {
	var s DemoC = DemoC{
		DemoA: DemoA{"张三", 101},
		DemoB: DemoB{17, "男"},
		score: 101,
	}
	//s.name = "张三"
	//s.score = 100
	//s.age = 18
	//s.id = 101
	//s.sex = "女"

	fmt.Println(s)
}
