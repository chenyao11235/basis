package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue(0)
	q.EnQueue(1)
	q.EnQueue(2)
	q.Print()
	assert.Equal(t, 0, q.DeQueue(), "should be 0")
	assert.Equal(t, 1, q.DeQueue(), "should be 1")
	assert.Equal(t, 2, q.DeQueue(), "should be 2")
	q.Print()
	assert.Nil(t, q.DeQueue(), "should be nil")
	q.EnQueue(100)
	q.Print()
}

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(3)
	q.EnQueue(0)
	q.EnQueue(1)
	q.EnQueue(2)
	assert.Equal(t, false, q.EnQueue(3), "queue is full")
	q.Print()
	assert.Equal(t, 0, q.DeQueue(), "should be 0")
	assert.Equal(t, 1, q.DeQueue(), "should be 1")
	q.Print()
	q.EnQueue(100)
	q.EnQueue(200)
	q.Print()
	assert.Equal(t, 2, q.DeQueue(), "should be 2")
	assert.Equal(t, 100, q.DeQueue(), "should be 100")
	assert.Equal(t, 200, q.DeQueue(), "should be 200")
}

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(4)
	q.EnQueue(0)
	q.EnQueue(1)
	q.EnQueue(2)
	assert.Equal(t, false, q.EnQueue(3), "queue is full")
	q.Print()
	assert.Equal(t, 0, q.DeQueue(), "should be 0")
	assert.Equal(t, 1, q.DeQueue(), "should be 1")
	q.Print()
	q.EnQueue(100)
	q.EnQueue(200)
	q.Print()
	assert.Equal(t, 2, q.DeQueue(), "should be 2")
	assert.Equal(t, 100, q.DeQueue(), "should be 100")
	assert.Equal(t, 200, q.DeQueue(), "should be 200")
}
