package queue

import (
	"math/rand"
	"sync"
	"time"
)

func ProcessMethod1(jobCount int, concurrency int) []int {
	valChan := make(chan int)
	queueChan := make(chan struct{}, concurrency)
	send(jobCount, queueChan, valChan)
	results := receive(jobCount, valChan)
	close(valChan)
	return results
}

func send(jobCount int, queueChan chan struct{}, valChan chan int) {
	for i := 0; i < jobCount; i++ {
		go func() {
			queueChan <- struct{}{}
		}()
		go func(i int) {
			time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))
			valChan <- i
			<-queueChan
		}(i)
	}
}

func receive(jobCount int, valChan chan int) []int {
	var results []int
	for val := range valChan {
		results = append(results, val)
		if len(results) == jobCount {
			break
		}
	}
	return results
}

func ProcessMethod1ByWaitGroup(jobCount int, concurrency int) []int {
	valChan := make(chan int)
	queueChan := make(chan struct{}, concurrency)
	wg := &sync.WaitGroup{}
	wg.Add(jobCount)
	sendByWg(jobCount, queueChan, valChan, wg)
	stopWhenDone(valChan, wg)
	return receiveByWg(valChan)
}

func sendByWg(jobCount int, queueChan chan struct{}, valChan chan int, wg *sync.WaitGroup) {
	for i := 0; i < jobCount; i++ {
		go func() {
			queueChan <- struct{}{}
		}()
		go func(i int) {
			// 在valChan发送后并被接收后，用defer才会被调用
			defer func() {
				wg.Done()
				<-queueChan
			}()
			time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))
			valChan <- i
		}(i)
	}
}

func stopWhenDone(valChan chan int, wg *sync.WaitGroup) {
	// 代替上面方法的跳出for循环的判断逻辑
	go func() {
		wg.Wait()
		close(valChan)
	}()
}

func receiveByWg(valChan chan int) []int {
	var results []int
	for val := range valChan {
		results = append(results, val)
	}
	return results
}
