package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := pembagian(10, 0)

	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}
}
