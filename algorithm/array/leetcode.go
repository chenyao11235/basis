package array

/* 做过的leetcode中的题目
 */

import "sort"

// 寻找数据的中心索引
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
	for i := 0; i < n; i++ {
		if preSum*2+arr[i] == sum {
			return i
		}
		preSum += arr[i-1]
	}
	return -1
}

// 搜索插入位置 在一个排序数组中找到指定元素的索引，如果不存在就返回应该插入的位置 这种是暴力解题 时间复杂度是O(N)
// 可以使用二分查找 时间负载度是log(N)
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

//矩阵旋转

/*
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
], 旋转成为
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
*/

func rotate(matrix [][]int) {

}
