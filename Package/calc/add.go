package calc

import (
	"fmt"
	"learngo/package/snow"
)

//标识符首字母大写表示对外可见
//通常不对外的标识符都是要首字母小写
//类似Java的public 和private
var name = "张三"

// Add 是一个计算两个int类型数据和的函数
func Add(a, b int) int {
	snow.Snow()
	Sub(a, b) //同一个保重的不同文件可以直接调用
	return a + b
}

//init函数在包导入的时候自动执行,优先于main函数
//init函数没有参数也没有返回值
func init() {
	fmt.Println("calc.init()")
}
