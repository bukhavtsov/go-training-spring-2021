package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUsername(t *testing.T) {
	assert.Equal(t, true, isUsername("asddsa"))
	assert.Equal(t, false, isUsername("a"))
	assert.Equal(t, false, isUsername("Hass"))
	assert.Equal(t, false, isUsername("Hasd_12assssssasasasasasaasasasasas"))
	assert.Equal(t, false, isUsername(""))
	assert.Equal(t, true, isUsername("____"))
	assert.Equal(t, false, isUsername("012"))
	assert.Equal(t, true, isUsername("p1pp1"))
	assert.Equal(t, false, isUsername("asd43 34"))
	assert.Equal(t, true, isUsername("asd43_34"))
}
