package main

import (
	"log"
	"net/http"
)

type myHandler struct {
	message string
}

func (m myHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//panic("implement me")
	log.Print("get request: ", request)
	_, err := writer.Write([]byte(m.message))
	if err != nil {
		log.Print(err)
	}
}

func startupServerMux() {
	// ServeMux实际实现了Handler接口，是一个多路复用的HTTP请求路由器
	mux := http.NewServeMux()
	mux.Handle("/en", myHandler{"hello"})
	mux.Handle("/cn", myHandler{"你好"})
	mux.HandleFunc("/heartbeat", func(writer http.ResponseWriter, request *http.Request) {
		log.Print("i'm still alive!")
	})
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
