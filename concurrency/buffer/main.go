package main

import "fmt"

func main() {
	receivedByFor()
}

func receivedByRange() {
	// 带缓冲区的channel
	c := make(chan int, 2)
	go func() {
		// 发送方关闭channel。如果不关闭的话，发送方不生产数据了，接收方仍一直阻塞在获取数据上，会导致死锁
		// 如果channel被生产者关闭了，接收方消费数据完毕后会消费到零值，也可以获取channel是否已关闭
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
	// 表示有2个元素的缓冲区，加上原本的一个，共三个
	c := make(chan int, 2)
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
		fmt.Println("consumed ", <-c) // 0， 2
		// 如果通道被关闭，ok为false。这样就不会消费到零值
		v, ok := <-c // 1
		if !ok {
			break
		}
		fmt.Println("consumed ", v)
		//
	}
}
