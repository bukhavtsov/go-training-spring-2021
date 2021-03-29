package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShort(t *testing.T)  {
	assert.Equal(t, 3, findShort("bitcoin take over the world maybe who knows perhaps"))
	assert.Equal(t,  3, findShort("turns out random test cases are easier than writing out basic ones"))
}
