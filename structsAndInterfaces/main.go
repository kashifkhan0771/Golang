package main

import "fmt"

type Details interface {
	details()
}

type Book struct {
	id          int
	name        string
	price       float64
	author      Author
	totalCopies int
	buyers      []Buyer
}

type Author struct {
	id         int
	name       string
	rating     float64
	totalBooks int
}

type Buyer struct {
	id   int
	name string
}

func (b Book) details() {
	fmt.Println("Id           : ", b.id)
	fmt.Println("Name         : ", b.name)
	fmt.Println("Price        : ", b.price)
	fmt.Println("Author       : ", b.author.name)
	fmt.Println("Total Copies : ", b.totalCopies)
	fmt.Println("Buyers       : ", b.buyers)
}

func (auth Author) details() {
	fmt.Println("Id          : ", auth.id)
	fmt.Println("Name        : ", auth.name)
	fmt.Println("Rating      : ", auth.rating)
	fmt.Println("Total Books : ", auth.totalBooks)
}

func (by Buyer) details() {
	fmt.Println("Id   : ", by.id)
	fmt.Println("Name : ", by.name)
}

func getDetails(d Details) {
	d.details()
}

func main() {
	author1 := Author{
		id:         1,
		name:       "Kashif Khan",
		rating:     8.3,
		totalBooks: 3,
	}

	buyer1 := Buyer{
		id:   1,
		name: "Shahzad Haider",
	}

	buyer2 := Buyer{
		id:   2,
		name: "Ovais Tariq",
	}

	b1 := Book{
		id:          1,
		name:        "Golang",
		author:      author1,
		price:       8000,
		totalCopies: 50,
		buyers:      []Buyer{buyer1, buyer2},
	}

	getDetails(author1)
	getDetails(b1)
	getDetails(buyer1)
}
