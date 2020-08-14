package stack

//Browser 浏览器
type Browser struct {
	forwardStack Stack
	backStack    Stack
}

//NewBrowser 新建一个浏览器
func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewStackBaseOnLinkedList(),
		backStack:    NewStackBaseOnArray(),
	}
}

//CanForward 是否还能前进
func (b *Browser) CanForward() bool {
	if b.forwardStack.IsEmpty() {
		return true
	}
	return false
}

//CanBack 是否还能后退
func (b *Browser) CanBack() bool {
	if b.backStack.IsEmpty() {
		return true
	}
	return false
}

//Open 在chrome中打开了一个新的标签页
func (b *Browser) Open(addr string) {
	b.forwardStack.Flush()
}

//PushBack 相当于在一个标签页中不断的跳转到新的地址
func (b *Browser) PushBack(addr string) {
	b.backStack.Push(addr)
}

//Forward 前进
func (b *Browser) Forward() {
	if b.forwardStack.IsEmpty() {
		return
	}
	top := b.forwardStack.Pop()
	b.backStack.Push(top)
}

//Back 后退
func (b *Browser) Back() {
	if b.backStack.IsEmpty() {
		return
	}
	top := b.backStack.Pop()
	b.forwardStack.Push(top)
}
