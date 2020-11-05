package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("---perTest---")
	perTest()
	fmt.Println("---arrTest---")
	arrTest()
	fmt.Println("---sliTest---")
	sliTest()

}

func perTest() {
	var Mani Person
	Mani.name = "manisan"
	Mani.age = 18
	fmt.Println("[name]:", Mani.name, "[age]:", Mani.age)

	abe := Person{"abe", 18}
	fmt.Println("[name]:", abe.name, "[age]:", abe.age)

	taku := Person{name: "taku", age: 18}
	fmt.Println("[name]:", taku.name, "[age]:", taku.age)

	fmt.Println(Mani)
	fmt.Println(abe)
	fmt.Println(taku)
}

func arrTest() {
	var arr [2]string
	arr[0] = "Golang"
	arr[1] = "Test"
	fmt.Println("[arr[0]]:", arr[0], "[arr[1]]:", arr[1])
	fmt.Println(arr)

	arr2 := [...]string{"Golang", "Test"}
	fmt.Println("[arr2[0]]:", arr2[0], "[arr2[1]]:", arr2[1])
	fmt.Println(arr2)
}

func sliTest() {
	slice := []string{"Golang", "Test"}
	fmt.Println(slice[0:2])
	fmt.Println(len(slice), cap(slice))
}
