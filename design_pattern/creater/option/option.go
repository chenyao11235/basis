package option

import "fmt"

// 当阅读golang源码的时候，常常会遇见的一种设计模式，那就是选项模式

// 1. 需要为结构设置的某些字段设置默认值
// 2. 结构体的成员的值会发生变动

// 这种设计模式是golang中常用的一种设计模式，在其他语言可能不是很常见

// 优点： 支持传递多个参数，并且在参数个数、类型发生变化时保持兼容性
//任意顺序传递参数
//支持默认值
//方便拓展
type Option struct {
	A string
	B string
	C string
}

type OptionFunc func(option *Option)

// 用于设置不同字段的函数
func WithA(a string) OptionFunc {
	return func(option *Option) {
		option.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(option *Option) {
		option.B = b
	}
}

func WithC(c string) OptionFunc {
	return func(option *Option) {
		option.C = c
	}
}

var (
	defaultOption = &Option{
		A: "",
		B: "",
		C: "",
	}
)

//NewOption 工厂函数
func NewOption(opts ...OptionFunc) *Option {
	opt := defaultOption
	for _, o := range opts {
		o(opt)
	}

	return opt
}

func main() {
	opt := NewOption(
		WithA("hello"),
		WithB("world"),
		WithC("!"),
	)
	fmt.Println(opt)
}
