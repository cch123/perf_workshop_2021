package main

import (
	"io"
	"log"
	_ "net/http/pprof" // 一般开 pprof，在你的 main.go 里加上这一行就行了

	"net/http"
)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	wr.Header()["Content-Type"] = []string{"application/json"}
	io.WriteString(wr, "hello")

}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":10003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// 也可以像 etcd 那样做一些定制
// https://github.com/etcd-io/etcd/blob/e2d67f2e3bfa6f72178e26557bb22cc1482c418c/pkg/debugutil/pprof.go#L26
// 有些公司内有热开启/关闭 pprof 的需求
