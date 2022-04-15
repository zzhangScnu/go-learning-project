package main

import (
	"log"
	"net/http"
)

func main() {
	getByHttp()
	getByClient()
}

func getByHttp() {
	r, err := http.Get("https://baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r)
}

func getByClient() {
	req, err := http.NewRequest("GET", "https://baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
