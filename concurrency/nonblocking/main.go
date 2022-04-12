package main

import (
	"fmt"
)

// 不能写在test方法里，test方法貌似会等待所有goroutine结束
// 而main方法结束会暴力结束所有goroutine
func main() {
	c := make(chan bool)
	// 匿名函数，最后的()表示立即调用这个无参函数
	go func() {
		// 主方法结束得早，导致goroutine里的方法执行不完
		fmt.Println("i am in a goroutine, I consumed value ")
		// 相当于接收并丢弃
		<-c
	}()
	fmt.Println("i am in main, I produced value ")
	c <- true
}
