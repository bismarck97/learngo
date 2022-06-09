package main

type Person struct {
	name, sex string
	age       int
}

//结构体是值类型，利用构造函数返回指针类型内存开销会小很多
func newPerson(name, sex string, age int) *Person {
	return &Person{
		name: name,
		sex:  sex,
		age:  age,
	}
}
