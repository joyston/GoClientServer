package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func header(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v", name, h)
		}
	}
}

func ExecuteServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", header)
	http.ListenAndServe(":8090", nil)
}
