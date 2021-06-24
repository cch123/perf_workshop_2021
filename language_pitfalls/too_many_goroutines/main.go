package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func sayhello(wr http.ResponseWriter, r *http.Request) {}

func main() {
	for i := 0; i < 1000000; i++ {
		go func() {
			time.Sleep(time.Second * 10)
		}()
	}
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":10003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
