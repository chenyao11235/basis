package sort

func getMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//BucketSort 桶排序
func BucketSort(arr []int) {
	num := len(arr)
	if num <= 1 {
		return
	}
	max := getMax(arr)
	buckets := make([][]int, num)

	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num - 1) / max
		buckets[index] = append(buckets[index], arr[i])
	}

	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}
