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
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("<h1>This is the about page</h1>"))
}

func Redirect(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, 301)
	}
}

func ExecuteServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", subtreepath)  //Subtree path
	mux.HandleFunc("/hello", hello)   //Fixed path
	mux.HandleFunc("/header", header) //Fixed path
	mux.HandleFunc("/example", Redirect("http://example.com"))
	http.ListenAndServe(":8090", mux)
}
