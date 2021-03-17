package main

import "fmt"

// function with parameters
func add(x, y int) {
	result := x + y
	fmt.Println("RESULT OF SUM : ", result)
}

// function with parameters and return
func sub(x, y int) int {
	return x - y
}

// function with parameters and multiple return
func divMul(x, y int) (int, int) {
	multiplication := x * y
	division := x / y

	return multiplication, division
}

// function without parameter
func message() {
	fmt.Println("Welcome To Functions Example Program")
}

func main() {
	// calling a function
	message()
	// calling function with parameters
	add(12, 33)
	fmt.Println("RESULT OF SUB : ", sub(30, 11))
	// saving return from a function in variables
	r1, r2 := divMul(13, 12)
	fmt.Printf("RESULT OF MULTIPLICATION : %d \nRESULT OF DIVISION : %d\n", r1, r2)
}
