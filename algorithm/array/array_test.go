package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	a := NewArray(7)
	a.Print()
	assert.Nil(t, a.Insert(0, 0), "no error")
	assert.Nil(t, a.Insert(1, 1), "no error")
	assert.Equal(t, 2, a.Len(), "len is 2")
	assert.Nil(t, a.Insert(2, 2), "no error")
	assert.Nil(t, a.Insert(3, 3), "no error")
	a.Print()
	assert.Nil(t, a.Insert(3, 3), "no error")
	a.Print()

	if v, err := a.Find(3); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
	a.Print()

	if v, err := a.Delete(3); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
	a.Print()
}

func TestPivotIndex(t *testing.T) {
	arr := []int{1, 7, 3, 6, 5, 6}
	assert.Equal(t, 3, pivotIndex(arr), "index should be 3")
	arr = []int{1, 7, 3, 6}
	assert.Equal(t, -1, pivotIndex(arr), "index should be 3")
}

func TestSearchInsert(t *testing.T) {
	arr := []int{1, 2, 3, 6, 8, 11}
	assert.Equal(t, 0, searchInsert(arr, 0), "should be 0")
	assert.Equal(t, 2, searchInsert(arr, 3), "should be 2")
	assert.Equal(t, 6, searchInsert(arr, 15), "should be 6")
}

func TestMerge(t *testing.T) {
	arr := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	marr := merge(arr)
	t.Log(marr)
}
