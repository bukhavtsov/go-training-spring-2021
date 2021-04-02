package main

import "fmt"

//Node this is a LinkedList element.
type Node struct {
	next *Node
	key  interface{}
}

//List this is Singly LinkedList.
type List struct {
	head *Node
	tail *Node
	len  int
}

//isEmpty this function determines whether the LinkedList is empty.
func (L *List) isEmpty() bool {
	return L.head == nil
}

//Insert − Adds an element at the beginning of the list.
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

//Display − Displays the complete list.
func (L *List) Display() {
	list := L.head
	for list != nil {
		fmt.Printf(" --> %v", list.key)
		list = list.next
	}
	fmt.Println()
}

//Deletion − Deletes an element at the beginning of the list.
func (L *List) Deletion() {
	if L.isEmpty() {
		fmt.Println("List is empty")
		return
	}
	L.head = L.head.next
	L.len--
}

//Search − Searches an element using the id.
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

//Delete − Deletes an element using the id.
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
