package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
void output(char *str) {
    usleep(1000000);
    printf("%s\n", str);
}
*/
import "C"

import (
	"net/http"
	"unsafe"

	"log"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

func init() {
	go http.ListenAndServe(":10003", nil)
}

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			str := "hello cgo"
			//change to char*
			cstr := C.CString(str)
			C.output(cstr)
			C.free(unsafe.Pointer(cstr))

		}()
	}
	killThreadService()
	select {}
}

func sayhello(wr http.ResponseWriter, r *http.Request) {
	KillOne()
}

func killThreadService() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":10003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// KillOne kills a thread
func KillOne() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		runtime.LockOSThread()
		return
	}()

	wg.Wait()
}


