package memento

import (
	"fmt"
)

// 备忘录模式： 在不破坏对象封装型的前提下，捕获对象的内部状态，并在该对象之外保存这个状态，这样就可以将该对象恢复到原先保存的状态

type Originator struct {
	state string
}

type Memento struct {
	state string
}

// 创建备忘录
func (o *Originator) CreateMemento() *Memento {
	return &Memento{
		state: o.state,
	}
}

//回复备忘录
func (o *Originator) SetMemento(m *Memento) {
	o.state = m.state
}

// 展示
func (o *Originator) Show() {
	fmt.Println(o.state)
}

// 管理者类, 如果需要保存多个快照的话，可以使用栈（stack）来保存多个memento
type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) SetMemento(m *Memento) {
	c.memento = m
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}
