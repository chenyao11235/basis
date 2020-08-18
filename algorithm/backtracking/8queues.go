package backtracking

import (
	"fmt"
)

//八皇后问题
//有一个8x8棋盘，有八个皇后（棋子），每个棋子所在的行，列，对角线都不能有另外一个棋子，找出所有满足这种要求的防棋子的方式

//用这个一维数组来表示8皇后怎么存放
//[0, 4, 7, 5, 3, 6, 2, 3] 表示第1行放在第0列，第2行放在第4列，第3行放在第7列......
var result = make([]int, 8)

// 通过一维数组把8皇后问题打印出来
func printQueues(result []int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// 计算每一行的皇后应该放在第几列，调用方法calQueues(0)
func calQueues(row int) {
	// 说明0-7行都已经放好了，已经完事了
	if row == 8 {
		printQueues(result)
		return
	}
	for column := 0; column < 8; column++ {
		if isOK(row, column) {
			result[row] = column
			// 计算下一行
			calQueues(row + 1)
		}
	}
}

func isOK(row, column int) bool {
	leftUp := column - 1  // 左上对角线
	rightUp := column + 1 // 右上对角线
	// 逐行往上检查
	for i := row - 1; i >= 0; i-- {
		if result[i] == column {
			return false
		}
		if leftUp >= 0 {
			if result[i] == leftUp {
				return false
			}
		}
		if rightUp < 8 {
			if result[i] == rightUp {
				return false
			}
		}
		leftUp--
		rightUp++
	}
	return true
}
