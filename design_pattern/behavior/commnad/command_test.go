package command

import "testing"

func TestCommand(t *testing.T) {
	boy := &Barbecuer{}

	command1 := &BakeMuttonCommand{
		receiver: boy,
	}

	command2 := &BakeChickendWingCommand{
		receiver: boy,
	}

	girl := &Waiter{}
	girl.SetOrder(command1)
	girl.Notify()
	girl.SetOrder(command2)
	girl.Notify()
}
