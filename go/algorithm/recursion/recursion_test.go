package recursion

import "testing"

func TestFac(t *testing.T) {
	f := NewFactorial(5)
	f.Factorial(5)
	f.Print(5)
}

func TestFib(t *testing.T) {
	f := NewFibs(10)
	f.Fibonacci(10)
	f.Print(10)
}
