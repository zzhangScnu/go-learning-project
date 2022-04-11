package main

import (
	"github.com/go-learning-project/interface/combination"
	"testing"
)

func TestRaiseUp(t *testing.T) {
	// 无论是interface还是struct，都可以用组合的方式嵌入其他的interface或struct
	// 作用域提升，直接调用成员变量或方法
	student := combination.Student{Course: "Go Development"}
	student.Say()
	programmer := combination.Programmer{Language: "Go"}
	programmer.Say()
}
