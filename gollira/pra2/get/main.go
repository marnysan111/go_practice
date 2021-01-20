package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	r, err := http.Get("http://192.168.1.10:8080/json")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(r.Body)
	var user []User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}
	fmt.Println(user[0].Name)
	data, _ := json.Marshal(user[1])
	fmt.Println(string(data))

	/*
		a := "test" + "test"
		fmt.Println(a)
		f := 8080
		q := strconv.Itoa(f)
		b := "aaa" + q + "aaa"
		fmt.Println(b)
	*/
}
