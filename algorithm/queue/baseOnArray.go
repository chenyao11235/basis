package queue

import "fmt"

//ArrayQueue 用数组实现的队列
type ArrayQueue struct {
	data     []interface{}
	head     int
	tail     int
	capacity int
}

//NewArrayQueue 新建
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		data:     make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

//EnQueue 入队
func (q *ArrayQueue) EnQueue(v interface{}) bool {
	if q.tail == q.capacity {
		// 队列已满
		if q.head == 0 {
			return false
		}
		// 说明slice中有废弃的item，废弃的item集中在slice的前边，需要把所有的item往前移动
		for i := q.head; i < q.tail; i++ {
			q.data[i-q.head] = q.data[i]
		}
		q.tail -= q.head
		q.head = 0
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

//DeQueue 出队, 注意这里： 虽然是出队, 但是并没有真正的删除slice的item，只是把head往后边移动一个位置
func (q *ArrayQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}
	v := q.data[q.head]
	q.head++
	return v
}

//Print 打印
func (q *ArrayQueue) Print() {
	format := ""
	for i := q.head; i < q.tail; i++ {
		format += fmt.Sprintf("|%+v", q.data[i])
	}
	fmt.Println(format)
}
