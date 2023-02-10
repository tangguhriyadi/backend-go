package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("tangguh")

	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(data)
	}
}
