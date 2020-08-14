package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	a := []int{10, 4, 5, 6, 1, 2, 3}
	BubbleSort(a)
	t.Log(a)
}

func TestInsertSort(t *testing.T) {
	a := []int{1000, 1, 3, 10, 5, 6, 1, 2, 3}
	InsertSort(a)
	t.Log(a)
}

func TestSelectSort(t *testing.T) {
	a := []int{1, 10, 5, 6, 1, 2, 3}
	SelectSort(a)
	t.Log(a)
}

func TestMergeSort(t *testing.T) {
	a := []int{1, 10, 5, 6, 1, 2, 3}
	MergeSort(a)
	t.Log(a)
}

func TestQuickSort(t *testing.T) {
	a := []int{1, 10, 5, 6, 1, 2, 3}
	QuickSort(a)
	t.Log(a)
}

func TestBucketSort(t *testing.T) {
	a := []int{1, 10, 5, 6, 1, 2, 3}
	BucketSort(a)
	t.Log(a)
}

func TestCountSort(t *testing.T) {
	a := []int{1, 10, 5, 6, 1, 2, 3}
	CountSort(a)
	t.Log(a)
}
