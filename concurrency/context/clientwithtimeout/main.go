package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type ResponseData struct {
	Resp *Data
	Err  error
}

func call() <-chan *ResponseData {
	// 1是为了防止读取阻塞？
	resChan := make(chan *ResponseData, 1)
	go func() {
		response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
		if err != nil {
			resChan <- &ResponseData{nil, fmt.Errorf("get url failed: %s", err)}
			// 不return的话会继续往下走，引发panic
			return
		}
		defer func() {
			err := response.Body.Close()
			if err != nil {
				log.Fatal("close response body failed: ", err)
			}
		}()
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			resChan <- &ResponseData{nil, fmt.Errorf("read response failed: %s", err)}
			return
		}
		data := &Data{}
		err = json.Unmarshal(bytes, data)
		if err != nil {
			resChan <- &ResponseData{nil, fmt.Errorf("unmarshal failed: %s", err)}
		}
		resChan <- &ResponseData{data, nil}
	}()
	return resChan
}

func getResponseData(ctx context.Context) (*Data, error) {
	select {
	case res := <-call():
		return res.Resp, res.Err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	data, err := getResponseData(ctx)
	if err != nil {
		fmt.Println("failed! ", err)
	} else {
		fmt.Printf("query success! data is: %v ", data)
	}
}
