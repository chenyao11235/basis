package strategy

import (
	"fmt"
	"sync"
)

/*策略模式
翻译成中文就是：定义一族算法类，将每个算法分别封装起来，让它们可以互相替换。
策略模式可以使算法的变化独立于使用它们的客户端（这里的客户端代指使用算法的代码）。
*/

//Strategy 策略接口
type Strategy interface {
	Set(string)
	algorithmInterface()
}

//ConcreteStrategyA 接口实现A
type ConcreteStrategyA struct {
	name string
}

func (s *ConcreteStrategyA) Set(name string) {
	s.name = name
}

func (s *ConcreteStrategyA) algorithmInterface() {
	fmt.Printf("我是 %s, 我将执行算法A\n", s.name)
}

//ConcreteStrategyB 接口实现B
type ConcreteStrategyB struct {
	name string
}

func (s *ConcreteStrategyB) Set(name string) {
	s.name = name
}

func (s *ConcreteStrategyB) algorithmInterface() {
	fmt.Println("算法B")
}

var (
	strategyMap map[string]Strategy
	setMap      sync.Once
)

//NewStrategyNoStatus 对客户端屏蔽策略的创建细节，使用工厂函数创建  适用于无状态的对象，通过反射可以解决这一问题
func NewStrategyNoStatus(kind string) Strategy {
	// 只需要设置一次, 这种情况适用于Strategy的实现是无状态的，没有属性字段，可以事先把对象存储在map
	// 如果是有状态的，就需要每次获取strategy的都创建
	setMap.Do(func() {
		strategyMap = make(map[string]Strategy)
		strategyMap["A"] = new(ConcreteStrategyA)
		strategyMap["B"] = new(ConcreteStrategyB)
	})

	if kind == "" {
		fmt.Println("kind can not be empty.")
		return nil
	}

	if impl, ok := strategyMap[kind]; ok {
		return impl
	}

	return nil
}
