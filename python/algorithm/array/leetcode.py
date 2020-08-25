
from typing import List


# 寻找数组中心位置 leetcode:724
def pivot_index(arr:List[int])->int:
    Sum = sum(arr)

    preSum = 0
    # 先求出来前半部分的和
    for i, v  in enumerate(arr):
        if preSum*2+v == Sum:
            return i
        preSum += v

    return - 1


# 搜索插入位置 leetcode: 35
def searchInsert(arr:List[int], target:int)->int:
    # 数组是有序的
    if target in arr:
        return arr.index(target)
    
    for i, v in enumerate(arr):
        if target <= v:
            return i
    # 插入到末尾的位置
    return len(arr)


# 求两数之和 leetcode: 1
def twoSum(arr:List[int], target: int):
    hashmap = {}

    # 先把数组存储起来
    for i, v in enumerate(arr):
        hashmap[v] = i

    for i, v in enumerate(arr):
        j = hashmap.get(target- v)
        if j is not None and j != i:
            return [i, j]


