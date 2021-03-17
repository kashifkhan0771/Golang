package main

import "fmt"

func main(){
	// defer execute line after surrounding function ends
	defer fmt.Print("World\n")
	fmt.Print("Hello ")
}
