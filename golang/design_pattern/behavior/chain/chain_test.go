package chain

import "testing"

func TestChain1(t *testing.T) {
	h1 := new(HandlerA)
	h2 := new(HandlerB)
	h3 := new(HandlerC)

	h1.SetSuccessor(h2)
	h2.SetSuccessor(h3)

	requests := []int{1, 5, 23, 16, 23, 32, 7, 9}
	for _, req := range requests {
		h1.HandleReq(req)
	}
}

func TestChain2(t *testing.T) {
	chain := NewHandlerChain()
	requests := []int{1, 5, 23, 16, 23, 32, 7, 9}

	for _, req := range requests {
		chain.Handle(req)
	}
}
