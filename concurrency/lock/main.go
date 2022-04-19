package main

import (
	"sync"
)

func AccumulateByLock(concurrency int) []int {
	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	mux := sync.Mutex{}
	result := make([]int, 0, concurrency) // make([]int, concurrency) 会创建一个含有concurrency个int零值的切片
	//var result []int
	for i := 0; i < concurrency; i++ {
		// defer wg.Done()
		// 隔离的作为参数；共享的直接读取外部变量
		go func(i int) {
			defer wg.Done()
			mux.Lock()
			result = append(result, i)
			mux.Unlock()
		}(i)
	}
	wg.Wait()
	return result
}

func AccumulateByChannel(concurrency int) []int {
	numChan := make(chan int, concurrency)
	var result []int
	defer close(numChan)
	for i := 0; i < concurrency; i++ {
		go func(i int) {
			numChan <- i
		}(i)
	}
	// select + for + break loop
	for num := range numChan {
		result = append(result, num)
		if len(result) == concurrency {
			break
		}
	}
	return result
}
