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
func (q *Queue) Enqueue(key interface{}) error {
	if q.size == q.len {
		return fmt.Errorf("the queue is full")
	}
	if !q.IsEmpty() && reflect.TypeOf(q.tail.key) != reflect.TypeOf(key) {
		return fmt.Errorf("wrong type: %v", key)
	}
	node := &Node{
		next: q.tail,
		key:  key,
	}
	q.tail = node
	element := q.tail
	for element.next != nil {
		element = element.next
	}
	q.head = element
	q.len++
	return nil
}

//Dequeue - Remove an element from the front of the queue.
func (q *Queue) Dequeue() error {
	if q.IsEmpty() {
		return fmt.Errorf("the queue is empty")
	}
	element := q.tail
	for i := 0; i < q.len-2; i++ {
		element = element.next
	}
	q.head = element
	q.head.next = nil
	q.len--
	return nil
}

//IsEmpty - Check if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

//IsFull - Check if the queue is full.
func (q *Queue) IsFull() bool {
	return q.size <= q.len
}

//Peek - Get the value of the front of the queue without removing it.
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("the queue is empty")
	}
	return q.head.key, nil
}

//Sort - method sorts data.
func (q *Queue) Sort() {
	for i := 0; i < q.len; i++ {
		element := q.tail
		for element.next != nil {
			if fmt.Sprint(element.key) > fmt.Sprint(element.next.key) {
				element.key, element.next.key = element.next.key, element.key
			}
			element = element.next
		}
	}
}
