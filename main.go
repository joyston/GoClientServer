package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
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
