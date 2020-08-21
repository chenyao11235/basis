package sort

//MergeSort 归并排序  利用的是递归的思想
func MergeSort(a []int) {
	n := len(a)
	sort(a, 0, n-1)
}

func sort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	sort(arr, start, mid)
	sort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	// 下面这两个for循环，就是如果某个数组中有剩余的元素，把剩余的元素放到tmpArr的最后
	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		j++
	}
	copy(arr[start:end+1], tmpArr)
}
