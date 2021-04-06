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

func TestList_Delete(t *testing.T) {
	assert := assert.New(t)
	list := List{}
	_ = list.Insert("a")
	_ = list.Insert("b")
	_ = list.Insert("c")
	_ = list.Insert("d")
	_ = list.Delete(0)
	_ = list.Delete(1)
	expectedList := []string{"c", "a"}
	for i := 0; i < list.len; i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
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

//func TestList_Sort(t *testing.T) {
//	assert := assert.New(t)
//	list := List{}
//	_ = list.Insert("b")
//	_ = list.Insert("a")
//	_ = list.Insert("c")
//	_ = list.Insert("d")
//	_ = list.Display()
//	list.Sort()
//	_ = list.Display()
//	expectedList := []string{"a", "b", "c", "d"}
//	for i := 0; i < list.len; i++ {
//		expected := expectedList[i]
//		actual, _ := list.Search(i)
//		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
//	}
//}
