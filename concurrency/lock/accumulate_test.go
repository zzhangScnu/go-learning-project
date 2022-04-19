package main

import (
	"fmt"
	"testing"
)

func TestLock(t *testing.T) {
	for i := 0; i < 100; i++ {
		runAccumulator(AccumulateByLock)
	}
}

func TestChannel(t *testing.T) {
	for i := 0; i < 100; i++ {
		runAccumulator(AccumulateByChannel)
	}
}

func runAccumulator(accumulator func(int) []int) {
	concurrency := 100
	result := accumulator(concurrency)
	if len(result) != concurrency {
		fmt.Printf("test failed! len of result is: %d\n", len(result))
	}
	fmt.Println(result)
}
