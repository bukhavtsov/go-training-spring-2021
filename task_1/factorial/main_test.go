package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 1, factorial(0))
	assert.Equal(t, 2, factorial(2))
	assert.Equal(t, 6, factorial(3))
}
