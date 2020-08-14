package array

import (
	"errors"
	"fmt"
)

//Array 数组
type Array struct {
	data   []int
	length int
}

//NewArray 新建数组
func NewArray(capacity int) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

//Len 长度
func (a *Array) Len() int {
	return a.length
}

//isIndexOutOfRange 超出索引范围
func (a *Array) isIndexOutOfRange(index int) bool {
	if index >= cap(a.data) {
		return true
	}
	return false
}

//Find 通过索引查找数组
func (a *Array) Find(index int) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index], nil
}

//Insert 插入
func (a *Array) Insert(index int, v int) error {
	if a.Len() == cap(a.data) {
		return errors.New("full array")
	}
	if index != a.length && a.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	//需要把index之后的item全部向后移动一位
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

//InsertTail 插入到末尾
func (a *Array) InsertTail(v int) error {
	return a.Insert(a.Len(), v)
}

//Delete 删除指定位置的元素
func (a *Array) Delete(index int) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := a.data[index]
	//index之后的元素全部向前移动一个位置
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}

//Print 打印数组
func (a *Array) Print() {
	var format string
	for i := int(0); i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
