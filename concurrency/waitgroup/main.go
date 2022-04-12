package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		fmt.Println("loop i ", i)
		go func(i int) {
			fmt.Println("run i ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
