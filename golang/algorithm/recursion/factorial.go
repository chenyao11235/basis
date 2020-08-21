package recursion

//Fac 迭代实现阶乘
type Fac struct {
	val map[int]int
}

//NewFactorial 新建
func NewFactorial(n int) *Fac {
	return &Fac{
		make(map[int]int, n),
	}
}

//Factorial 阶乘计算
func (fac *Fac) Factorial(n int) int {
	if fac.val[n] != 0 {
		return fac.val[n]
	}

	if n <= 1 {
		fac.val[n] = 1
		return 1
	} else {
		res := n * fac.Factorial(n-1)
		fac.val[n] = res
		return res
	}
}

//Print 打印
func (fac *Fac) Print(n int) {
	println(fac.val[n])
}
