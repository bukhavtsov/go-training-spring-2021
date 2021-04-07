package linked_list

import (
	"fmt"
	"reflect"
)

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
func (L *List) IsEmpty() bool {
	return L.len == 0
}

//Insert − Adds an element at the beginning of the list.
func (L *List) Insert(key interface{}) error {
	if !L.IsEmpty() && reflect.TypeOf(L.head.key) != reflect.TypeOf(key) {
		return fmt.Errorf("wrong type: %v", key)
	}
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
	return nil
}

//Display − Displays the complete list.
func (L *List) Display() error {
	if L.IsEmpty() {
		return fmt.Errorf("linked List is empty")
	}
	list := L.head
	for list != nil {
		fmt.Printf(" --> %v", list.key)
		list = list.next
	}
	fmt.Println()
	return nil
}

//Deletion − Deletes an element at the beginning of the list.
func (L *List) Deletion() error {
	if L.IsEmpty() {
		return fmt.Errorf("linked List is empty")
	}
	L.head = L.head.next
	L.len--
	return nil
}

//Search − Searches an element using the id.
func (L *List) Search(id int) (interface{}, error) {
	if L.IsEmpty() {
		return nil, fmt.Errorf("linked List is empty")
	}
	if id >= L.len || id < 0 {
		return nil, fmt.Errorf("incorrect id: %v", id)
	}
	element := L.head
	for i := 0; i < id; i++ {
		element = element.next
	}
	return element.key, nil
}

//Delete − Deletes an element using the id.
func (L *List) Delete(id int) error {
	if L.IsEmpty() {
		return fmt.Errorf("linked List is empty")
	}
	if id >= L.len || id < 0 {
		return fmt.Errorf("incorrect id: %v", id)
	}
	if id == 0 {
		_ = L.Deletion()
	} else {
		element := L.head
		for i := 0; i < id-1; i++ {
			element = element.next
		}
		element.next = element.next.next
	}
	L.len--
	return nil
}

//Sort - method sorts data.
func (L *List) Sort() {
	for i := 0; i < L.len; i++ {
		element := L.head
		for element.next != nil {
			if fmt.Sprint(element.key) > fmt.Sprint(element.next.key) {
				element.key, element.next.key = element.next.key, element.key
			}
			element = element.next
		}
	}
}
