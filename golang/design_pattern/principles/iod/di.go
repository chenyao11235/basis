package iod

/*依赖反转原则之依赖注入
那到底什么是依赖注入呢？
我们用一句话来概括就是：不通过 new() 的方式在类内部创建依赖类对象，
而是将依赖的类对象在外部创建好之后，通过构造函数、函数参数等方式传递（或注入）给类使用。
使用依赖注入的方式可以使得写的代码易于测试，写测试用例比较方便
*/

// -------------------------------------->非依赖注入的实现方法

//Notification 告警
type Notification struct {
	messageSender *MessageSender
}

func (n *Notification) sendMessage() {
	// 外部依赖类是在内部初始化的
	n.messageSender = &MessageSender{}
	n.messageSender.send()
}

//MessageSender 消息发送器(邮件发送器，短信发送器....)
type MessageSender struct {
}

func (ms *MessageSender) send() {

}

// ------------------------------------> 通过依赖注入的方法实现

/*
通过依赖注入的方式来将依赖的类对象传递进来，
这样就提高了代码的扩展性，我们可以灵活地替换依赖的类。
这一点在我们之前讲“开闭原则”的时候也提到过。
当然，上面代码还有继续优化的空间，我们还可以把 MessageSender 定义成接口，基于接口而非实现编程。
*/

//NewNotification 初始化
func NewNotification(messageSender *MessageSender) *Notification {
	return &Notification{
		messageSender: messageSender,
	}
}
