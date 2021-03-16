package main

import "fmt"

func main(){
	names := []string{"kashif", "ovais", "shahzad", "moeed", "ahsan", "abiha", "ahmed"}

	for index, name := range names{
		fmt.Printf("%d : %s \n", index, name)
	}
}
