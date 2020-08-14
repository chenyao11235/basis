package sort

//QuickSort 快速排序
func QuickSort(arr []int) {
	n := len(arr)
	separate(arr, 0, n-1)
}

func separate(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separate(arr, start, i-1)
	separate(arr, i+1, end)

}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
