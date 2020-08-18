package visitor

import "testing"

func TestVisitor(t *testing.T) {
	objs := &ObjectStructure{
		persons: make([]Person, 0),
	}

	man := &Man{
		name: "男人",
	}
	woman := &Woman{
		name: "女人",
	}

	objs.Attach(man)
	objs.Attach(woman)

	success := &Success{
		name: "成功",
	}

	objs.Display(success)
}
