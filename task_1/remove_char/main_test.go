package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveChar(t *testing.T) {
	assert.Equal(t, "loquen", removeChar("eloquent"))
	assert.Equal(t,  "ountr", removeChar("country"))
	assert.Equal(t,"erso", removeChar("person"))
	assert.Equal(t, "lac", removeChar("place"))
}
