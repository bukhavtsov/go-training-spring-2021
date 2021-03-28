package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	assert.Equal(t, 6, compute(30, 12))
	assert.Equal(t, 1, compute(8, 9))
	assert.Equal(t, 1, compute(1, 1))
}
