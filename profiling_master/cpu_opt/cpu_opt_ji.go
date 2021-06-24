package main

import (
	"io"
	"log"
	_ "net/http/pprof"

	"net/http"
	json "github.com/json-iterator/go"

)

var bigMap = make(map[int]int)
func init() {
	for i:=0;i<10000;i++ {
		bigMap[i] = i + 1
	}
}


func sayhello(wr http.ResponseWriter, r *http.Request) {
	d, _ := json.Marshal(bigMap)
	wr.Header()["Content-Type"] = []string{"application/json"}
	io.WriteString(wr, string(d))

}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":10003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

