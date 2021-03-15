package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name : ")
	scanner.Scan()
	userName := scanner.Text()
	fmt.Printf("Hey %s, Please enter your age : \n", userName)
	scanner.Scan()
	userAge, err := strconv.Atoi(scanner.Text())
	if err != nil{
		_ = fmt.Errorf("please enter correct number %s", err)
	}

	if userAge < 18{
		fmt.Print("You are less then 18 \n")
	}else if userAge >= 18 && userAge <= 40{
		fmt.Print("You are young, aged between 18 and 40 \n")
	}else if userAge > 40{
		fmt.Print("Your age is greater then 40 \n")
	}
}
