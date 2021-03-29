package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, "dlrow", reverse("world"))
	assert.Equal(t, "", reverse(""))
	assert.Equal(t, "h", reverse("h"))
}
