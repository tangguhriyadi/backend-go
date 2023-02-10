package main

import "fmt"

type HasName interface {
	getName() string
}

func hello(hasName HasName) {
	fmt.Println("hello", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func Ups() interface{} { // empty interface return
	return "Ups"
}

func main() {
	var tangguh Person
	tangguh.Name = "tangguh"
	hello(tangguh)

	cat := Animal{Name: "KOCHENG"}
	hello(cat)

	kosong := Ups()
	fmt.Println(kosong)
}
