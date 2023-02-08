/* package main

import "fmt"

func sumAll(numbers ...int) int { // numbers dikonversi menjadi slice
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)

	numbers := []int{10, 20, 30, 40}

	totals := sumAll(numbers...)

	fmt.Println(total)
	fmt.Println(totals)
}
*/