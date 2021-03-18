package main

import "fmt"

func main(){
	fmt.Println("Divide : ", div(4, 1))
	fmt.Println("Divide : ", div(13, 0))
}

func div(num1, num2 int) int{
	// recover will be done if panic occurs
	defer func() {
		if r:= recover(); r!= nil{
			fmt.Println(r)
		}
	}()

	solution := num1 / num2

	return solution
}
