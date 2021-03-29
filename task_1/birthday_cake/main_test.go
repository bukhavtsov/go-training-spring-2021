package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCake(t *testing.T) {
	assert.Equal(t, "That was close!", cake(900, "abc"))
	assert.Equal(t, "That was close!", cake(739, "coatckhs"))
	assert.Equal(t, "That was close!", cake(517, "pbbi"))
	assert.Equal(t, "That was close!", cake(721, "qprffxgz"))
	assert.Equal(t, "That was close!", cake(932, "bbgpawya"))

	assert.Equal(t, "Fire!", cake(256, "aaaaaddddr"))
	assert.Equal(t, "Fire!", cake(56, "ifkhchlhfd"))
}
