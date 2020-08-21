package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, 8, BinarySearch(arr, 8), "should be 8")
	assert.Equal(t, -1, BinarySearch(arr, 17), "should be -1")
}

func TestRecursiveBinarySearch(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, 8, RecursiveBinarySearch(arr, 8), "should be 8")
	assert.Equal(t, -1, RecursiveBinarySearch(arr, 17), "should be -1")
}

func TestSearchInsert(t *testing.T) {
	arr := []int{1, 3, 5, 6}
	assert.Equal(t, 1, searchInsert(arr, 2), "")
	assert.Equal(t, 0, searchInsert(arr, 0), "")
	assert.Equal(t, 4, searchInsert(arr, 10), "")
}
