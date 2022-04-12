package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("i am in a goroutine, I produced value")
		c <- true
	}()
	// 阻塞在获取channel值的语句上
	fmt.Println("i am in main, I consumed value ", <-c)
}
