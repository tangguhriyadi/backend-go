/* package main

import "fmt"

func main() {
	var months = [...]string{
		"jan",
		"feb",
		"mar",
		"apr",
		"mei",
		"jun",
		"jul",
		"aug",
		"sep",
		"okt",
		"nov",
		"des",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "diubah"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "bulan 13")
	fmt.Println(slice3)

	slice3[1] = "bukan des"
	fmt.Println(slice3)
	fmt.Println(slice2)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "tangguh"
	newSlice[1] = "riyadi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)

	fmt.Println(copySlice)

}
 */