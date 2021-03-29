package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	assert.Equal(t, true, isAnagram("foefet", "toffee"))
	assert.Equal(t, true, isAnagram("Buckethead", "DeathCubeK"))
	assert.Equal(t, true, isAnagram("Twoo", "WooT"))
	assert.Equal(t, false, isAnagram("dumble", "WooT"))
	assert.Equal(t, false, isAnagram("ound", "round"))
	assert.Equal(t, false, isAnagram("apple", "pale"))
}
