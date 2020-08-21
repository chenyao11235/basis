package decorator

//Component 基于同一个接口
type Component interface {
	Calc() int
}

//BasicComponent 基础对象
type BasicComponent struct {
}

//Calc 目标函数
func (c *BasicComponent) Calc() int {
	return 1
}

//MulDecorator 乘法装饰器
type MulDecorator struct {
	Component
	n int
}

//Calc 函数调用
func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.n
}

//WarpMulDecorator 装饰器函数
func WarpMulDecorator(c Component, n int) Component {
	return &MulDecorator{
		Component: c,
		n:         n,
	}
}
