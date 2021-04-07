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
func (l *List) IsEmpty() bool {
	return l.len == 0
}

//Insert − Adds an element at the beginning of the list.
func (l *List) Insert(key interface{}) error {
	if !l.IsEmpty() && reflect.TypeOf(l.head.key) != reflect.TypeOf(key) {
		return fmt.Errorf("wrong type: %v", key)
	}
	list := &Node{
		next: l.head,
		key:  key,
	}
	l.head = list
	element := l.head
	for element.next != nil {
		element = element.next
	}
	l.tail = element
	l.len++
	return nil
}

//Display − Displays the complete list.
func (l *List) Display() error {
	if l.IsEmpty() {
		return fmt.Errorf("linked List is empty")
	}
	list := l.head
	for list != nil {
		fmt.Printf(" --> %v", list.key)
		list = list.next
	}
	fmt.Println()
	return nil
}

//Deletion − Deletes an element at the beginning of the list.
func (l *List) Deletion() error {
	if l.IsEmpty() {
		return fmt.Errorf("linked List is empty")
	}
	l.head = l.head.next
	l.len--
	return nil
}

//Search − Searches an element using the id.
func (l *List) Search(id int) (interface{}, error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("linked List is empty")
	}
	if id >= l.len || id < 0 {
		return nil, fmt.Errorf("incorrect id: %v", id)
	}
	element := l.head
	for i := 0; i < id; i++ {
		element = element.next
	}
	return element.key, nil
}

//Delete − Deletes an element using the id.
func (l *List) Delete(id int) error {
	if l.IsEmpty() {
		return fmt.Errorf("linked List is empty")
	}
	if id >= l.len || id < 0 {
		return fmt.Errorf("incorrect id: %v", id)
	}
	if id == 0 {
		_ = l.Deletion()
	} else {
		element := l.head
		for i := 0; i < id-1; i++ {
			element = element.next
		}
		element.next = element.next.next
	}
	l.len--
	return nil
}

//Sort - method sorts data.
func (l *List) Sort() {
	for i := 0; i < l.len; i++ {
		element := l.head
		for element.next != nil {
			if fmt.Sprint(element.key) > fmt.Sprint(element.next.key) {
				element.key, element.next.key = element.next.key, element.key
			}
			element = element.next
		}
	}
}
