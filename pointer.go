package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIna(address Address) {
	address.Country = "Indonesia"
}
func changeCountryToInaPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"cimahi", "jawa barat", "indonesia"}
	address2 := &address1 //jadi pass by reference

	address2.City = "bandung"

	*address2 = Address{"malang", "jawa timur", "indonesia"} //address 1 ngikut address 2

	/* fmt.Println(address1)
	fmt.Println(address2) */

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	changeCountryToIna(alamat)
	fmt.Println(alamat) // tidak berubah

	changeCountryToInaPointer(&alamat)
	fmt.Println(alamat) // berubah

	var alamatPointer *Address = &alamat
	changeCountryToInaPointer(alamatPointer)
	fmt.Println(alamat) // berubah

}
