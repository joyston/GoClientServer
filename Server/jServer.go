package jServer

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	log.Println("Hello world")
}

func header(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v", name, h)
		}
	}
}

func subtreepath(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is the about page</h1>"))
}

func ExecuteServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", subtreepath)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/header", header)
	http.ListenAndServe(":8090", mux)
}
