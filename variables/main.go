package main

import "fmt"

func main() {
	// defining a string variable
	var name string = "kashif"
	fmt.Println("Name : ", name)

	// defining a variable with short method
	age := 22
	fmt.Println("Age : ", age)

	// defining a variable with another method
	var gpa float64
	gpa = 3.5
	fmt.Println("GPA : ", gpa)

	// defining a constant
	const totalNumber = 550
	fmt.Println("Total Number : ", totalNumber)

	// define multiple variable
	var(
		name2 = "ovais"
		age2 = 25
	)
	fmt.Println("Name 2 : ", name2, " Age 2 : ", age2)

}
