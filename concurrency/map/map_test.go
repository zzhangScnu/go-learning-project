package _map

import "testing"

func TestFetchData(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fetchDataConcurrently()
	}
}
