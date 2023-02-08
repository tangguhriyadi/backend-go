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

func getCompleteName() (pertama, kedua, ketiga string) { // returning named value
	pertama = "muhammad"
	kedua = "tangguh"
	ketiga = "riyadi"

	return
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

	a, b, c := getCompleteName()

	fmt.Println(a, b, c)

}
