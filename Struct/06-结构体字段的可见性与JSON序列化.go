package main

//结构体字段的可见性与JSON序列化
type student struct {
	Id   int
	Name string
}
type class struct {
	Title    string
	Students []student
}

func main() {

}
