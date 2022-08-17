package _map

import (
	"log"
	"sync"
	"time"
)

type resp struct {
	k string
	v string
}

func fetchDataConcurrently() {
	res, err := fetchData()
	if err != nil {
		log.Println(err)
	}
	log.Print(res)
}

func rpcWork() resp {
	// do some rpc work
	time.Sleep(time.Second * 2) // 下游接口超时时，有概率会触发map的panic
	return resp{}
}

func fetchData() (map[string]string, error) {
	var result = map[string]string{} // result is k -> v
	var keys = []string{"a", "b", "c"}
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < len(keys); i++ {
		wg.Add(1)

		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()

			// do some rpc
			resp := rpcWork()

			result[resp.k] = resp.v
		}()
	}

	waitTimeout(&wg, time.Second)
	return result, nil
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
