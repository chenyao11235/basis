package chain

import (
	"fmt"
)

/*职责链模式
将请求的发送和接收解耦，让多个接收对象都有机会处理这个请求，将这些接收对象穿成串成一条链
并沿着这条链传递这个请求，直到某个接收对象能够处理为止，还有一种应用场景是需要请求经过每个接收对象的处理，而不是每个对象都能处理

一种方式是在每个实现类中设置一个继任者类，通过这个继任者类形成一个链式结构
另一种方式是不用设置继任者类，将这些实现类放到一个数组中
*/

//Handler 处理接口
type Handler interface {
	// 处理请求的方法
	HandleReq(req int)
	//设置下一个继任者
	SetSuccessor(successor Handler)
}

//HandlerA 处理接口实现A
type HandlerA struct {
	successor Handler
}

//HandleReq 处理方法
func (h *HandlerA) HandleReq(req int) {
	if req >= 0 && req < 10 {
		fmt.Println("req被HandlerA处理了")
	} else if h.successor != nil {
		h.successor.HandleReq(req)
	}
}

//SetSuccessor 设置继任者
func (h *HandlerA) SetSuccessor(handler Handler) {
	h.successor = handler
}

//HandlerB 处理接口实现B
type HandlerB struct {
	successor Handler
}

//HandleReq 处理方法
func (h *HandlerB) HandleReq(req int) {
	if req >= 10 && req < 20 {
		fmt.Println("req被HandlerB处理了")
	} else if h.successor != nil {
		h.successor.HandleReq(req)
	}
}

//SetSuccessor 设置继任者
func (h *HandlerB) SetSuccessor(handler Handler) {
	h.successor = handler
}

//HandlerC 处理接口实现B
type HandlerC struct {
	successor Handler
}

//HandleReq 处理方法
func (h *HandlerC) HandleReq(req int) {
	if req >= 20 {
		fmt.Println("req被HandlerC处理了")
	} else if h.successor != nil {
		h.successor.HandleReq(req)
	}
}

//SetSuccessor 设置继任者
func (h *HandlerC) SetSuccessor(handler Handler) {
	h.successor = handler
}

//HandlerChain 处理链类
type HandlerChain struct {
	// 使用数组保存接收者的化就不必设置successor了
	filters []Handler
}

func (c *HandlerChain) addHandler(handler Handler) {
	c.filters = append(c.filters, handler)
}

//Handle 处理
func (c *HandlerChain) Handle(req int) {
	for _, handler := range c.filters {
		handler.HandleReq(req)
	}
}

//NewHandlerChain 新建
func NewHandlerChain() *HandlerChain {
	chain := new(HandlerChain)

	chain.addHandler(new(HandlerA))
	chain.addHandler(new(HandlerB))
	chain.addHandler(new(HandlerC))

	return chain
}
