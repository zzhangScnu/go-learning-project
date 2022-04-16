package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var intChan = make(chan int)

var errChan = make(chan error)

var finChan = make(chan struct{})

var wg = sync.WaitGroup{}

func main() {
	run()
Loop:
	for {
		select {
		case i := <-intChan:
			fmt.Println("received num: ", i)
		case err := <-errChan:
			fmt.Println(err.Error())
			break Loop
		case <-finChan:
			fmt.Println("task end successfully!")
			// 如果不加这句break Loop，循环会继续下去，由于已经执行了finChan已经被执行关闭，会一直走到这条分支
			break Loop
		case <-time.After(time.Millisecond * 1_00_00):
			fmt.Println("times up")
			break Loop
		}
	}
}

func run() {
	wg.Add(100)
	go finalize()
	for i := 0; i < 100; i++ {
		go doRun(i, &wg)
	}
}

// 这里要用指针传参。“A WaitGroup must not be copied after first use.”
func doRun(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))
	intChan <- i
	// try to change this to -1 to test the condition of normal completion
	if i == 66 {
		errChan <- errors.New("reached error num")
	}
}

func finalize() {
	wg.Wait()
	close(finChan)
}
