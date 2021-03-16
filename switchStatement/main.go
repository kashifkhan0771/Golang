package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var day string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please Enter Day Name : ")
	scanner.Scan()
	day = scanner.Text()

	switch day {
	case "Monday":
		fmt.Print("Today Is ", day)
	case "Tuesday":
		fmt.Print("Today Is ", day)
	case "Wednesday":
		fmt.Print("Today Is ", day)
	case "Thursday":
		fmt.Print("Today Is ", day)
	case "Friday":
		fmt.Print("Today Is ", day)
	case "Saturday":
		fmt.Print("Today Is ", day)
	case "Sunday":
		fmt.Print("Today Is ", day)
	default:
		fmt.Print("Please Enter a valid day")
	}
	fmt.Print("\n")
}
