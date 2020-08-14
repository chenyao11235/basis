package decorator

import (
	"fmt"
	"strconv"
)

//Decoer 装饰函数和被装饰函数都应该属于同一种函数
type Decoer func(i int, s string) bool

//Foo 被装饰函数
func Foo(i int, s string) bool {
	return strconv.Itoa(i) == s
}

//Wrapper 装饰器函数
func Wrapper(f Decoer) Decoer {
	return func(i int, s string) bool {
		fmt.Println("Before...")
		result := f(i, s)
		fmt.Println("After...")
		return result
	}
}
