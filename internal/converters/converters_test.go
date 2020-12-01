package converters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToInt(t *testing.T) {
	r, err := StringToInt("6")
	assert.Equal(t, 6, r, err)
}

func TestStringsToInts(t *testing.T) {
	r, err := StringsToInts("6", "5", "3")
	assert.Equal(t, []int{6, 5, 3}, r, err)
}
