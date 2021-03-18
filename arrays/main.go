package main

import "fmt"

func main(){
	evenNumbers := [5]int{0, 2, 4, 6, 8}

	for index, value := range evenNumbers{
		// %d is used for integers
		fmt.Printf("Index : %d => %d\n", index, value)
	}

	// len = 5 and index = 4, so we use < rather then <=
	for i:=0; i<len(evenNumbers); i++{
		fmt.Printf("Index : %d => %d\n", i, evenNumbers[i])
	}
}
