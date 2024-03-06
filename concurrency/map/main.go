package _map

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

/**
这里的 wg 使用不恰当
若下游系统发生超时时，该 wg 其实并没有完成
这也就意味着，其子任务也并没有全部完成。
虽然在 fetchData 内部对 map 的修改加了写锁，但若下游超时，在 fetchData 返回后，fetchData 内部启动的 goroutine 仍然可能对返回的 map 进行修改。
当 map 对象同时进行加锁的 write 和不加锁的读取时，也会发生崩溃。
不加锁的读取发生在什么地方呢？其实就是这里例子的 log.Printf。如果你做个 json.Marshal 之类的，效果也差不多。
至于为什么是偶发，超时本来也不是经常发生的，看起来这个 bug 就变成了一个偶现 bug。
*/
type resp struct {
	k string
	v string
}

func fetchDataConcurrently() {
	res, err := fetchData()
	if err != nil {
		log.Println(err)
	}
	str, _ := json.Marshal(res)
	log.Printf("%s", str)
}

func rpcWork() resp {
	// do some rpc work
	time.Sleep(time.Millisecond * 5) // 下游接口超时时，有概率会触发map的panic
	return resp{}
}

func fetchData() (map[string]string, error) {
	var result = map[string]string{} // result is k -> v
	var keys = []string{"a", "b", "c"}
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < len(keys); i++ {
		wg.Add(1)

		go func(i int) {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()

			// do some rpc
			resp := rpcWork()

			result[keys[i]] = resp.v
		}(i)
	}

	waitTimeout(&wg, time.Millisecond)
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
