package heap

import (
	"github.com/pkg/errors"
)

//Heap 堆
type Heap struct {
	data     []int
	capacity int
	length   int
}

//NewHeap 新建
func NewHeap(capacity int) *Heap {
	return &Heap{
		data:     make([]int, capacity),
		capacity: capacity,
		length:   0,
	}
}

//Insert 往堆中插入一个元素
func (h *Heap) Insert(value int) error {
	if h.length >= h.capacity {
		return errors.New("heap is full")
	}
	h.length++
	// 注意 index为0的位置是不存储数据的
	h.data[h.length] = value

	// 堆化，自下而上
	index := h.length - 1

	for index/2 > 0 && h.data[index] > h.data[index/2] {
		// swap , 与父节点比较，交换
		h.data[index], h.data[index/2] = h.data[index/2], h.data[index]
		index = index / 2
	}
	return nil
}

//RemoveTop 删除堆顶元素
func (h *Heap) RemoveTop() {
	// 将堆中的最后一个元素置于堆顶，然后从上往下进行堆化
	if h.length == 0 {
		return
	}

	if h.length == 1 {
		return
	}
	// 把最后一个元素置于堆顶
	h.data[1] = h.data[h.length]
	h.length--
	// 然后从上往下进行堆化
	heapifyUpToDown(h.data, h.length)
}

// 堆化 自上而下
func heapifyUpToDown(arr []int, length int) {
	index := 1
	var position int
	for {
		// 比较左子节点
		if index*2 <= length && arr[index] < arr[index*2] {
			position = index * 2
		}
		// 比较右子节点
		if index*2+1 <= length && arr[index] < arr[index*2+1] {
			position = index*2 + 1
		}
		// 直到遍历到最后一个节点就退出
		if index >= length {
			break
		}
		// 交换值
		arr[index], arr[position] = arr[position], arr[index]
		index = position
	}
}
