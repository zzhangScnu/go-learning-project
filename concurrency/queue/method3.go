package queue

import (
	"math/rand"
	"sync"
	"time"
)

func ProcessMethod3(jobCount int, concurrency int) []int {
	queueChan := make(chan int, jobCount)
	valChan := make(chan int, jobCount)
	wg := &sync.WaitGroup{}
	wg.Add(jobCount)
	produce(jobCount, queueChan)
	transfer(concurrency, valChan, queueChan, wg)
	stop(valChan, wg)
	return collectResult(valChan)
}

func produce(jobCount int, queueChan chan int) {
	defer close(queueChan)
	for i := 0; i < jobCount; i++ {
		queueChan <- i
	}
}

func transfer(concurrency int, valChan chan<- int, queueChan <-chan int, wg *sync.WaitGroup) {
	for i := 0; i < concurrency; i++ {
		go func() {
			// defer wg.Done()
			for val := range queueChan {
				defer wg.Done()
				time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))
				valChan <- val
			}
		}()
	}
}

func stop(valChan chan<- int, wg *sync.WaitGroup) {
	wg.Wait()
	close(valChan)
}

func collectResult(valChan <-chan int) []int {
	var results []int
	for val := range valChan {
		results = append(results, val)
	}
	return results
}
