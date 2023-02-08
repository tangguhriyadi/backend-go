package main

import "fmt"

func getGoodbye(name string) string {
	return "goodbye " + name
}

func main() {
	goodbye := getGoodbye

	fmt.Println(goodbye("tangguh"))
}
