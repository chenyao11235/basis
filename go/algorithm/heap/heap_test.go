package heap

import "testing"

func TestHeap(t *testing.T) {
	heap := NewHeap(15)
	heap.Insert(5)
	heap.Insert(7)
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(6)
	heap.Insert(4)

	heap.RemoveTop()
}

func TestHeapSort(t *testing.T) {
	arr := []int{7, 5, 19, 8, 4, 1, 20, 13, 16}
	HeapSort(arr, 9)
	for _, v := range arr {
		t.Log(v)
	}
}
