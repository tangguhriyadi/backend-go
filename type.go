package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case bool:
		fmt.Println("Boolean", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown")
	}

}
