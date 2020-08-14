package stack

import (
	"fmt"
)

//StackBaseOnArray 基于数组实现的栈
type StackBaseOnArray struct {
	data []interface{}
	top  int //保存的是索引的位置
}

//NewStackBaseOnArray 新建一个stack
func NewStackBaseOnArray() *StackBaseOnArray {
	return &StackBaseOnArray{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

//Push 压栈
func (s *StackBaseOnArray) Push(v interface{}) {
	s.top++
	// 因为在pop操作中并没有真正的删除item，在这里通过覆盖操作来实现删除
	// 如果top<=len-1 说明有本应该废弃的元素，并且该元素位于栈顶，就应该覆盖
	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

//Pop 出栈
func (s *StackBaseOnArray) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top--
	return v
}

//IsEmpty 判断是否为空
func (s *StackBaseOnArray) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

//Top 取栈顶
func (s *StackBaseOnArray) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

//Flush 清空
func (s *StackBaseOnArray) Flush() {
	s.top = -1
}

//Print 打印
func (s *StackBaseOnArray) Print() {
	if s.IsEmpty() {
		return
	}
	for i := s.top; i >= 0; i-- {
		fmt.Println(s.data[i])
	}
}
