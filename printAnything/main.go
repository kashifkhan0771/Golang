package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please Enter Anything : ")
	scanner.Scan()
	userInput := scanner.Text()
	fmt.Println("You Typed : ", userInput)
}
