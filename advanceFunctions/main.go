package main

import "fmt"

func main(){
	fmt.Println("Sum of all is : ", sum(1, 2, 3, 4, 5))
}

func sum(args ...int) int{
	sum := 0

	for _, value := range args{
		sum += value
	}

	return sum
}
