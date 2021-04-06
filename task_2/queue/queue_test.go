package queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNewQueue(t *testing.T) {
	assert := assert.New(t)
	expected := &Queue{
		size: 6,
	}
	actual := NewQueue(6)
	assert.Equal(expected, actual, "not equal, but expected:")
}

func TestQueue_Enqueue(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(6)
	_ = q.Enqueue("one")
	_ = q.Enqueue("two")
	_ = q.Enqueue("three")
	_ = q.Enqueue("four")
	_ = q.Enqueue("five")
	_ = q.Enqueue("six")
	expectedQueue := []string{"one", "two", "three", "four", "five", "six"}
	for i, expected := range expectedQueue {
		element := q.tail
		for j := 1; j < q.len-i; j++ {
			element = element.next
		}
		actual := element.key
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}

func TestQueue_Dequeue(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(6)
	_ = q.Enqueue("one")
	_ = q.Enqueue("two")
	_ = q.Enqueue("three")
	_ = q.Enqueue("four")
	_ = q.Enqueue("five")
	_ = q.Enqueue("six")
	_ = q.Dequeue()
	_ = q.Dequeue()
	expectedQueue := []string{"three", "four", "five", "six"}
	for i, expected := range expectedQueue {
		element := q.tail
		for j := 1; j < q.len-i; j++ {
			element = element.next
		}
		actual := element.key
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(6)
	actual := q.IsEmpty()
	assert.Equal(true, actual, fmt.Sprintf("%s and %s not equal, but expected:", "true", strconv.FormatBool(actual)))
}

func TestQueue_IsFull(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(6)
	_ = q.Enqueue("one")
	_ = q.Enqueue("two")
	_ = q.Enqueue("three")
	_ = q.Enqueue("four")
	_ = q.Enqueue("five")
	_ = q.Enqueue("six")
	actual := q.IsFull()
	assert.Equal(true, actual, fmt.Sprintf("%s and %s not equal, but expected:", "true", strconv.FormatBool(actual)))
}

func TestQueue_Peek(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(6)
	_ = q.Enqueue("one")
	_ = q.Enqueue("two")
	_ = q.Enqueue("three")
	_ = q.Enqueue("four")
	_ = q.Enqueue("five")
	_ = q.Enqueue("six")
	actual, _ := q.Peek()
	assert.Equal("one", actual, fmt.Sprintf("%s and %s not equal, but expected:", "one", actual))
}

//func TestList_Sort(t *testing.T) {
//	assert := assert.New(t)
//	list := List{}
//	_ = list.Insert("a")
//	_ = list.Insert("b")
//	_ = list.Insert("c")
//	_ = list.Insert("d")
//	list.Sort()
//	expectedList := []string{"a", "b", "c", "d"}
//	for i := 0; i < list.len; i++ {
//		expected := expectedList[i]
//		element := list.head
//		for j := 0; j < i; j++ {
//			element = element.next
//		}
//		fmt.Println(element.key)
//		actual := element.key
//		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
//	}
//}
