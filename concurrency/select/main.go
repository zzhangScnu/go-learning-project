package main

import (
	"fmt"
	"time"
)

const (
	tickTime = time.Second
)

// 模拟一个带缓冲区的生产者-消费者模型
// 满足以下条件时会发送数据：
// 1. 缓冲区满
// 2. 缓冲区未满，但经过了一定的时间
func main() {
	c, d := make(chan int), make(chan bool)
	defer close(c)
	defer close(d)
	go consumer(c, d)
	producer(c, d)
}

// c chan int也能正常执行
// 但是 <- 提供了更加严格的语义上的约束：表示该通道作为方法参数，只能生产，不能用作接收
// 这种用法也可以用在方法 / 函数的返回值上
func producer(c chan<- int, d chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			time.Sleep(time.Second * 10)
		}
		c <- i
	}
	d <- true
}

func consumer(c <-chan int, d <-chan bool) {
	timer := time.NewTicker(tickTime)
	defer timer.Stop()
	buf := make([]int, 0, 5)
	// 如果不把select套在for里面，容易出现select执行完了，channel才接收到值，导致死锁
	for {
		select {
		case data := <-c:
			timer.Reset(tickTime)
			buf = append(buf, data)
			// 缓冲区满了
			if len(buf) >= cap(buf) {
				buf = sendAndResetBuf(buf)
			}
		case <-timer.C:
			if len(buf) > 0 {
				fmt.Println("reached time, prepare to send data ")
				buf = sendAndResetBuf(buf)
			}
		case <-d:
			fmt.Println("job done...")
			return
			//default: 用于处理阻塞时的情况
		}
	}
}

func sendAndResetBuf(buf []int) []int {
	// 切片是引用类型。goroutine接收的是拷贝吧？否则会有安全问题的
	go send(buf)
	// buf是地址的拷贝。这意味着对buf里的元素进行修改，会直接反映在外面的buf上
	// 但如果不将新的切片return出去，在外面接收，而是赋值给这里面的buf的话，外面的buf不会改变
	return make([]int, 0, 5)
}

func send(buf []int) {
	fmt.Println("send data...")
	// 如果只写 for i := range buf 的话，实际上获取的是下标值
	for _, v := range buf {
		fmt.Print(v, " ")
	}
	fmt.Println("send completed!")
}
