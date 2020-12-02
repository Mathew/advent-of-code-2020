package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	result, err := parsePassword("2-5 z: zzztvz")

	assert.Equal(t, nil, err)
	assert.Equal(t, 2, result.min)
	assert.Equal(t, 5, result.max)
	assert.Equal(t, "z", result.letter)
	assert.Equal(t, "zzztvz", result.password)
}
