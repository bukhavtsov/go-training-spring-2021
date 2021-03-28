package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenOrOdd(t *testing.T) {
	assert.Equal(t, "Even", evenOrOdd(2))
	assert.Equal(t, "Even", evenOrOdd(0))
	assert.Equal(t, "Odd", evenOrOdd(7))
	assert.Equal(t, "Odd", evenOrOdd(1))
}
