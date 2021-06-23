package main

import (
	"io"
	"log"
	_ "net/http/pprof"

	//"os/exec"

	"net/http"

	uuid "github.com/satori/go.uuid"
)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	uu, _ := uuid.NewV4()

	wr.Header()["Content-Type"] = []string{"application/text"}
	io.WriteString(wr, uu.String())
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":10003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
