package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Host struct {
	Title string
	Date  time.Time
	Text  string
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/temp", handlerTemp)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Hello GET")
	case "POST":
		fmt.Fprint(w, "Hello POST")
	default:
		fmt.Fprint(w, "Hello")
	}
}

func handlerTemp(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatalf("template ERROR")
	}
	host := Host{
		Title: "テストタイトル",
		Date:  time.Now(),
		Text:  "これは、netパッケージの練習です。",
	}

	t.Execute(w, host)
}
