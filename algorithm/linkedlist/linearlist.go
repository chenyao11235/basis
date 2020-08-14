package linearlist

import "fmt"

//Node 单向链表节点
type Node struct {
	value interface{}
	next  *Node
}

//GetValue 获取节点的值
func (n *Node) GetValue() interface{} {
	return n.value
}

//GetNext 获取下一个节点
func (n *Node) GetNext() *Node {
	return n.next
}

//LinearList 单向链表
type LinearList struct {
	head   *Node // 带头链表，链表中的第一个节点是head.next
	length int
}

//NewNode 创建一个单向链表节点
func NewNode(value interface{}) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

//NewLinearList 创建一个单向链表
func NewLinearList() *LinearList {
	return &LinearList{NewNode(0), 0}
}

//insertAfter 在某个节点之后插入一个节点
func (s *LinearList) insertAfter(p *Node, value interface{}) bool {
	newNode := NewNode(value)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	s.length++
	return true
}

//insertBefore 在某个节点之前插入一个节点
func (s *LinearList) insertBefore(p *Node, value interface{}) bool {
	if p == s.head {
		return false
	}
	cur := s.head.next
	pre := s.head
	// 找到p节点的前一个节点
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 链表中只有一个节点(头节点)
	if nil == cur {
		return false
	}
	newNode := NewNode(value)
	pre.next = newNode
	newNode.next = cur
	s.length++
	return true
}

//InsterHead 在头部插入节点
func (s *LinearList) InsterHead(value interface{}) bool {
	return s.insertAfter(s.head, value)
}

//InsertTail 在尾部插入节点
func (s *LinearList) InsertTail(value interface{}) bool {
	curl := s.head
	for curl.next != nil {
		curl = curl.next
	}
	return s.insertAfter(curl, value)
}

//FindByIndex 查找指定值的索引
func (s *LinearList) FindByIndex(index int) *Node {
	if index >= s.length {
		return nil
	}

	cur := s.head.next
	for i := 0; i < index; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.next
	}
	return cur
}

//DeleteNode 删除节点
func (s *LinearList) DeleteNode(p *Node) bool {
	if p == nil {
		return false
	}

	pre := s.head
	cur := s.head.next

	for cur != nil {
		if cur == p {
			pre.next = p.next
			s.length--
			return true
		}
		pre = cur
		cur = cur.next
	}
	return false
}

//Print 打印链表
func (s *LinearList) Print() {
	cur := s.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("=>%+v", cur.GetValue())
		cur = cur.next
	}
	fmt.Println(format)
}

//Reverse 链表反转
func (s *LinearList) Reverse() {
	if s.head == nil || s.head.next == nil || s.head.next.next == nil {
		return
	}
	// 遍历链表，直到pre指向链表的最后一个节点
	var pre *Node = nil
	cur := s.head.next
	// 这三个变量的关系大概是: cur.next = pre  tmp在遍历的时候是走在最前面的
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	s.head.next = pre
}

//HashCycle 判断链表时候有环(两个节点的next互相指向对方)
func (s *LinearList) HashCycle() bool {
	if s.head != nil {
		slow := s.head.next
		fast := s.head.next
		for slow == nil || fast == nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

//MergeSortedLinearList 合并两个有序链表
func (s *LinearList) MergeSortedLinearList(l1, l2 *LinearList) *LinearList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}
	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}
	l := &LinearList{head: &Node{}}
	cur := l.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) > cur2.value.(int) {
			cur.next = cur2
			cur2 = cur2.next
		} else {
			cur.next = cur1
			cur1 = cur1.next
		}
		cur = cur.next
	}
	if cur1 != nil {
		cur.next = cur1
	} else if cur2 != nil {
		cur.next = cur2
	}
	return l
}

//DeleteBottomN 删除倒数第n个元素： 用两个指针，通过n设计好两个指针之间的差多少个节点，然后遍历整个链表
func (s *LinearList) DeleteBottomN(n int) {
	if s == nil || s.head.next == nil {
		return
	}
	pre := s.head
	for i := 1; i <= n && pre != nil; i++ {
		pre = pre.next
	}
	if nil == pre {
		return
	}
	cur := s.head
	// 直到遍历到最后一个节点
	for pre.next != nil {
		pre = pre.next
		cur = cur.next
	}
	cur.next = cur.next.next
}

//FindMiddleNode 获取中间节点
func (s *LinearList) FindMiddleNode() *Node {
	if nil == s.head || nil == s.head.next {
		return nil
	}
	if nil == s.head.next.next {
		return s.head.next
	}

	slow, fast := s.head, s.head
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
