package linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestList_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	actual := list.IsEmpty()
	assert.Equal(true, actual, fmt.Sprintf("%s and %s not equal, but expected:", "true", strconv.FormatBool(actual)))
}

func TestList_Insert(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("a")
	_ = list.Insert("b")
	_ = list.Insert("c")
	_ = list.Insert("d")
	expected := 4
	actual := list.len
	assert.Equal(expected, actual, fmt.Sprintf("%d and %d not equal, but expected:", expected, actual))
}
func TestList_Insert2(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("a")
	err := list.Insert(2)
	expected := fmt.Errorf("wrong type: %v", 2)
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}

func TestList_Search(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("a")
	_ = list.Insert("b")
	_ = list.Insert("c")
	_ = list.Insert("d")
	expectedList := []string{"d", "c", "b", "a"}
	for i := 0; i < 4; i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}
func TestList_Search2(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_, err := list.Search(2)
	expected := fmt.Errorf("linked List is empty")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}
func TestList_Search3(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("a")
	_ = list.Insert("b")
	_, err := list.Search(2)
	expected := fmt.Errorf("incorrect id: %v", 2)
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}

func TestList_Delete(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("a")
	_ = list.Insert("b")
	_ = list.Insert("c")
	_ = list.Insert("c")
	_ = list.Insert("d")
	_ = list.Delete(4)
	_ = list.Delete(0)
	_ = list.Delete(1)
	expectedList := []string{"c", "b"}
	for i := 0; i < list.len; i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}
func TestList_Delete2(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	err := list.Delete(2)
	expected := fmt.Errorf("linked List is empty")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}
func TestList_Delete3(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("a")
	_ = list.Insert("b")
	err := list.Delete(2)
	expected := fmt.Errorf("incorrect id: %v", 2)
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}

func TestList_Deletion(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("a")
	_ = list.Insert("b")
	_ = list.Insert("c")
	_ = list.Insert("d")
	_ = list.Deletion()
	_ = list.Deletion()
	expectedList := []string{"b", "a"}
	for i := 0; i < list.len; i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}
func TestList_Deletion2(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	err := list.Deletion()
	expected := fmt.Errorf("linked List is empty")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}

func TestList_Display(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("a")
	_ = list.Insert("b")
	err := list.Display()
	actual := err
	assert.Equal(nil, actual, fmt.Sprintf("%v and %v not equal, but expected:", nil, actual))
}
func TestList_Display2(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	err := list.Display()
	expected := fmt.Errorf("linked List is empty")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}

func TestList_Sort(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("b")
	_ = list.Insert("a")
	_ = list.Insert("c")
	_ = list.Insert("d")
	list.Sort()
	expectedList := []string{"a", "b", "c", "d"}
	for i := 0; i < list.len; i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}
