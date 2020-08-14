package recursion

import "fmt"

//Fibs 递归实现斐波那契数列
type Fibs struct {
	val map[int]int // 使用字典存储结果，用来优化性能，避免重复计算
}

//NewFibs 新建
func NewFibs(n int) *Fibs {
	return &Fibs{
		make(map[int]int, n),
	}
}

//Fibonacci 斐波那契计算
func (fib *Fibs) Fibonacci(n int) int {
	if fib.val[n] != 0 {
		return fib.val[n]
	}
	if n <= 1 {
		fib.val[1] = 1
		return 1
	} else if n == 2 {
		fib.val[2] = 1
		return 1
	} else {
		res := fib.Fibonacci(n-1) + fib.Fibonacci(n-2)
		fib.val[n] = res
		return res
	}
}

//Print 打印  递归过程中已经把每次计算出来的值保存起来了
func (fib *Fibs) Print(n int) {
	fmt.Println(fib.val[n])
}
