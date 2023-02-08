package main

import "fmt"

func sayHello(count int) {
	fmt.Println(count)
}

func hasReturn(name string) string {
	return "my name is " + name
}

func getFullname() (string, string) { //return multiple value
	return "tangguh", "riyadi"
}

func main() {
	for i := 0; i <= 10; i++ {
		sayHello(i)
	}

	namaSaya := hasReturn("tangguh")

	fmt.Println(namaSaya)

	firstName, lastName := getFullname()
	fmt.Println(firstName, lastName)

	first, _ := getFullname()
	fmt.Println(first)
}
