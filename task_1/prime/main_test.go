package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	assert.Equal(t, false, isPrime(0))
	assert.Equal(t, false, isPrime(1))
	assert.Equal(t, false, isPrime(75))
	assert.Equal(t, false, isPrime(-1))
	assert.Equal(t, false, isPrime(4))
	assert.Equal(t, false, isPrime(6))
	assert.Equal(t, false, isPrime(8))
	assert.Equal(t, false, isPrime(9))
	assert.Equal(t, false, isPrime(45))
	assert.Equal(t, false, isPrime(-5))
	assert.Equal(t, false, isPrime(-8))
	assert.Equal(t, false, isPrime(-41))

	assert.Equal(t, true, isPrime(2))
	assert.Equal(t, true, isPrime(73))
	assert.Equal(t, true, isPrime(3))
	assert.Equal(t, true, isPrime(5))
	assert.Equal(t, true, isPrime(7))
	assert.Equal(t, true, isPrime(41))
	assert.Equal(t, true, isPrime(5099))
}
