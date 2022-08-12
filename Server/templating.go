package jServer

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func ExecuteTemplating() {
	file, err := os.OpenFile("teplatelogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	todos := []string{"Eat", "Sleep", "Code", "Repeat"}
	/*err = t.Execute(os.Stdout, todos)
	if err != nil {
		log.Fatal(err)
	}*/

	mux := http.NewServeMux()
	mux.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, todos)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.ListenAndServe(":8080", mux)
	log.SetOutput(file)
}
