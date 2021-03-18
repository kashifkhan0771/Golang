package main

import "fmt"

func main(){
	var x int = 10
	// with pointer change
	fmt.Println("Before Change : ", x)
	changeValue(&x)
	fmt.Println("After Change : ", x)

	// without pointer change (value will not be changed)
	fmt.Println("Before Without Pointer Change : ", x)
	changeValueWithoutPointer(x)
	fmt.Println("After Without Pointer Change : ", x)

}

func changeValue(x *int){
	*x = 5
}

func changeValueWithoutPointer(x int){
	x = 2
}
