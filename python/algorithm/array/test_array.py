
from leetcode import pivot_index, searchInsert

def test_pivot_index():
    # 为毛不3呢？
   assert pivot_index([1, 7, 3, 6, 5, 6]) == 3

def test_searchInsert():
    assert searchInsert([1, 2, 3, 6, 8, 11], 3) == 2
    assert searchInsert([1, 2, 3, 6, 8, 11], 15) == 6
    assert searchInsert([1, 2, 3, 6, 8, 11],11) == 5
    assert searchInsert([1, 2, 3, 6, 8, 11],0) == 0