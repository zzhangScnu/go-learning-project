package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	go doSomething(ctx)
	go func() {
		time.Sleep(time.Second * 3)
		cancelFunc()
	}()
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("main function finished job: ", ctx.Err())
	//		return
	//	}
	//}
	time.Sleep(time.Second * 5)
}

func doSomething(ctx context.Context) {
	go innerDoSomething(ctx)
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("outer doing something...")
		case <-ctx.Done():
			fmt.Println("outer finished work: ", ctx.Err())
			// 跳出很重要！
			return
		}
	}
}

func innerDoSomething(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("inner doing something...")
		case <-ctx.Done():
			fmt.Println("inner finished work: ", ctx.Err())
			return
		}
	}
}
