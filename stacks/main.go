package main

import "fmt"

type Stack struct {
	list []int
}

func (s *Stack) Push(value int) {
	s.list = append(s.list, value)
}

func (s *Stack) PushMultiple(values ...int) {
	for _, value := range values {
		s.list = append(s.list, value)
	}
}

func (s *Stack) Pop() int {
	removedItem := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return removedItem
}

func main() {
	myStack := Stack{}
	myStack.Push(10)
	myStack.PushMultiple(20, 30, 40, 50, 60, 70)
	fmt.Println("My Stack: ", myStack)
	fmt.Println("Last Item Pop: ", myStack.Pop())
	fmt.Println("Stack After Pop: ", myStack)
}
