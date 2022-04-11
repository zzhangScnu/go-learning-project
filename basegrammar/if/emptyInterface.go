package _if

import "fmt"

// Judge .If的语法，if 表达式... 这个表达式里可以赋值，只要最后给出true / false值
func Judge(i interface{}) {
	// 这里的t或ok都在各自if分支的作用域里，所以可以重名
	if t, ok := i.(int); ok {
		fmt.Println("type is int: ", t)
	} else if t, ok := i.(float64); ok {
		fmt.Println("type is float ", t)
	} else {
		fmt.Println("type is other ", t)
	}
}
