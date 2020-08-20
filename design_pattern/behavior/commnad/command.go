package command

import (
	"fmt"
)

// 做烧烤的师傅
type Barbecuer struct {
}

func (b *Barbecuer) BakeMutton() {
	fmt.Println("我要烤一个羊肉串")
}

func (b *Barbecuer) BakeChickenWing() {
	fmt.Println("我要烤一个鸡翅膀")
}

type Command interface {
	ExcuteCommand()
}

// 烤羊肉串的命令
type BakeMuttonCommand struct {
	receiver *Barbecuer
}

func (c *BakeMuttonCommand) ExcuteCommand() {
	c.receiver.BakeMutton()
}

// 烤鸡翅膀的命令
type BakeChickendWingCommand struct {
	receiver *Barbecuer
}

func (c *BakeChickendWingCommand) ExcuteCommand() {
	c.receiver.BakeChickenWing()
}

// 服务员，接收顾客的指令，并把指令传达给做烧烤的师傅
// 这里是可以优化一下的：服务员不应该接受到一个命令就通知师傅一下，而是接收顾客的所有点单之后一次性通知给师傅
type Waiter struct {
	command Command
}

func (w *Waiter) SetOrder(command Command) {
	w.command = command
}

func (w *Waiter) Notify() {
	w.command.ExcuteCommand()
}
