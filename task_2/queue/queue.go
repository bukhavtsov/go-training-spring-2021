package queue

import (
	"fmt"
	"reflect"
)

//Node this is a Queue element.
type Node struct {
	next *Node
	key  interface{}
}

//Queue is a data structure that, like a stack, has restrictions on adding and removing elements.
type Queue struct {
	head *Node
	tail *Node
	size int
	len  int
}

//NewQueue - Creates a new queue
func NewQueue(size int) *Queue {
	return &Queue{
		size: size,
	}
}

//Enqueue - Add an element to the end of the queue.
func (Q *Queue) Enqueue(key interface{}) error {
	if Q.size == Q.len {
		return fmt.Errorf("the queue is full")
	}
	if !Q.IsEmpty() && reflect.TypeOf(Q.tail.key) != reflect.TypeOf(key) {
		return fmt.Errorf("wrong type: %v", key)
	}
	node := &Node{
		next: Q.tail,
		key:  key,
	}
	Q.tail = node
	element := Q.tail
	for element.next != nil {
		element = element.next
	}
	Q.head = element
	Q.len++
	return nil
}

//Dequeue - Remove an element from the front of the queue.
func (Q *Queue) Dequeue() error {
	if Q.IsEmpty() {
		return fmt.Errorf("the queue is empty")
	}
	element := Q.tail
	for i := 0; i < Q.len-2; i++ {
		element = element.next
	}
	Q.head = element
	Q.head.next = nil
	Q.len--
	return nil
}

//IsEmpty - Check if the queue is empty.
func (Q *Queue) IsEmpty() bool {
	return Q.len == 0
}

//IsFull - Check if the queue is full.
func (Q *Queue) IsFull() bool {
	return Q.size <= Q.len
}

//Peek - Get the value of the front of the queue without removing it.
func (Q *Queue) Peek() (interface{}, error) {
	if Q.IsEmpty() {
		return nil, fmt.Errorf("the queue is empty")
	}
	return Q.head.key, nil
}

//Sort - method sorts data.
func (Q *Queue) Sort() {
	for i := 0; i < Q.len; i++ {
		element := Q.tail
		for element.next != nil {
			if fmt.Sprint(element.key) > fmt.Sprint(element.next.key) {
				element.key, element.next.key = element.next.key, element.key
			}
			element = element.next
		}
	}
}
