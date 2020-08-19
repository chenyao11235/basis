package backtracking

//0-1背包问题
/*
我们有一个背包，背包总的承载重量是 Wkg。
现在我们有 n 个物品，每个物品的重量不等，并且不可分割。
我们现在期望选择几件物品，装载到背包中。
在不超过背包所能装载重量的前提下，如何让背包中物品的总重量最大？
*/

const MAX_WEIGHT = 100

// i
func cal(i, cw int, items []int, n, w int) {
	if i == n || cw >= w {
		if cw > MAX_WEIGHT {
			cw = MAX_WEIGHT
		}
		return
	}

	cal(i+1, cw, items, n, w)

	if cw+items[i] <= w {
		cal(i+1, cw+items[i], items, n, w)
	}
}
