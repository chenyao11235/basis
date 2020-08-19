package memento

import "testing"

func TestMemento(t *testing.T) {
	o := new(Originator)
	o.state = "On"
	o.Show()

	c := new(Caretaker)
	c.memento = o.CreateMemento()

	o.state = "Off"
	o.Show()

	o.SetMemento(c.GetMemento())
	o.Show()
}
