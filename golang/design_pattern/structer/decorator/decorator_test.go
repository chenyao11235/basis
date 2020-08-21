package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoratorFunc(t *testing.T) {
	foo := Wrapper(Foo)
	assert.Equal(t, true, foo(1, "1"), "")
}

func TestDecoratorClass(t *testing.T) {
	c := WarpMulDecorator(&BasicComponent{}, 2)
	assert.Equal(t, 2, c.Calc(), "")
}
