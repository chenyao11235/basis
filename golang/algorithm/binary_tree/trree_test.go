package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Add(8)
	tree.Add(3)
	tree.Add(10)
	tree.Add(1)
	tree.Add(6)
	tree.Add(13)
	tree.Add(4)
	tree.Add(7)
	tree.Add(14)

	tree.MidOrder()

	assert.Equal(t, 1, tree.FindMinValue().value, "")
	assert.Equal(t, 14, tree.FindMaxValue().value, "")
	assert.Equal(t, 6, tree.Find(6).value, "")
	assert.Equal(t, 3, tree.FindParent(6).value, "")

	tree.Delete(10)
	tree.Delete(8)
	tree.MidOrder()
}
