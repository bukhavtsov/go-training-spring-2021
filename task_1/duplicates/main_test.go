package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDuplicate(t *testing.T) {
	assert.Equal(t, getDuplicate([]int{1, 1, 2, 3}), 1)
	assert.Equal(t, getDuplicate([]int{5, 4, 3, 2, 1, 1}), 1)
	assert.Equal(t, getDuplicate([]int{1, 3, 2, 5, 4, 5, 7, 6}), 5)
	assert.Equal(t, getDuplicate([]int{8, 2, 6, 3, 7, 2, 5, 1, 4}), 2)
}
