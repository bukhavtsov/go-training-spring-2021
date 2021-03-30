package main

import "fmt"

type Node struct {
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
	len  int
}

func (L *List) isEmpty() bool {
	return L.head == nil
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	L.head = list
	element := L.head
	for element.next != nil {
		element = element.next
	}
	L.tail = element
	L.len++
}

func (L *List) Display() {
	list := L.head
	for list != nil {
		fmt.Printf(" --> %v", list.key)
		list = list.next
	}
	fmt.Println()
}
func (L *List) Deletion() {
	if L.isEmpty() {
		fmt.Println("List is empty")
		return
	}
	L.head = L.head.next
	L.len--
}

func (L *List) Search(id int) interface{} {
	if id >= L.len || id < 0 {
		return nil
	}
	element := L.head
	for i := 0; i < L.len-id-1; i++ {
		element = element.next
	}
	return element.key
}

func (L *List) Delete(id int) {
	if L.isEmpty() {
		fmt.Println("List is empty")
		return
	}
	if id >= L.len || id < 0 {
		return
	}
	element := L.head
	for i := 0; i < L.len-id-2; i++ {
		element = element.next
	}
	element.next = element.next.next
	L.len--
}

func main() {
	link := List{}
	link.Insert(1)
	link.Insert(2)
	link.Insert(3)
	link.Insert(4)
	link.Insert(5)
	link.Insert(6)
	link.Insert(7)
	link.Insert(8)
	link.Insert(9)
	link.Display()
	fmt.Println(link.Search(5))
	link.Delete(5)
	link.Display()
	fmt.Println(link.len)
}
