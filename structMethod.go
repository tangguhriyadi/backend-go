package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Is_Married    bool
}

func (customer Customer) hello(admin string) {
	fmt.Println("hello", customer.Name, "my name is", admin)
}

func main() {

	rully := Customer{Name: "Rully"}
	rully.hello("tangguh")

}
