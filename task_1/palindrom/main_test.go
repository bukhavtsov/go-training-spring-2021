package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome("anna"))
	assert.Equal(t, false, isPalindrome("walter"))
	assert.Equal(t, true, isPalindrome("12321"))
	assert.Equal(t, false, isPalindrome("123456"))
	assert.Equal(t, true, isPalindrome("."))
	assert.Equal(t, true, isPalindrome(".!!."))
}
