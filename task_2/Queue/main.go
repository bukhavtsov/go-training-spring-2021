package main

import "fmt"

type Queue struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

func (q *Queue) Enqueue(value interface{}) {
	node := &Node{
		prev:  q.last,
		value: value,
	}
	if q.last != nil {
		q.last.next = node
	}
	q.last = node
	for node.prev != nil {
		node = node.prev
	}
	q.first = node
	q.size++
}

func (q *Queue) Dequeue() {
	if q.size > 0 {
		node := q.first.next
		q.first = node
		if q.size != 1 {
			q.first.prev = nil
		}
		q.size--
	} else {
		fmt.Println("queue has not been initialized")
	}
}

func (q *Queue) Display() {
	firstNode := q.first
	for firstNode != nil {
		fmt.Println(firstNode.value)
		firstNode = firstNode.next
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size != 0
}

func (q *Queue) Peek() interface{} {
	return q.first.value
}

func main() {
	q := Queue{}
	fmt.Println(q.IsEmpty())
	fmt.Println(q.IsFull())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Display()
	q.Dequeue()
	q.Display()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.IsFull())
	fmt.Println(q.Peek())
}
