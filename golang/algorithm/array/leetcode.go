package array

/* 做过的leetcode中的题目
 */

import (
	"sort"
)

// 寻找数据的中心索引 leetcode:724
func pivotIndex(arr []int) int {
	n := len(arr)
	if n <= 2 {
		return -1
	}
	var sum = 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	var preSum = 0
	for i := 0; i < n-1; i++ {
		if preSum*2+arr[i] == sum {
			return i
		}
		preSum += arr[i]
	}
	return -1
}

// 可以使用二分查找 时间负载度是log(N) leetcode:35
func searchInsert(nums []int, target int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if target <= nums[i] {
			return i
		}
	}
	return n
}

//给出一个区间的集合，请合并所有重叠的区间。
func merge(intervals [][]int) [][]int {
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	// 把第一个数组放到merged中
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]
		c := intervals[i]

		if c[0] > m[1] {
			merged = append(merged, c)
			continue
		}

		if c[1] > m[1] {
			m[1] = c[1]
		}
	}
	return merged
}

// 两数之和 leetcode：1
func twoSum(nums []int, target int) []int {
	hashamp := map[int]int{}
	for i, v := range nums {
		hashamp[v] = i
	}
	for i, v := range nums {
		if j, ok := hashamp[target-v]; ok {
			if i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}

// 三数之和 leetcode:15
func threeSum(nums []int) [][]int {
	result := [][]int{}
	hashmap := map[int]int{}
	for i, v := range nums {
		hashmap[v] = i
	}
	return result
}
