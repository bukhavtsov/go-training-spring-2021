package main

import (
	"fmt"
)

type LinkedList struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

func (list *LinkedList) Insert(value interface{}) {
	node := &Node{
		next:  list.first,
		value: value,
	}
	if list.first != nil {
		list.first.prev = node
	}
	list.first = node
	for node.next != nil {
		node = node.next
	}
	list.last = node
	list.size++
}

func (list *LinkedList) Deletion() {
	if list.size > 0 {
		node := list.first.next
		list.first = node
		if list.size != 1 {
			list.first.prev = nil
		}
		list.size--
	} else {
		fmt.Println("list has not been initialized")
	}
}

func (list *LinkedList) Display() {
	firstNode := list.first
	for firstNode != nil {
		fmt.Println(firstNode.value)
		firstNode = firstNode.next
	}
}

func (list *LinkedList) Search(index int) (*Node, error) {
	if index < 0 || index > list.size-1 {
		return nil, fmt.Errorf("Index out of list")
	}
	node := list.first
	for i := 0; i != index; i++ {
		node = node.next
	}
	return node, nil
}

func (list *LinkedList) Delete(index int) {
	node, err := list.Search(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	if node.prev == nil {
		node.next.prev = nil
		list.first = node.next
	} else if node.next == nil {
		node.prev.next = nil
		list.last = node.prev
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	list.size--
}

func main() {
	linkedList := LinkedList{}
	linkedList.Insert(1)
	/*linkedList.Insert(6)
	linkedList.Insert(22)
	linkedList.Insert(3)
	linkedList.Insert(13)
	linkedList.Insert(22)
	linkedList.Insert(11)
	linkedList.Insert(65)*/
	linkedList.Display()
	linkedList.Deletion()
	fmt.Println("after deletion")
	linkedList.Display()
	linkedList.Delete(4)
	linkedList.Display()
}
