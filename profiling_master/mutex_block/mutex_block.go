package main

import (
	"io"
	"log"
	"time"
	"sync"
	_ "net/http/pprof"

	"net/http"
)

var mu sync.Mutex
var data = map[string]string{
	"hint" : "hello world",
}

func sayhello(wr http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	buf := data["hint"]
	
	// 假设这里是一些非常慢的 io 操作
	time.Sleep(time.Millisecond * 10)

	wr.Header()["Content-Type"] = []string{"application/json"}
	io.WriteString(wr, buf)
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":10003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

