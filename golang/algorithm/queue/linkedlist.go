package queue

import "fmt"

//Node 节点
type Node struct {
	value interface{}
	next  *Node
}

//LinkedListQueue 用链表实现的队列
type LinkedListQueue struct {
	head   *Node // 始终指向第一个节点
	tail   *Node // 始终指向最后一个节点, 新进来的item一直保存在这个位置
	length int
}

//NewLinkedListQueue 新建
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

//EnQueue 进入队列
func (q *LinkedListQueue) EnQueue(v interface{}) {
	node := &Node{v, nil}
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

//DeQueue 出队列
func (q *LinkedListQueue) DeQueue() interface{} {
	if q.head == nil {
		return nil
	}
	v := q.head.value
	if q.head == q.tail {
		q.tail = q.head.next
	}
	q.head = q.head.next
	q.length--
	return v
}

//Print 打印队列
func (q *LinkedListQueue) Print() {
	if q.head == nil {
		fmt.Println("empty queue")
	}
	result := ""
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("%+v|", cur.value)
	}
	fmt.Println(result)
}
