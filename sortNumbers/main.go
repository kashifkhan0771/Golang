package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Numbers Sorting Function")
	fmt.Print("How Many Numbers You Wanna Add Into List : ")
	scanner.Scan()
	totalNumbers, err := strconv.Atoi(scanner.Text())
	if err != nil {
		_ = fmt.Errorf("please enter a integer value")
	}
	allNumbers := getAllInputs(totalNumbers, *scanner)
	sortedNumbers := sortOut(allNumbers, *scanner)
	if sortedNumbers == nil {
		_ = fmt.Errorf("some error occur please try again")
	} else {
		fmt.Println("Sorted List : ", sortedNumbers)
	}
}

func getAllInputs(inputsToGet int, scanner bufio.Scanner) []int {
	var allNumbers []int
	for i := 1; i <= inputsToGet; i++ {
		fmt.Printf("Please Enter %s Number : ", strconv.Itoa(i))
		scanner.Scan()
		noToAdd, err := strconv.Atoi(scanner.Text())
		if err != nil {
			_ = fmt.Errorf("please enter a integer value")
		} else {
			allNumbers = append(allNumbers, noToAdd)
		}
	}

	return allNumbers
}

func sortOut(allNumbers []int, scanner bufio.Scanner) []int {
	fmt.Print("Please Enter 1 For Ascending and 2 For Descending : ")
	scanner.Scan()
	sortType, err := strconv.Atoi(scanner.Text())
	if err != nil {
		_ = fmt.Errorf("please enter correct number")
		return nil
	} else if sortType == 1 {
		sort.Ints(allNumbers)

		return allNumbers
	} else if sortType == 2{
		sort.Slice(allNumbers, func(i, j int) bool {
			return allNumbers[i] > allNumbers[j]
		})

		return allNumbers
	} else{
		fmt.Print("PLEASE CHOOSE 1 OR 2")
		return nil
	}
}
