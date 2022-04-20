package queue

import (
	"fmt"
	"testing"
)

var jobCount = 100

var concurrency = 10

func TestMethod1(t *testing.T) {
	runTest(t, ProcessMethod1)
}

func runTest(t *testing.T, f func(int, int) []int) {
	results := f(jobCount, concurrency)
	fmt.Println(results)
	if len(results) != jobCount {
		t.Errorf("test failed!")
	}
}

func TestProcessMethod1ByWaitGroup(t *testing.T) {
	runTest(t, ProcessMethod1ByWaitGroup)
}

func TestMethod2(t *testing.T) {
	runTest(t, ProcessMethod2)
}
