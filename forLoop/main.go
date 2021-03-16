package main

import "fmt"

func main() {
	var i int

	// increment
	for i = 0; i <= 10; i++ {
		fmt.Print(i ," ")
	}

	fmt.Print("\n")
	// decrement
	for i = 10; i >= 1; i-- {
		fmt.Print(i, " ")
	}

	fmt.Print("\n")
}
