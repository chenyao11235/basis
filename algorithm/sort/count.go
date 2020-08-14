package sort

import "math"

//CountSort 计数排序
func CountSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	var max int = math.MinInt32

	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	// 把arr中的value作为c中的下标，c的value是arr中value的个数， 逐个元素统计
	c := make([]int, max+1)
	for i := range arr {
		c[arr[i]]++
	}
	// 统计c中小于等于自己的(排在自己前面的元素)有多少个
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}
	// 从arr中取出value，这个value是c的下标，通过这个下标从c中查到该value前面有多少个item
	// value有了，这个value应该排第几也直到了，然后把这个value放到r中
	// 把c中的该value的统计数量减1
	r := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		index := c[arr[i]] - 1
		r[index] = arr[i]
		c[arr[i]]--
	}
	copy(arr, r)
}
