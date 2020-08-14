package binarysearch

//BinarySearch 二分查找  使用for循环实现
func BinarySearch(arr []int, value int) int {
	n := len(arr)
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == value {
			return arr[mid]
		} else if value < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//RecursiveBinarySearch 递归实现二分查找
func RecursiveBinarySearch(arr []int, value int) int {
	n := len(arr)
	return bs(arr, 0, value, n-1)
}

func bs(arr []int, low, value, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if value == arr[mid] {
		return arr[mid]
	} else if arr[mid] > value {
		return bs(arr, low, value, mid-1)
	} else {
		return bs(arr, mid+1, value, high)
	}
}

// 在循序数组中搜索元素插入位置  如果元素在数组中存在就返回索引，不存在就返回其应该插入位置的索引
func searchInsert(nums []int, target int) int {
	n := len(nums)
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
