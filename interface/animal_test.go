package main

import (
	"fmt"
	"github.com/go-learning-project/interface/cat"
	"testing"
)

func TestAnimal(t *testing.T) {
	wantedColor := "yellow"
	oneCat := cat.NewCat(wantedColor, 10.0, 10.0)
	if oneCat.Color() != wantedColor {
		t.Errorf("test failed. result is: %s", oneCat.Color())
	}
}

func TestSetter(t *testing.T) {
	catAddr := cat.NewCatAddr("black", 10, 10)
	catAddr.SetColor("white")
	fmt.Println(catAddr)
	// 指针会隐式转换为值，去调用set方法。但是只是设置到了拷贝的入参里
	catAddr.SetColorForValue("yellow")
	fmt.Println(catAddr)
}

func TestSetterForValue(t *testing.T) {
	catVal := cat.NewCat("black", 10, 10)
	// 值会隐式转换为指针，去调用set方法。由于方法接收器是指针，设置到了对应的内存地址上
	catVal.SetColor("white")
	fmt.Println(catVal)
	catVal.SetColorForValue("yellow")
	fmt.Println(catVal)
}

func TestPointer(t *testing.T) {
	catAddr := cat.NewCatAddr("black", 10, 10)
	fmt.Println(catAddr)  // &值
	fmt.Println(&catAddr) // 内存地址
	fmt.Println(*catAddr) // 值
	num := 5
	fmt.Println(&num) // 内存地址
}
