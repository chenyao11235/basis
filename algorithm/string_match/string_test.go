package stringmatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBF(t *testing.T) {
	assert.Equal(t, true, BFMatch("hello", "he"), "")
	assert.Equal(t, false, BFMatch("hello", "hi"), "")
	assert.Equal(t, true, BFMatch("abcdcba", "cdc"), "")
	assert.Equal(t, true, BFMatch("abcDcB", "cB"), "")
}
