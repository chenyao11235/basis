package queue

import (
	"fmt"
)

//CircularQueue 循环队列
type CircularQueue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

//NewCircularQueue 新建
func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		data:     make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

//IseEmpty 是否为空
func (q *CircularQueue) IseEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

//IsFull 是否满了, 这个方法是实现环形队列的关键
func (q *CircularQueue) IsFull() bool {
	if (q.tail+1)%q.capacity == q.head {
		return true
	}
	return false
}

//EnQueue 入栈
func (q *CircularQueue) EnQueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	//将元素移动到0的位置，
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

//DeQueue 出栈
func (q *CircularQueue) DeQueue() interface{} {
	if q.IseEmpty() {
		return nil
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	return v
}

//Print 打印
func (q *CircularQueue) Print() {
	if q.IseEmpty() {
		return
	}
	result := "head"
	var i = q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.data[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	fmt.Println(result)
}
