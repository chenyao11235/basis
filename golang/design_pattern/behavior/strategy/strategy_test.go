package strategy

import "testing"

// 决定使用哪一种策略类型
func decideStrategyKind() string {
	return "A"
}

func TestStrategy(t *testing.T) {
	strategy1 := NewStrategyNoStatus(decideStrategyKind())
	strategy1.Set("alex")
	strategy1.algorithmInterface()

	strategy2 := NewStrategyNoStatus(decideStrategyKind())
	strategy2.Set("eric")
	strategy2.algorithmInterface()
	// strategy1被strategy2修改，说明有状态的对象不行
	strategy1.algorithmInterface()
}
