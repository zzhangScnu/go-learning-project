package queue

import (
	"math/rand"
	"sync"
	"time"
)

func ProcessMethod2(jobCount int, concurrency int) []int {
	wg := &sync.WaitGroup{}
	wg.Add(concurrency)
	valChan := make(chan int)
	queueChan := make(chan int)
	startQueue(jobCount, queueChan)
	transport(concurrency, valChan, queueChan, wg)
	stopWhenFinish(valChan, wg)
	return consume(valChan)
}

func startQueue(jobCount int, queueChan chan int) {
	go func() {
		defer close(queueChan)
		for i := 0; i < jobCount; i++ {
			queueChan <- i
		}
	}()
}

func transport(concurrency int, valChan chan int, queueChan chan int, wg *sync.WaitGroup) {
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for val := range queueChan {
				time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))
				valChan <- val
			}
		}()
	}
}

func stopWhenFinish(valChan chan int, wg *sync.WaitGroup) {
	go func() {
		wg.Wait()
		close(valChan)
	}()
}

func consume(valChan chan int) []int {
	var results []int
	for val := range valChan {
		results = append(results, val)
	}
	return results
}
