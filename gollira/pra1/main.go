package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := mux.NewRouter()
	// "www.example.com"をドメイン名設定
	r.HandleFunc("/sample", handler)
	r.HandleFunc("/json", jsonCheck)
	//r.HandleFunc("/json2", jsonCheck2)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func jsonCheck(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("test.json")
	if err != nil {
		log.Fatalf("json error")
	}

	bytes, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}
	var user []User
	json.Unmarshal(bytes, &user)
	fmt.Println(user)
	for _, p := range user {
		fmt.Printf("%d : %s\n", p.ID, p.Name)
	}
	fmt.Fprint(w, t)
}

/*
func jsonCheck2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/
}
*/
