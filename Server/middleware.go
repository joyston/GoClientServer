package jServer

import (
	"context"
	"net/http"
	"time"
)

type hellohandler struct {
	name string
}

func (h hellohandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + h.name))
}

func requestTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "requestTIme", time.Now().Format(time.RFC3339))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func helloClosure(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseText := "<h1>Hello " + name + "</h1>"

		if requestTime := r.Context().Value("requestTime"); requestTime != nil {
			if str, ok := requestTime.(string); ok {
				responseText = responseText + "<small>Generated at:" + str + "</small>"
			}
		}
		w.Write([]byte(responseText))
	})
}

func ExecuteMiddleware() {
	mux := http.NewServeMux()
	helloJohn := hellohandler{name: "John"}
	mux.Handle("/john", helloJohn)
	mux.Handle("/doe", requestTime(helloClosure("Doe")))
	http.ListenAndServe(":8080", mux)
}
