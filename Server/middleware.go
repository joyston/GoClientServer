package jServer

import "net/http"

type hellohandler struct {
	name string
}

func (h hellohandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + h.name))
}

func helloClosure(h string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + h))
	})
}

func ExecuteMiddleware() {
	mux := http.NewServeMux()
	helloJohn := hellohandler{name: "John"}
	mux.Handle("/john", helloJohn)
	mux.Handle("/doe", helloClosure("Doe"))
	http.ListenAndServe(":8080", mux)
}
