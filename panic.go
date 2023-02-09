/* package main

import "fmt"

func endApp() {

	message := recover()

	if message != nil {
		fmt.Println("error occured", message)
	}
	fmt.Println("ended")

}

func runApp(err bool) {
	defer endApp()

	if err {
		panic("ERROR")
	}

	fmt.Println("app running")

}

func main() {
	runApp(false)
}
*/