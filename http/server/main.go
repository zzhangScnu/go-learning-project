package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print("get request: ", r)
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
}

// 自定义服务器
var server = &http.Server{
	Addr:         ":8080",
	ReadTimeout:  time.Second * 10,
	WriteTimeout: time.Second * 10,
}

func assignHandler() {
	// http.Handle("...", myHandler{...})
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("get request: ", r)
		_, err := fmt.Fprintf(w, "health check")
		if err != nil {
			log.Fatal(err)
		}
	})
}

func startupServer() {
	// err := http.ListenAndServe(server.Addr, nil)
	// server.Handler = mux / myHandler
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("server startup failed")
	}
}

func init() {
	assignHandler()
	fmt.Println("server startup...")
	// 启动两个服务器，监听不同的端口
	go startupServer()
	startupServerMux()
}

func main() {
}
