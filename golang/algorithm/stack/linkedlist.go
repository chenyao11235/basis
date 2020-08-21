package stack

import (
	"fmt"
)

//Node 节点
type Node struct {
	value interface{}
	next  *Node
}

//StackBaseOnLinkedList 基于链表实现的栈
type StackBaseOnLinkedList struct {
	topNode *Node
}

//NewStackBaseOnLinkedList 新建stack
func NewStackBaseOnLinkedList() *StackBaseOnLinkedList {
	return &StackBaseOnLinkedList{nil}
}

//Push 压栈
func (s *StackBaseOnLinkedList) Push(v interface{}) {
	s.topNode = &Node{
		value: v,
		next:  s.topNode,
	}
}

//Pop 出栈
func (s *StackBaseOnLinkedList) Pop() interface{} {
	if s.topNode == nil {
		return nil
	}
	v := s.topNode.value
	s.topNode = s.topNode.next
	return v
}

//IsEmpty 判断是否为空
func (s *StackBaseOnLinkedList) IsEmpty() bool {
	if s.topNode == nil {
		return true
	}
	return false
}

//Top 查看栈顶的元素，会返回但是不会删除
func (s *StackBaseOnLinkedList) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.topNode.value
}

//Flush 清空
func (s *StackBaseOnLinkedList) Flush() {
	s.topNode = nil
}

//Print 打印
func (s *StackBaseOnLinkedList) Print() {
	cur := s.topNode

	for cur != nil {
		fmt.Println(cur.value)
		cur = cur.next
	}
}
