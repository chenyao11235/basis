package bad

import (
	"fmt"
)

// 下面这段代码虽然使用了面向对象的方式，但是如果需要添加一个恋爱的action就需要修改Man，Woman两个类，不符合开闭原则

type Person interface {
	GetConclusion()
}

type Man struct {
	action string
}

func (m *Man) GetConclusion() {
	if m.action == "成功" {
		fmt.Println("男人成功的时候，背后多半都有一个伟大的女人。")
	}
	if m.action == "失败" {
		fmt.Println("男人失败的时候，低头喝闷酒，谁也不用劝。")
	}
}

type Woman struct {
	action string
}

func (w *Woman) GetConclusion() {
	if w.action == "成功" {
		fmt.Println("女人成功的时候，背后多半都有一个不成功的男人。")
	}
	if w.action == "失败" {
		fmt.Println("女人失败的时候，眼泪汪汪，谁也劝不了。")
	}
}
