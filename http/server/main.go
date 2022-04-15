package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	addr string
}

func assignHandler() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "health check")
		if err != nil {
			log.Fatal(err)
		}
	})
}

func startupServer() {
	server := Server{":8080"}
	err := http.ListenAndServe(server.addr, nil)
	if err != nil {
		log.Fatal("server startup failed")
	}
}

func init() {
	assignHandler()
	fmt.Println("server startup...")
	startupServer()
}

func main() {
}
