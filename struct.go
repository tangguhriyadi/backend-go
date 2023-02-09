package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Is_special    bool
}

func main() {
	// var tangguh Customer
	// tangguh.Name = "tangguh"
	// tangguh.Address = "cimahi"
	// tangguh.Age = 26
	// tangguh.Is_special = true

	// fmt.Println(tangguh)

	tangguh := Customer{
		Name:       "tangguh",
		Address:    "cimahi",
		Age:        26,
		Is_special: true,
	}

	fmt.Println(tangguh)
}
