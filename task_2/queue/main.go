package main

import "fmt"

//Node this is a Queue element.
type Node struct {
	value interface{}
}

//Queue is a data structure that, like a stack, has restrictions on adding and removing elements.
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

//NewQueue - Creates a new queue
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

//Enqueue - Add an element to the end of the queue.
func (Q *Queue) Enqueue(n interface{}) {
	if Q.head == Q.tail && Q.count > 0 {
		nodes := make([]*Node, len(Q.nodes)+Q.size)
		copy(nodes, Q.nodes[Q.head:])
		copy(nodes[len(Q.nodes)-Q.head:], Q.nodes[:Q.head])
		Q.head = 0
		Q.tail = len(Q.nodes)
		Q.nodes = nodes
	}
	Q.nodes[Q.tail] = &Node{value: n}
	Q.tail = (Q.tail + 1) % len(Q.nodes)
	Q.count++
}

//Dequeue - Remove an element from the front of the queue.
func (Q *Queue) Dequeue() interface{} {
	if Q.IsEmpty() {
		fmt.Println("The queue is empty!")
		return nil
	}
	node := Q.nodes[Q.head]
	Q.head = (Q.head + 1) % len(Q.nodes)
	Q.count--
	return node.value
}

//IsEmpty - Check if the queue is empty.
func (Q *Queue) IsEmpty() bool {
	return Q.count == 0
}

//IsFull - Check if the queue is full.
func (Q *Queue) IsFull() bool {
	return Q.size <= Q.count
}

//Peek - Get the value of the front of the queue without removing it.
func (Q *Queue) Peek() interface{} {
	if Q.IsEmpty() {
		fmt.Println("The queue is empty!")
		return nil
	}
	return Q.nodes[Q.head].value
}

func main() {

	q := NewQueue(10)

	fmt.Println(q.count)
	fmt.Println(q.nodes)
	fmt.Println(q.head)
	fmt.Println(q.tail)
	fmt.Println(q.size)
	fmt.Println()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue("fdhjgf")
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(8)
	q.Enqueue(8)
	q.Enqueue(8)
	q.Enqueue(8)

	fmt.Println(q.count)
	fmt.Println(q.nodes)
	fmt.Println(q.head)
	fmt.Println(q.tail)
	fmt.Println(q.size)
	fmt.Println()

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println()

	fmt.Println(q.count)
	fmt.Println(q.nodes)
	fmt.Println(q.head)
	fmt.Println(q.tail)
	fmt.Println(q.size)
	fmt.Println()

	fmt.Println(q.Peek())
	fmt.Println(q.Peek())
	fmt.Println(q.Peek())
	fmt.Println()

	fmt.Println(q.count)
	fmt.Println(q.nodes)
	fmt.Println(q.head)
	fmt.Println(q.tail)
	fmt.Println(q.size)

}
