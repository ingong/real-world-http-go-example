package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888?")
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
	log.Println("Content Length:", resp.ContentLength)
	log.Println("Body:", resp.Body)
}
