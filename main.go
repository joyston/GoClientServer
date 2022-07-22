package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func ExecuteClient() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	Scanner := bufio.NewScanner(resp.Body)

	for i := 0; Scanner.Scan() && i < 5; i++ {
		fmt.Println(Scanner.Text())
	}

	if err := Scanner.Err(); err != nil {
		panic(err)
	}
}

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

func main() {
	// ExecuteClient()
	ExecuteServer()
}
