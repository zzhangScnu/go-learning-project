package main

import "fmt"

func main() {
	receivedByFor()
}

func receivedByRange() {
	// 带缓冲区的channel
	c := make(chan int, 3)
	go func() {
		// 发送方关闭channel。如果不关闭的话，发送方不生产数据了，接收方仍一直阻塞在获取数据上，会导致死锁
		// 关闭必须由生产者来做。如果放在外面，没有新数据产生的时候，消费者已经阻塞了，不会调到defer函数
		defer close(c)
		for i := 0; i < 3; i++ {
			fmt.Println("produced ", i)
			c <- i
		}
	}()
	for v := range c {
		fmt.Println("consumed ", v)
	}
}

func receivedByFor() {
	c := make(chan int, 3)
	go func() {
		// 如果没有关闭语句，发送者和生产者都会阻塞，产生死锁和goroutine泄漏
		defer close(c)
		for i := 0; i < 3; i++ {
			fmt.Println("produced ", i)
			c <- i
		}
	}()
	for {
		// 到最后一直消费到零值，因为通道其实已经被发送者关闭了，这就是goroutine泄漏
		// fmt.Println("consumed ", <-c)
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Println("consumed ", v)
	}
}
