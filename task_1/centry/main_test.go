package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCentury(t *testing.T) {
	assert.Equal(t,  18, century(1705))
	assert.Equal(t, 19, century(1900))
	assert.Equal(t, 17, century(1601))
	assert.Equal(t, 20, century(2000))
	assert.Equal(t, 1, century(89))
}
