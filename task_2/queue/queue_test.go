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
func TestQueue_Enqueue2(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(6)
	_ = q.Enqueue("one")
	err := q.Enqueue(2)
	expected := fmt.Errorf("wrong type: %v", 2)
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}
func TestQueue_Enqueue3(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(5)
	_ = q.Enqueue("one")
	_ = q.Enqueue("two")
	_ = q.Enqueue("three")
	_ = q.Enqueue("four")
	_ = q.Enqueue("five")
	err := q.Enqueue("six")
	expected := fmt.Errorf("the queue is full")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))

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
func TestQueue_Dequeue2(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(5)
	err := q.Dequeue()
	expected := fmt.Errorf("the queue is empty")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
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
func TestQueue_Peek2(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(5)
	_, err := q.Peek()
	expected := fmt.Errorf("the queue is empty")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}

func TestList_Sort(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(6)
	_ = q.Enqueue("d")
	_ = q.Enqueue("c")
	_ = q.Enqueue("a")
	_ = q.Enqueue("b")
	_ = q.Sort()
	expectedList := []string{"a", "b", "c", "d"}
	for i := 0; i < q.len; i++ {
		expected := expectedList[i]
		element := q.tail
		for j := 0; j < i; j++ {
			element = element.next
		}
		actual := element.key
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}

func TestList_Sort2(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(6)
	_ = q.Enqueue("d")
	err := q.Sort()
	expected := fmt.Errorf("the queue contains one element")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}

func TestList_Sort3(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue(6)
	err := q.Sort()
	expected := fmt.Errorf("the queue is empty")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}
