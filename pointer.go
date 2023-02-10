package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func main() {
	address1 := Address{"cimahi", "jawa barat", "indonesia"}
	address2 := &address1 //jadi pass by reference

	address2.City = "bandung"

	*address2 = Address{"malang", "jawa timur", "indonesia"} //address 1 ngikut address 2

	fmt.Println(address1)
	fmt.Println(address2)
}
