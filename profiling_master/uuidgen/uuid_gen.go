package main

import (
	"io"
	"log"
	_ "net/http/pprof"
	"os/exec"

	"net/http"
)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	uuid, _ := exec.Command("uuidgen").Output()

	wr.Header()["Content-Type"] = []string{"application/text"}
	io.WriteString(wr, string(uuid))
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":10003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
