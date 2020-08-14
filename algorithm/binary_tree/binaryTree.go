package binarytree

import (
	"fmt"
)

//BinarySearchTreeNode 节点
type BinarySearchTreeNode struct {
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
	value int
	times int // 值出现的次数
}

func newBinarySearchTreeNode(value int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{
		value: value,
	}
}

func (n *BinarySearchTreeNode) find(value int) *BinarySearchTreeNode {
	if n.value == value {
		return n
	}
	if value < n.value {
		if n.left == nil {
			return nil
		}
		return n.left.find(value)
	} else {
		if n.right == nil {
			return nil
		}
		return n.right.find(value)
	}
}

func (n *BinarySearchTreeNode) findParent(value int) *BinarySearchTreeNode {
	if value < n.value {
		leftNode := n.left
		if leftNode == nil {
			return nil
		}
		if leftNode.value == value {
			return n
		} else {
			return leftNode.findParent(value)
		}
	} else {
		rightNode := n.right
		if rightNode == nil {
			return nil
		}
		if rightNode.value == value {
			return n
		} else {
			return rightNode.findParent(value)
		}
	}
}

func (n *BinarySearchTreeNode) add(value int) {
	// 添加到左子节点
	if value < n.value {
		if n.left == nil {
			n.left = newBinarySearchTreeNode(value)
		} else {
			// 递归
			n.left.add(value)
		}
	} else if value > n.value {
		if n.right == nil {
			n.right = newBinarySearchTreeNode(value)
		} else {
			// 递归
			n.right.add(value)
		}
	} else {
		n.times++
	}
	return
}

//findMinValue 找出最小值
func (n *BinarySearchTreeNode) findMinValue() *BinarySearchTreeNode {
	if n.left == nil {
		return n
	}
	return n.left.findMinValue()
}

//findMaxValue 找出最大值
func (n *BinarySearchTreeNode) findMaxValue() *BinarySearchTreeNode {
	if n.right == nil {
		return n
	}
	return n.right.findMaxValue()
}

// 中序遍历
func (n *BinarySearchTreeNode) midOrder() {
	if n == nil {
		return
	}

	n.left.midOrder()

	fmt.Println(n.value)

	n.right.midOrder()
}

//BinarySearchTree 二叉查找树
type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

//NewBinarySearchTree 新建
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

//Add 添加
func (t *BinarySearchTree) Add(value int) {
	if t.root == nil {
		t.root = newBinarySearchTreeNode(value)
	} else {
		t.root.add(value)
	}
}

//FindMinValue 找出最小值的节点
func (t *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if t.root == nil {
		return nil
	}
	return t.root.findMinValue()
}

//FindMaxValue 找出最大值的节点
func (t *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if t.root == nil {
		return nil
	}
	return t.root.findMaxValue()
}

//FindParent 查找父亲节点
func (t *BinarySearchTree) FindParent(value int) *BinarySearchTreeNode {
	if t.root == nil {
		return nil
	}
	// 父节点没有父节点，直接返回
	if t.root.value == value {
		return nil
	}
	return t.root.findParent(value)
}

//Find 查找指定节点
func (t *BinarySearchTree) Find(value int) *BinarySearchTreeNode {
	if t.root == nil {
		return nil
	}
	return t.root.find(value)
}

//Delete 删除节点
func (t *BinarySearchTree) Delete(value int) {
	// 第一种情况，删除的是根节点，且根节点没有儿子，直接删除即可
	// 第二种情况，删除的节点有父亲节点，但没有子树，也就是删除的是叶子节点，直接删除即可
	// 第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点，这时二叉查找树的性质又满足了。右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到
	// 第四种情况，删除的节点只有一个子树，那么该子树直接替换被删除的节点即可

	if t.root == nil {
		return
	}
	// 如果值不存在，直接返回
	node := t.Find(value)
	if node == nil {
		return
	}

	parent := t.FindParent(value)

	if parent == nil && node.left == nil && node.right == nil {
		t.root = nil
		return
	} else if parent != nil && node.left == nil && node.right == nil {
		if node.value < parent.value {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	} else if node.left != nil && node.right != nil {
		// 找到右子树中的最小节点
		rightTreeMinNode := node.right
		for rightTreeMinNode.left != nil {
			rightTreeMinNode = rightTreeMinNode.left
		}
		// 替换
		t.Delete(rightTreeMinNode.value)
		node.value = rightTreeMinNode.value
		node.times = rightTreeMinNode.times
	} else {
		// 删除的是根节点，并且根节点只有一个子树
		if parent == nil {
			if node.left != nil {
				t.root = node.left
			} else {
				t.root = node.right
			}
			return
		}

		if node.left != nil {
			if parent.left != nil && value == parent.left.value {
				parent.left = node.left
			} else {
				parent.right = node.right
			}
		} else {
			if parent.left != nil && value == parent.left.value {
				parent.left = node.right
			} else {
				parent.right = node.right
			}
		}
	}

}

//PreOrder 前序遍历
func (t *BinarySearchTree) PreOrder() {

}

//MidOrder 中序遍历
func (t *BinarySearchTree) MidOrder() {
	if t.root == nil {
		return
	}
	t.root.midOrder()
}

//InOrder 后序遍历
func (t *BinarySearchTree) InOrder() {

}
