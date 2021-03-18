package main

import "fmt"

func main() {
	// maps are just key value pairs
	internNames := make(map[int]string)
	internAges := make(map[string]int)

	internNames[1] = "Kashif Khan"
	internNames[2] = "Ovais Tariq"
	internNames[3] = "Shahzad Haider"

	internAges["Kashif Khan"] = 22
	internAges["Ovais Tariq"] = 25
	internAges["Shahzad Haider"] = 23

	for index, name := range internNames {
		fmt.Printf("%d = %s\n", index, name)
	}

	fmt.Println(internAges)
}
