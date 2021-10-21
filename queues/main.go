package main

import "fmt"

type Queue struct {
	list []int
}

func (q *Queue) Enqueue(item int) {
	q.list = append(q.list, item)
}

func (q *Queue) EnqueueMultiple(items ...int) {
	for _, item := range items {
		q.list = append(q.list, item)
	}
}

func (q *Queue) Dequeue() int {
	itemToRemove := q.list[0]
	q.list = q.list[1:len(q.list)]

	return itemToRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.EnqueueMultiple(20, 30, 40, 50, 60, 70, 80)
	fmt.Println("Queue: ", myQueue)
	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println("Queue After Two Dequeue: ", myQueue)
}
