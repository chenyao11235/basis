package linearlist

import "testing"

func TestInsertHead(t *testing.T) {
	l := NewLinearList()
	for i := 0; i < 10; i++ {
		l.InsterHead(i + 1)
	}
	l.Print()
}

func TestInsertTail(t *testing.T) {
	l := NewLinearList()
	for i := 0; i < 10; i++ {
		l.InsertTail(i + 1)
	}
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	l := NewLinearList()
	for i := 0; i < 10; i++ {
		l.InsertTail(i + 1)
	}

	t.Log(l.FindByIndex(2).GetValue())
	t.Log(l.FindByIndex(5).GetValue())
	t.Log(l.FindByIndex(6).GetValue())
	t.Log(l.FindByIndex(9).GetValue())
}

func TestDeleteNode(t *testing.T) {
	l := NewLinearList()
	for i := 0; i < 10; i++ {
		l.InsertTail(i + 1)
	}
	l.DeleteNode(l.head.next)
	l.Print()

	l.DeleteNode(l.head.next.next)
	l.Print()
}

func TestIsPalindrome(t *testing.T) {
	l := NewLinearList()
	l.InsertTail(0)
	l.InsertTail(1)
	// l.InsertTail(1)
	l.InsertTail(2)
	l.InsertTail(1)
	l.InsertTail(0)
	l.InsertTail(0)

	t.Log(isPalindrome(l))
	t.Log(isPalindrome1(l))
}

func TestReverse(t *testing.T) {
	l := NewLinearList()
	l.InsertTail(0)
	l.InsertTail(1)
	l.InsertTail(2)
	l.Print()

	l.Reverse()
	l.Print()
}

func TestLru(t *testing.T) {
	l := NewLru(5)
	l.add(0)
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	l.Print()
}
