package function

import "fmt"

// step 1
func convert() int {
	fmt.Println("convert value")
	return 1
}

var v = convert()

// step 2
// 所有go文件中声明的init函数，都会在服务启动的时候被调用
func init() {
	fmt.Println("initialize value")
	v = 2
}

func Show() {
	fmt.Println("final value is ", v)
}
