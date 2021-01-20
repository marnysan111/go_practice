package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Users []User

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/json", jsonCheck)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func jsonCheck(w http.ResponseWriter, r *http.Request) {

	var users []User
	/*
		user := Users{
			User{Name: "root", Age: 19},
			User{Name: "Robot", Age: 0},
		}
	*/
	bytes, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}
	//outputJson, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &users)
	test, err := json.Marshal(users)
	if err != nil {
		fmt.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	/*
		fmt.Println(string(bytes))
		fmt.Fprint(w, string(bytes))
		fmt.Println(string(outputJson))
		fmt.Fprint(w, string(outputJson))
		fmt.Println(outputJson)
		fmt.Println(bytes)
	*/
	fmt.Println(string(test))
	fmt.Fprint(w, string(test))
}
