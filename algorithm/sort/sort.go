package sort

import "testing"

//BubbleSort 冒泡排序
func BubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		//如果某次冒泡没有发生数据交换，说明已经完全有序了
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

//InsertSort 插入排序 从未排序区拿到一个元素插入到已排序区和合适位置
func InsertSort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

//SelectSort 选择排序, 从未排序去找出最小的数据放到已排序区的末尾
func SelectSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

//FastSort 快速排序
func FastSort(t *testing.T) {

}
