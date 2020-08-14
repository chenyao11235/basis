package isp

//Result 结果
type Result struct {
	avarage int
	min     int
	max     int
}

//Count 统计
type Count struct {
}

//count 统计
func (c *Count) count() *Result {
	return nil
}

// 一个cont函数返回了三种统计结果，但是有时候可能三个数据不会完全用到，
//这样就会造成程序资源，性能的浪费，可以把这三个操作分开来
func (c *Count) min() int {
	return 0
}

func (c *Count) avarage() int {
	return 0
}

func (c *Count) max() int {
	return 0
}
