package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackBaseOnLinkedList(t *testing.T) {
	s := NewStackBaseOnLinkedList()
	assert.Equal(t, true, s.IsEmpty(), "should be empty")
	assert.Nil(t, s.Pop(), "pop should be nil")
	s.Push(1)
	assert.Equal(t, false, s.IsEmpty(), "should not be empty")
	assert.Equal(t, 1, s.Pop(), "value = 1")
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Pop(), "value = 3")
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Print()
}

func TestStackBaseOnArray(t *testing.T) {
	s := NewStackBaseOnArray()
	assert.Equal(t, true, s.IsEmpty(), "should be empty")
	assert.Nil(t, s.Pop(), "pop should be nil")
	s.Push(1)
	assert.Equal(t, false, s.IsEmpty(), "should not be empty")
	assert.Equal(t, 1, s.Pop(), "value = 1")
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Pop(), "value = 3")
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Print()
}
